package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/ejecs/ejecs/internal/ast"
	"github.com/ejecs/ejecs/internal/lexer"
	"github.com/ejecs/ejecs/internal/token"
)

// Parser represents a JECS parser
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

// Error represents a parsing error
type Error struct {
	Line    int
	Column  int
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("line %d, column %d: %s", e.Line, e.Column, e.Message)
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
func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := &ast.Program{
		Statements: []ast.Node{},
	}

	for p.curToken.Type != token.EOF {
		var stmt ast.Node
		var err error

		switch p.curToken.Type {
		case token.COMPONENT:
			stmt, err = p.parseComponent()
		case token.RELATIONSHIP:
			stmt, err = p.parseRelationship()
		case token.SYSTEM:
			stmt, err = p.parseSystem()
		default:
			return nil, fmt.Errorf("unexpected token %s", p.curToken.Type)
		}

		if err != nil {
			return nil, err
		}

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program, nil
}

func (p *Parser) parseComponent() (*ast.Component, error) {
	comp := &ast.Component{}

	// Skip 'component' keyword
	p.nextToken()

	// Parse component name
	if !p.curTokenIs(token.IDENT) {
		return nil, p.newError("expected component name, got %s", p.curToken.Type)
	}
	comp.Name = p.curToken.Literal

	// Skip name
	p.nextToken()

	// Expect '{'
	if !p.curTokenIs(token.LBRACE) {
		return nil, p.newError("expected '{', got %s", p.curToken.Type)
	}
	p.nextToken()

	// Parse fields
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		field, err := p.parseField()
		if err != nil {
			return nil, err
		}
		if field != nil {
			comp.Fields = append(comp.Fields, field)
		}
	}

	// Expect closing brace RBRACE
	if !p.curTokenIs(token.RBRACE) {
		return nil, p.newError("expected '}' to close component, got %s", p.curToken.Type)
	}

	return comp, nil
}

func (p *Parser) parseField() (*ast.Field, error) {
	field := &ast.Field{}

	// Check if the type is 'table'
	if p.curTokenIs(token.TABLE) {
		fmt.Println("DEBUG: p.curTokenIs(token.TABLE) is TRUE") // DEBUG
		os.Stdout.Sync()                                        // DEBUG
		field.Type = "table"
		p.nextToken() // Consume 'table'

		// Expect <
		if !p.curTokenIs(token.LT) {
			return nil, p.newError("expected '<' after table keyword, got %s", p.curToken.Type)
		}
		p.nextToken() // Consume <

		// Parse Key Type (expect IDENT)
		if !p.curTokenIs(token.IDENT) && !p.curTokenIs(token.STRING) { // Allow string literal too?
			return nil, p.newError("expected table key type (identifier or string), got %s", p.curToken.Type)
		}
		field.MapKeyType = p.curToken.Literal
		p.nextToken() // Consume key type

		// Expect ,
		if !p.curTokenIs(token.COMMA) {
			return nil, p.newError("expected ',' between table key and value types, got %s", p.curToken.Type)
		}
		p.nextToken() // Consume ,

		// Parse Value Type (expect IDENT)
		if !p.curTokenIs(token.IDENT) {
			return nil, p.newError("expected table value type (identifier), got %s", p.curToken.Type)
		}
		field.MapValueType = p.curToken.Literal
		p.nextToken() // Consume value type

		// Expect >
		if !p.curTokenIs(token.GT) {
			return nil, p.newError("expected '>' after table value type, got %s", p.curToken.Type)
		}
		p.nextToken() // Consume >

		// Now parse the field name (IDENT) that follows the table type definition
		if !p.curTokenIs(token.IDENT) {
			return nil, p.newError("expected field name after table type, got %s", p.curToken.Type)
		}
		field.Name = p.curToken.Literal
		p.nextToken() // Consume field name

	} else if p.curTokenIs(token.IDENT) {
		// --- Regular Type Parsing (Type name;) ---
		field.Type = p.curToken.Literal

		// Check for optional type
		if p.peekTokenIs(token.QUESTION) {
			field.Optional = true
			p.nextToken() // Consume '?'
		}

		// Skip type (and optional '?')
		p.nextToken()

		// Expect field name
		if !p.curTokenIs(token.IDENT) {
			return nil, p.newError("expected field name, got %s", p.curToken.Type)
		}
		field.Name = p.curToken.Literal
		p.nextToken() // Consume field name
	} else {
		// Error: Expected type name (IDENT or TABLE)
		return nil, p.newError("expected field type (identifier or 'table'), got %s", p.curToken.Type)
	}

	// --- Common Logic: Default Value and Semicolon ---

	// Expect optional default value
	if p.curTokenIs(token.ASSIGN) {
		p.nextToken()                // Consume '='
		startLine := p.curToken.Line // For error reporting or storing skipped value
		startCol := p.curToken.Column

		// If default value starts with {, skip until matching }
		if p.curTokenIs(token.LBRACE) {
			braceLevel := 1
			for braceLevel > 0 {
				p.nextToken()
				if p.curTokenIs(token.EOF) {
					return nil, p.newErrorf(startLine, startCol, "unmatched '{'")
				}
				if p.curTokenIs(token.LBRACE) {
					braceLevel++
				} else if p.curTokenIs(token.RBRACE) {
					braceLevel--
				}
			}
			// At this point, curToken is the matching RBRACE
			field.DefaultValue = "{...}" // Indicate skipped table
			p.nextToken()                // Consume the final RBRACE
		} else if p.curTokenIs(token.STRING) || p.curTokenIs(token.INT) || p.curTokenIs(token.FLOAT) || p.curTokenIs(token.TRUE) || p.curTokenIs(token.FALSE) {
			// Handle simple literals
			field.DefaultValue = p.curToken.Literal
			p.nextToken() // Consume the literal
		} else {
			// Unknown default value - skip until semicolon
			fmt.Printf("DEBUG: Skipping unknown default value starting with %s (%q) at %d:%d\n",
				p.curToken.Type, p.curToken.Literal, p.curToken.Line, p.curToken.Column)
			os.Stdout.Sync()
			for !p.curTokenIs(token.SEMICOLON) && !p.curTokenIs(token.EOF) {
				p.nextToken()
			}
			if p.curTokenIs(token.EOF) {
				return nil, p.newErrorf(startLine, startCol, "missing ';' after complex default value?")
			}
			// Do not consume the semicolon here, let the final check handle it.
			field.DefaultValue = "<skipped>" // Indicate skipped complex value
			// Current token should now be SEMICOLON
		}
	}

	// Expect ';'
	if !p.curTokenIs(token.SEMICOLON) {
		return nil, p.newError("expected ';', got %s (%q)", p.curToken.Type, p.curToken.Literal)
	}
	p.nextToken() // Consume ';'

	return field, nil
}

