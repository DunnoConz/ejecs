package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/ejecs/ejecs/internal/ast"
	"github.com/ejecs/ejecs/internal/lexer"
	"github.com/ejecs/ejecs/internal/token"
)

// Precedence levels for operators (add more as needed)
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index]
	DOT         // table.field
)

// Operator precedence map (add more operators)
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.LTE:      LESSGREATER,
	token.GTE:      LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL, // For function calls
	token.DOT:      DOT,  // For member access like CFrame.new
}

// Pratt parser function types
type (
	prefixParseFn func() (ast.Expression, error)
	infixParseFn  func(ast.Expression) (ast.Expression, error)
)

// Parser represents a JECS parser
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
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

	// Initialize parsing function maps
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseNumberLiteral)
	p.registerPrefix(token.FLOAT, p.parseNumberLiteral)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.TRUE, p.parseBooleanLiteral)
	p.registerPrefix(token.FALSE, p.parseBooleanLiteral)
	p.registerPrefix(token.LBRACE, p.parseTableConstructor)  // For table constructors {}
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression) // For ( expression )
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.LPAREN, p.parseCallExpression)      // For func()
	p.registerInfix(token.DOT, p.parseMemberAccessExpression) // For table.field or CFrame.new
	// Add other infix operators (+, -, *, /, ==, <, etc.) if needed

	// Read two tokens, so curToken and peekToken are both set.
	p.nextToken()
	p.nextToken()

	return p
}

// Helper registration functions
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
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
		case token.RELATIONSHIP, token.AT:
			if p.curTokenIs(token.AT) && !p.peekTokenIs(token.IDENT) {
				return nil, p.newError("expected identifier after @ for relationship type, got %s", p.peekToken.Type)
			}
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
	var defaultValueExpr ast.Expression
	var err error

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

	// --- Optional Default Value ---
	if p.curTokenIs(token.ASSIGN) {
		p.nextToken() // Consume '='
		defaultValueExpr, err = p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		field.DefaultValue = defaultValueExpr
		// parseExpression leaves curToken on the last token of the expression.
	}

	// --- Find the semicolon ---
	// After type/name and optional default value, skip until we find the semicolon.
	// This simplifies handling the end of potentially complex default value expressions.
	for !p.curTokenIs(token.SEMICOLON) && !p.curTokenIs(token.EOF) {
		// Maybe add a check here to prevent infinite loops if semicolon is missing on the line?
		p.nextToken()
	}

	if !p.curTokenIs(token.SEMICOLON) {
		// If we hit EOF without finding semicolon, report error at the field name's position (approx)
		// TODO: Improve error position reporting
		return nil, p.newError("missing ';' after field definition for field '%s'", field.Name)
	}

	p.nextToken() // Consume the SEMICOLON.

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
			if system.Frequency != nil {
				return nil, p.newError("duplicate frequency definition")
			}
			p.nextToken() // Consume 'frequency'
			if !p.curTokenIs(token.COLON) {
				return nil, p.newError("expected ':' after frequency, got %s", p.curToken.Type)
			}
			p.nextToken() // Consume ':'

			freqExpr, err := p.parseExpression(LOWEST)
			if err != nil {
				return nil, err
			}
			system.Frequency = freqExpr
			// parseExpression leaves curToken on the last token of the expression (e.g., the closing ')' of fixed(60) )
			p.nextToken() // Consume the last token of the frequency expression

		case token.PRIORITY:
			if system.Priority != nil {
				return nil, p.newError("duplicate priority definition")
			}
			p.nextToken() // Consume 'priority'
			if !p.curTokenIs(token.COLON) {
				return nil, p.newError("expected ':' after priority, got %s", p.curToken.Type)
			}
			p.nextToken() // Consume ':'

			prioExpr, err := p.parseExpression(LOWEST)
			if err != nil {
				return nil, err
			}
			system.Priority = prioExpr
			// parseExpression leaves curToken on the last token of the expression (e.g., the INT 100)
			p.nextToken() // Consume the last token of the priority expression

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
			defaultValueExpr, err := p.parseExpression(LOWEST)
			if err != nil {
				return nil, err
			}
			param.DefaultValue = defaultValueExpr // Assign expression node
			// parseExpression leaves curToken on last token of expr
			p.nextToken() // Consume last token of default value expr
		}

		if !p.curTokenIs(token.SEMICOLON) {
			return nil, p.newError("expected ';' after parameter definition, got %s (%q)", p.curToken.Type, p.curToken.Literal)
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

// Reverted parseCodeBlock: Reconstruct with heuristic spacing (imperfect)
func (p *Parser) parseCodeBlock() string {
	var code strings.Builder
	startLine := p.curToken.Line // Line where the opening { was - FOR ERROR REPORTING ONLY

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		code.WriteString(p.curToken.Literal)

		// Add space only if syntactically likely needed
		if shouldAddSpace(p.curToken, p.peekToken) {
			code.WriteString(" ")
		}
		p.nextToken()
	}

	if p.curTokenIs(token.EOF) {
		// Error: reached EOF before finding matching RBRACE
		p.newErrorf(startLine, 0, "unterminated code block starting on line %d", startLine)
		// Return what we have, maybe?
	}

	return code.String()
}

