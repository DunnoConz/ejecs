package parser

import (
	"fmt"
	"strings"

	"github.com/ejecs/ejecs/internal/ast"
	"github.com/ejecs/ejecs/internal/lexer"
)

// Parser represents a JECS parser
type Parser struct {
	l *lexer.Lexer

	curToken  lexer.Token
	peekToken lexer.Token

	errors []string
}

// New creates a new Parser instance
func New(input string) *Parser {
	l := lexer.New(input)
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Parse parses the JECS content and returns an AST
func (p *Parser) Parse() (*ast.Program, error) {
	program := &ast.Program{}

	for p.curToken.Type != lexer.EOF {
		switch p.curToken.Type {
		case lexer.COMPONENT:
			if comp := p.parseComponent(); comp != nil {
				program.Components = append(program.Components, comp)
			}
		case lexer.RELATIONSHIP:
			if rel := p.parseRelationship(); rel != nil {
				program.Relationships = append(program.Relationships, rel)
			}
		case lexer.SYSTEM:
			if sys := p.parseSystem(); sys != nil {
				program.Systems = append(program.Systems, sys)
			}
		default:
			p.nextToken()
		}
	}

	if len(p.errors) > 0 {
		return nil, fmt.Errorf("parse errors: %v", p.errors)
	}

	return program, nil
}

func (p *Parser) parseComponent() *ast.Component {
	comp := &ast.Component{}

	// Skip 'component' keyword
	p.nextToken()

	// Parse component name
	if p.curToken.Type != lexer.IDENT {
		p.errors = append(p.errors, fmt.Sprintf("expected component name, got %s", p.curToken.Type))
		return nil
	}
	comp.Name = p.curToken.Literal

	// Skip name
	p.nextToken()

	// Expect '{'
	if p.curToken.Type != lexer.LBRACE {
		p.errors = append(p.errors, fmt.Sprintf("expected '{', got %s", p.curToken.Type))
		return nil
	}
	p.nextToken()

	// Parse fields
	for p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
		if field := p.parseField(); field != nil {
			comp.Fields = append(comp.Fields, field)
		}
		p.nextToken()
	}

	return comp
}

func (p *Parser) parseField() *ast.Field {
	if p.curToken.Type != lexer.IDENT {
		p.errors = append(p.errors, fmt.Sprintf("expected field name, got %s", p.curToken.Type))
		return nil
	}

	field := &ast.Field{Name: p.curToken.Literal}

	// Skip field name
	p.nextToken()

	// Expect ':'
	if p.curToken.Type != lexer.COLON {
		p.errors = append(p.errors, fmt.Sprintf("expected ':', got %s", p.curToken.Type))
		return nil
	}
	p.nextToken()

	// Parse type
	if p.curToken.Type != lexer.IDENT {
		p.errors = append(p.errors, fmt.Sprintf("expected type name, got %s", p.curToken.Type))
		return nil
	}
	field.Type = p.curToken.Literal

	// Skip type
	p.nextToken()

	// Expect ';'
	if p.curToken.Type != lexer.SEMICOLON {
		p.errors = append(p.errors, fmt.Sprintf("expected ';', got %s", p.curToken.Type))
		return nil
	}

	return field
}

func (p *Parser) parseRelationship() *ast.Relationship {
	rel := &ast.Relationship{}

	// Skip 'relationship' keyword
	p.nextToken()

	// Parse relationship name
	if p.curToken.Type != lexer.IDENT {
		p.errors = append(p.errors, fmt.Sprintf("expected relationship name, got %s", p.curToken.Type))
		return nil
	}
	rel.Name = p.curToken.Literal

	// Skip name
	p.nextToken()

	// Expect '{}'
	if p.curToken.Type != lexer.LBRACE {
		p.errors = append(p.errors, fmt.Sprintf("expected '{', got %s", p.curToken.Type))
		return nil
	}
	p.nextToken()

	if p.curToken.Type != lexer.RBRACE {
		p.errors = append(p.errors, fmt.Sprintf("expected '}', got %s", p.curToken.Type))
		return nil
	}

	return rel
}

func (p *Parser) parseSystem() *ast.System {
	sys := &ast.System{}

	// Skip 'system' keyword
	p.nextToken()

	// Parse system name
	if p.curToken.Type != lexer.IDENT {
		p.errors = append(p.errors, fmt.Sprintf("expected system name, got %s", p.curToken.Type))
		return nil
	}
	sys.Name = p.curToken.Literal

	// Skip name
	p.nextToken()

	// Expect '{'
	if p.curToken.Type != lexer.LBRACE {
		p.errors = append(p.errors, fmt.Sprintf("expected '{', got %s", p.curToken.Type))
		return nil
	}
	p.nextToken()

	// Parse query and code
	for p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
		switch p.curToken.Type {
		case lexer.QUERY:
			if query := p.parseQuery(); query != nil {
				sys.Query = query
			}
		case lexer.RUN:
			if code := p.parseCode(); code != "" {
				sys.Code = code
			}
		}
		p.nextToken()
	}

	return sys
}