func (p *Parser) parseRelationship() (*ast.Relationship, error) {
	rel := &ast.Relationship{}

	// Parse relationship type if present
	if p.curToken.Type == token.AT {
		p.nextToken()
		rel.Type = p.curToken.Literal
		p.nextToken()
	}

	// Expect 'relationship' keyword
	if p.curToken.Type != token.RELATIONSHIP {
		return nil, fmt.Errorf("expected 'relationship', got %s", p.curToken.Type)
	}
	p.nextToken()

	// Parse relationship name
	if p.curToken.Type != token.IDENT {
		return nil, fmt.Errorf("expected identifier, got %s", p.curToken.Type)
	}
	rel.Name = p.curToken.Literal
	p.nextToken()

	// Expect opening brace
	if p.curToken.Type != token.LBRACE {
		return nil, fmt.Errorf("expected '{', got %s", p.curToken.Type)
	}
	p.nextToken()

	// Parse child field
	if p.curToken.Type != token.IDENT || p.curToken.Literal != "child" {
		return nil, fmt.Errorf("expected 'child', got %s", p.curToken.Type)
	}
	p.nextToken()

	if p.curToken.Type != token.COLON {
		return nil, fmt.Errorf("expected ':', got %s", p.curToken.Type)
	}
	p.nextToken()

	if p.curToken.Type != token.IDENT {
		return nil, fmt.Errorf("expected identifier, got %s", p.curToken.Type)
	}
	rel.Child = p.curToken.Literal
	p.nextToken()

	// Parse parent field
	if p.curToken.Type != token.IDENT || p.curToken.Literal != "parent" {
		return nil, fmt.Errorf("expected 'parent', got %s", p.curToken.Type)
	}
	p.nextToken()

	if p.curToken.Type != token.COLON {
		return nil, fmt.Errorf("expected ':', got %s", p.curToken.Type)
	}
	p.nextToken()

	if p.curToken.Type != token.IDENT {
		return nil, fmt.Errorf("expected identifier, got %s", p.curToken.Type)
	}
	rel.Parent = p.curToken.Literal
	p.nextToken()

	// Expect closing brace
	if p.curToken.Type != token.RBRACE {
		return nil, fmt.Errorf("expected '}', got %s", p.curToken.Type)
	}

	return rel, nil
}