// shouldAddSpace heuristic: determines if a space should be added AFTER cur and BEFORE peek.
func shouldAddSpace(cur, peek token.Token) bool {
	// Never add space before the final closing brace of the code block
	if peek.Type == token.RBRACE {
		return false
	}

	// Never add space before punctuation that doesn't need leading space
	switch peek.Type {
	case token.DOT, token.COMMA, token.SEMICOLON, token.COLON,
		token.LPAREN, token.RPAREN, token.LBRACKET, token.RBRACKET, token.LBRACE: // Don't add space before {, but maybe after
		return false
	}

	// Never add space AFTER punctuation that doesn't need trailing space
	switch cur.Type {
	case token.DOT, token.LPAREN, token.LBRACKET, token.LBRACE, token.AT: // Don't add space after { or (
		return false
	}

	// General rule: Add space between two "word-like" tokens
	curIsWord := isWordToken(cur)
	peekIsWord := isWordToken(peek)

	return curIsWord && peekIsWord
}

// isWordToken checks if a token is typically treated as a word requiring spacing.
func isWordToken(tok token.Token) bool {
	switch tok.Type {
	case token.IDENT, token.INT, token.FLOAT, token.STRING, token.TRUE, token.FALSE,
		token.COMPONENT, token.SYSTEM, token.RELATIONSHIP, token.QUERY, // Removed PARAMS
		token.FREQUENCY, token.PRIORITY, token.RETURN, token.FUNCTION, // More keywords if needed
		token.IF, token.ELSE, token.FOR, token.WHILE: // Removed DO, END, LOCAL
		return true
	default:
		return false
	}
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

// --- Expression Parsing ---

func (p *Parser) parseExpression(precedence int) (ast.Expression, error) {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil, p.newError("no prefix parse function for %s found", p.curToken.Type)
	}
	leftExp, err := prefix()
	if err != nil {
		return nil, err
	}

	// Loop for infix operators
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp, nil // No infix operator found or lower precedence
		}

		p.nextToken() // Consume the infix operator
		leftExp, err = infix(leftExp)
		if err != nil {
			return nil, err
		}
	}

	return leftExp, nil
}

// Get precedence of the *next* token
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

// Get precedence of the *current* token
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

// Placeholder parsing functions
func (p *Parser) parseIdentifier() (ast.Expression, error) {
	return &ast.Identifier{Value: p.curToken.Literal}, nil
}

func (p *Parser) parseNumberLiteral() (ast.Expression, error) {
	return &ast.NumberLiteral{Value: p.curToken.Literal}, nil
}

func (p *Parser) parseStringLiteral() (ast.Expression, error) {
	return &ast.StringLiteral{Value: p.curToken.Literal}, nil
}

func (p *Parser) parseBooleanLiteral() (ast.Expression, error) {
	return &ast.BooleanLiteral{Value: p.curTokenIs(token.TRUE)}, nil
}

func (p *Parser) parseGroupedExpression() (ast.Expression, error) {
	p.nextToken() // Consume '('
	exp, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	if !p.expectPeek(token.RPAREN) { // Consume ')'
		return nil, p.newError("expected ')' after grouped expression")
	}
	return exp, nil
}