func (p *Parser) parseQuery() *ast.Query {
	query := &ast.Query{}

	// Skip 'query' keyword and ':'
	p.nextToken()
	if p.curToken.Type != lexer.COLON {
		p.errors = append(p.errors, fmt.Sprintf("expected ':', got %s", p.curToken.Type))
		return nil
	}
	p.nextToken()

	// Expect '('
	if p.curToken.Type != lexer.LPAREN {
		p.errors = append(p.errors, fmt.Sprintf("expected '(', got %s", p.curToken.Type))
		return nil
	}
	p.nextToken()

	// Parse components list
	for p.curToken.Type != lexer.RPAREN && p.curToken.Type != lexer.EOF {
		if p.curToken.Type == lexer.IDENT {
			query.Components = append(query.Components, p.curToken.Literal)
			p.nextToken()
		} else if p.curToken.Type == lexer.PAIR {
			if rel := p.parseRelationQuery(); rel != nil {
				query.Relations = append(query.Relations, rel)
			}
		} else {
			p.errors = append(p.errors, fmt.Sprintf("expected component name or pair, got %s", p.curToken.Type))
			return nil
		}

		if p.curToken.Type == lexer.COMMA {
			p.nextToken()
		}
	}

	// Expect ')'
	if p.curToken.Type != lexer.RPAREN {
		p.errors = append(p.errors, fmt.Sprintf("expected ')', got %s", p.curToken.Type))
		return nil
	}
	p.nextToken()

	// Expect ';'
	if p.curToken.Type != lexer.SEMICOLON {
		p.errors = append(p.errors, fmt.Sprintf("expected ';', got %s", p.curToken.Type))
		return nil
	}

	return query
}

func (p *Parser) parseRelationQuery() *ast.Relation {
	// Skip 'pair' keyword
	p.nextToken()

	// Expect '('
	if p.curToken.Type != lexer.LPAREN {
		p.errors = append(p.errors, fmt.Sprintf("expected '(', got %s", p.curToken.Type))
		return nil
	}
	p.nextToken()

	// Parse relationship type
	if p.curToken.Type != lexer.IDENT {
		p.errors = append(p.errors, fmt.Sprintf("expected relationship name, got %s", p.curToken.Type))
		return nil
	}
	rel := &ast.Relation{Type: p.curToken.Literal}
	p.nextToken()

	// Expect comma
	if p.curToken.Type != lexer.COMMA {
		p.errors = append(p.errors, fmt.Sprintf("expected ',', got %s", p.curToken.Type))
		return nil
	}
	p.nextToken()

	// Parse target (can be identifier or asterisk)
	if p.curToken.Type != lexer.IDENT && p.curToken.Type != lexer.ASTERISK {
		p.errors = append(p.errors, fmt.Sprintf("expected target or '*', got %s", p.curToken.Type))
		return nil
	}
	rel.Target = p.curToken.Literal
	p.nextToken()

	// Expect ')'
	if p.curToken.Type != lexer.RPAREN {
		p.errors = append(p.errors, fmt.Sprintf("expected ')', got %s", p.curToken.Type))
		return nil
	}
	p.nextToken()

	return rel
}

func (p *Parser) parseCode() string {
	// Skip 'run' keyword
	p.nextToken()

	// Expect '('
	if p.curToken.Type != lexer.LPAREN {
		p.errors = append(p.errors, fmt.Sprintf("expected '(', got %s", p.curToken.Type))
		return ""
	}
	p.nextToken()

	// Parse parameter list
	var params []string
	for p.curToken.Type != lexer.RPAREN && p.curToken.Type != lexer.EOF {
		if p.curToken.Type == lexer.IDENT {
			params = append(params, p.curToken.Literal)
			p.nextToken()
			if p.curToken.Type == lexer.COMMA {
				p.nextToken()
			}
		} else {
			p.errors = append(p.errors, fmt.Sprintf("expected parameter name, got %s", p.curToken.Type))
			return ""
		}
	}

	// Expect ')'
	if p.curToken.Type != lexer.RPAREN {
		p.errors = append(p.errors, fmt.Sprintf("expected ')', got %s", p.curToken.Type))
		return ""
	}
	p.nextToken()

	// Expect '{'
	if p.curToken.Type != lexer.LBRACE {
		p.errors = append(p.errors, fmt.Sprintf("expected '{', got %s", p.curToken.Type))
		return ""
	}
	p.nextToken()

	// Parse code block
	var code strings.Builder
	var needsSpace bool

	for p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
		switch p.curToken.Type {
		case lexer.SEMICOLON:
			if p.peekToken.Type == lexer.RBRACE {
				code.WriteString(";")
			} else {
				code.WriteString(";\n")
			}
			needsSpace = false
		case lexer.COMMA:
			code.WriteString(", ")
			needsSpace = false
		case lexer.LPAREN, lexer.RPAREN:
			code.WriteString(p.curToken.Literal)
			needsSpace = false
		case lexer.COLON, lexer.DOT:
			code.WriteString(p.curToken.Literal)
			needsSpace = false
		default:
			if needsSpace {
				code.WriteString(" ")
			}
			code.WriteString(p.curToken.Literal)
			needsSpace = true
		}
		p.nextToken()
	}

	// Expect '}'
	if p.curToken.Type != lexer.RBRACE {
		p.errors = append(p.errors, fmt.Sprintf("expected '}', got %s", p.curToken.Type))
		return ""
	}

	return code.String()
}