// Renaming to reflect it parses the content *inside* the query parens/braces
func (p *Parser) parseQueryContent() (*ast.Query, error) {
	query := &ast.Query{
		Components: []string{},
		Relations:  []*ast.Relation{},
	}

	// Expect first component name or relation
	for !p.curTokenIs(token.RPAREN) && !p.curTokenIs(token.EOF) { // Stop at RPAREN for query()
		if p.curTokenIs(token.IDENT) {
			// Check if it's a relation type (e.g., parent(...))
			if p.peekTokenIs(token.LPAREN) {
				rel, err := p.parseRelationCall()
				if err != nil {
					return nil, err
				}
				query.Relations = append(query.Relations, rel)
			} else {
				// Regular component name
				query.Components = append(query.Components, p.curToken.Literal)
				p.nextToken() // Consume component name
			}
		} else {
			return nil, p.newError("expected component name or relation type in query, got %s", p.curToken.Type)
		}

		// Expect comma or closing paren
		if p.curTokenIs(token.COMMA) {
			p.nextToken() // Consume comma, continue loop
		} else if !p.curTokenIs(token.RPAREN) {
			return nil, p.newError("expected ',' or ')' in query, got %s", p.curToken.Type)
		}
	}

	return query, nil
}

// Parses a relation call like parent(Component)
func (p *Parser) parseRelationCall() (*ast.Relation, error) {
	rel := &ast.Relation{}

	if !p.curTokenIs(token.IDENT) {
		return nil, p.newError("expected relation type identifier, got %s", p.curToken.Type)
	}
	rel.Type = p.curToken.Literal
	p.nextToken() // Consume relation type

	if !p.curTokenIs(token.LPAREN) {
		return nil, p.newError("expected '(' after relation type, got %s", p.curToken.Type)
	}
	p.nextToken() // Consume (

	if !p.curTokenIs(token.IDENT) {
		return nil, p.newError("expected component name inside relation parentheses, got %s", p.curToken.Type)
	}
	rel.Component = p.curToken.Literal
	p.nextToken() // Consume component name

	if !p.curTokenIs(token.RPAREN) {
		return nil, p.newError("expected ')' after relation component name, got %s", p.curToken.Type)
	}
	p.nextToken() // Consume )

	return rel, nil
}

// Simplified parseSystem - Expects query() first, then optionals
func (p *Parser) parseSystem() (*ast.System, error) {
	system := &ast.System{}
	startLine := p.curToken.Line // Record line/col of SYSTEM token
	startCol := p.curToken.Column

	// Current token is SYSTEM (checked by ParseProgram)
	p.nextToken() // Consume SYSTEM keyword

	// Now expect system name
	if !p.curTokenIs(token.IDENT) {
		return nil, p.newErrorf(startLine, startCol, "expected system name after 'system' keyword, got %s", p.curToken.Type)
	}
	system.Name = p.curToken.Literal
	system.Line = p.curToken.Line // Update line/col to name token
	system.Column = p.curToken.Column
	p.nextToken() // Consume name

	// Expect opening brace for system body
	if !p.curTokenIs(token.LBRACE) {
		return nil, p.newError("expected '{' after system name, got %s", p.curToken.Type)
	}
	p.nextToken() // Consume {

	// Parse optional blocks inside the system body
	codeParsed := false
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		switch p.curToken.Type {
		case token.QUERY:
			if system.Query != nil {
				return nil, p.newError("duplicate query block")
			}
			p.nextToken() // Consume 'query'
			if !p.curTokenIs(token.LPAREN) {
				return nil, p.newError("expected '(' after query keyword, got %s", p.curToken.Type)
			}
			p.nextToken() // Consume (
			queryContent, err := p.parseQueryContent()
			if err != nil {
				return nil, err
			}
			system.Query = queryContent
			if !p.curTokenIs(token.RPAREN) { // parseQueryContent stops at RPAREN
				return nil, p.newError("expected ')' to close query, got %s", p.curToken.Type)
			}
			p.nextToken() // Consume )
		case token.IDENT:
			if p.curToken.Literal == "params" {
				if system.Parameters != nil {
					return nil, p.newError("duplicate params block")
				}
				params, err := p.parseParametersBlock()
				if err != nil {
					return nil, err
				}
				system.Parameters = params
			} else {
				return nil, p.newError("unexpected identifier '%s' in system body", p.curToken.Literal)
			}
		case token.FREQUENCY:
			if system.Frequency != "" {
				return nil, p.newError("duplicate frequency definition")
			}
			p.nextToken() // Consume 'frequency'
			if !p.curTokenIs(token.COLON) {
				return nil, p.newError("expected ':' after frequency, got %s", p.curToken.Type)
			}
			p.nextToken() // Consume ':'

			// Skip tokens until end of line or next block starter
			startLine := p.curToken.Line
			for p.curToken.Line == startLine &&
				!p.curTokenIs(token.EOF) &&
				!p.curTokenIs(token.PRIORITY) &&
				!p.curTokenIs(token.LBRACE) {
				// TODO: Actually parse/store this value correctly
				p.nextToken()
			}
			system.Frequency = "<skipped>" // Placeholder
			// Do not consume the token that stopped the loop (priority, {, or EOF)

		case token.PRIORITY:
			if system.Priority != "" {
				return nil, p.newError("duplicate priority definition")
			}
			p.nextToken() // Consume 'priority'
			if !p.curTokenIs(token.COLON) {
				return nil, p.newError("expected ':' after priority, got %s", p.curToken.Type)
			}
			p.nextToken() // Consume ':'
			if !p.curTokenIs(token.INT) {
				// For now, let's skip non-int priority too until we parse expressions
				startLine := p.curToken.Line
				for p.curToken.Line == startLine && !p.curTokenIs(token.EOF) && !p.curTokenIs(token.LBRACE) {
					p.nextToken()
				}
				system.Priority = "<skipped>"
			} else {
				system.Priority = p.curToken.Literal
				p.nextToken() // Consume value
			}
		case token.LBRACE: // Start of the code block
			if codeParsed {
				return nil, p.newError("multiple code blocks found in system")
			}
			p.nextToken()                    // Consume {
			system.Code = p.parseCodeBlock() // Parse until matching }
			if !p.curTokenIs(token.RBRACE) {
				return nil, p.newError("expected '}' to close code block, got %s", p.curToken.Type)
			}
			p.nextToken() // Consume }
			codeParsed = true
		default:
			return nil, p.newError("unexpected token %s in system body", p.curToken.Type)
		}
	}

	// Expect closing brace for the system
	if !p.curTokenIs(token.RBRACE) {
		return nil, p.newError("expected '}' to close system body, got %s", p.curToken.Type)
	}
	// Note: The final } is consumed by the ParseProgram loop

	return system, nil
}