func (p *Parser) parseTableConstructor() (ast.Expression, error) {
	table := &ast.TableConstructor{Fields: []*ast.TableField{}}
	startLine, startCol := p.curToken.Line, p.curToken.Column // For error reporting

	// Handle empty table {}
	if p.peekTokenIs(token.RBRACE) {
		p.nextToken() // Consume {
		p.nextToken() // Consume }
		return table, nil
	}

	p.nextToken() // Consume {

	// Parse first field
	keyExpr, valueExpr, err := p.parseTableField()
	if err != nil {
		return nil, err
	}
	table.Fields = append(table.Fields, &ast.TableField{Key: keyExpr, Value: valueExpr})

	// Parse subsequent fields (comma-separated)
	for p.curTokenIs(token.COMMA) {
		p.nextToken() // Consume ,

		// Allow trailing comma
		if p.curTokenIs(token.RBRACE) {
			break
		}

		keyExpr, valueExpr, err := p.parseTableField()
		if err != nil {
			return nil, err
		}
		table.Fields = append(table.Fields, &ast.TableField{Key: keyExpr, Value: valueExpr})
	}

	// Expect closing brace
	if !p.curTokenIs(token.RBRACE) {
		return nil, p.newErrorf(startLine, startCol, "expected '}' or ',' in table constructor, got %s", p.curToken.Type)
	}
	p.nextToken() // Consume }

	return table, nil
}

// Parses a single field inside a table constructor: [expr]=expr, ident=expr, or just expr
func (p *Parser) parseTableField() (ast.Expression, ast.Expression, error) {
	var key, value ast.Expression
	var err error

	// Check for different key syntaxes or just a value
	if p.curTokenIs(token.LBRACKET) {
		// Key is an expression: [expr] = value
		p.nextToken() // Consume [
		key, err = p.parseExpression(LOWEST)
		if err != nil {
			return nil, nil, err
		}
		if !p.expectPeek(token.RBRACKET) {
			return nil, nil, p.newError("expected ']' after table key expression")
		}
		if !p.expectPeek(token.ASSIGN) {
			return nil, nil, p.newError("expected '=' after table key expression")
		}
		value, err = p.parseExpression(LOWEST)
		if err != nil {
			return nil, nil, err
		}
	} else if p.curTokenIs(token.IDENT) && p.peekTokenIs(token.ASSIGN) {
		// Key is an identifier: key = value
		key = &ast.Identifier{Value: p.curToken.Literal}
		p.nextToken() // Consume ident
		p.nextToken() // Consume =
		value, err = p.parseExpression(LOWEST)
		if err != nil {
			return nil, nil, err
		}
	} else {
		// Key is nil, just a value (array-like table)
		key = nil
		value, err = p.parseExpression(LOWEST)
		if err != nil {
			return nil, nil, err
		}
	}

	// Consume the last token of the value expression before returning
	p.nextToken()

	return key, value, nil
}

func (p *Parser) parseCallExpression(function ast.Expression) (ast.Expression, error) {
	call := &ast.CallExpression{Function: function}
	var err error
	call.Arguments, err = p.parseExpressionList(token.RPAREN)
	if err != nil {
		return nil, err
	}
	return call, nil
}

// Helper to parse comma-separated expressions until an end token
func (p *Parser) parseExpressionList(end token.TokenType) ([]ast.Expression, error) {
	list := []ast.Expression{}

	if p.peekTokenIs(end) { // Handle empty list like func()
		p.nextToken() // Consume the end token
		return list, nil
	}

	p.nextToken() // Consume LPAREN or COMMA
	exp, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	list = append(list, exp)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // Consume ,
		p.nextToken() // Move to the start of the next expression
		exp, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		list = append(list, exp)
	}

	if !p.expectPeek(end) { // Consume the end token
		return nil, p.newError("expected '%s' to end expression list", end)
	}

	return list, nil
}

// Parses member access like table.field or CFrame.new
func (p *Parser) parseMemberAccessExpression(left ast.Expression) (ast.Expression, error) {
	if !p.curTokenIs(token.DOT) {
		return nil, p.newError("expected '.' for member access")
	}
	p.nextToken() // Consume '.'

	if !p.curTokenIs(token.IDENT) {
		return nil, p.newError("expected identifier after '.'")
	}

	member := &ast.Identifier{Value: p.curToken.Literal}

	exp := &ast.MemberAccessExpression{
		Object:     left,
		MemberName: member,
	}

	// Do not consume the member identifier here;
	// the main parseExpression loop will handle the next token.

	return exp, nil
}

// Parsing function for prefix operators like - or !
func (p *Parser) parsePrefixExpression() (ast.Expression, error) {
	expression := &ast.PrefixExpression{
		Operator: p.curToken.Literal,
	}
	p.nextToken() // Consume the operator token (e.g., '-')
	var err error
	expression.Right, err = p.parseExpression(PREFIX) // Parse the operand with PREFIX precedence
	if err != nil {
		return nil, err
	}
	return expression, nil
}