func (p *Parser) parseParametersBlock() ([]*ast.Parameter, error) {
	params := []*ast.Parameter{}
	p.nextToken() // Consume 'params' identifier

	if !p.curTokenIs(token.LBRACE) {
		return nil, p.newError("expected '{' to start params block, got %s", p.curToken.Type)
	}
	p.nextToken() // Consume {

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		// Similar to parseField, but using ast.Parameter
		if !p.curTokenIs(token.IDENT) {
			return nil, p.newError("expected parameter type, got %s", p.curToken.Type)
		}
		paramType := p.curToken.Literal
		p.nextToken()

		if !p.curTokenIs(token.IDENT) {
			return nil, p.newError("expected parameter name, got %s", p.curToken.Type)
		}
		paramName := p.curToken.Literal
		p.nextToken()

		param := &ast.Parameter{Name: paramName, Type: paramType}

		// Optional default value
		if p.curTokenIs(token.ASSIGN) {
			p.nextToken() // Consume =
			if p.curTokenIs(token.STRING) || p.curTokenIs(token.INT) || p.curTokenIs(token.FLOAT) || p.curTokenIs(token.TRUE) || p.curTokenIs(token.FALSE) {
				param.DefaultValue = p.curToken.Literal
				p.nextToken() // Consume value
			} else {
				return nil, p.newError("expected simple literal for parameter default value, got %s", p.curToken.Type)
			}
		}

		if !p.curTokenIs(token.SEMICOLON) {
			return nil, p.newError("expected ';' after parameter definition, got %s", p.curToken.Type)
		}
		p.nextToken() // Consume ;

		params = append(params, param)
	}

	if !p.curTokenIs(token.RBRACE) {
		return nil, p.newError("expected '}' to close params block, got %s", p.curToken.Type)
	}
	p.nextToken() // Consume }

	return params, nil
}

func (p *Parser) parseCodeBlock() string {
	var code strings.Builder
	for p.curToken.Type != token.RBRACE {
		code.WriteString(p.curToken.Literal)
		if p.peekToken.Type != token.RBRACE {
			code.WriteString(" ")
		}
		p.nextToken()
	}
	return code.String()
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}
	p.errors = append(p.errors, fmt.Sprintf("expected %s, got %s", t, p.peekToken.Type))
	return false
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) newError(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return Error{
		Line:    p.curToken.Line,
		Column:  p.curToken.Column,
		Message: msg,
	}
}

// Add newErrorf helper for errors with specific positions
func (p *Parser) newErrorf(line, column int, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return Error{
		Line:    line,
		Column:  column,
		Message: msg,
	}
}
