package lexer

import (
	"testing"

	"github.com/ejecs/ejecs/internal/token"
)

func TestNextTokenBasic(t *testing.T) {
	input := `component Position {
		x: float
		y: float
	}

	system Movement {
		query: Position
		frequency: 60
		priority: 1
		code: {
			// Update position
		}
	}`

	basicTests := []struct {
		expectedTokenType token.TokenType
		expectedTokenLit  string
	}{
		{token.COMPONENT, "component"},
		{token.IDENT, "Position"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.COLON, ":"},
		{token.IDENT, "float"},
		{token.IDENT, "y"},
		{token.COLON, ":"},
		{token.IDENT, "float"},
		{token.RBRACE, "}"},
		{token.SYSTEM, "system"},
		{token.IDENT, "Movement"},
		{token.LBRACE, "{"},
		{token.QUERY, "query"},
		{token.COLON, ":"},
		{token.IDENT, "Position"},
		{token.FREQUENCY, "frequency"},
		{token.COLON, ":"},
		{token.INT, "60"},
		{token.PRIORITY, "priority"},
		{token.COLON, ":"},
		{token.INT, "1"},
		{token.CODE, "code"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range basicTests {
		tok := l.NextToken()

		if tok.Type != tt.expectedTokenType {
			t.Fatalf("basicTests[%d] - tokentype wrong. expected=%q, got=%q (%s)",
				i, tt.expectedTokenType, tok.Type, tok.Literal)
		}

		if tok.Literal != tt.expectedTokenLit {
			t.Fatalf("basicTests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedTokenLit, tok.Literal)
		}
	}
}

func TestNextTokenNumeric(t *testing.T) {
	input := `123 45.67 -89 -12.34`

	numericTests := []struct {
		expectedNumType token.TokenType
		expectedNumLit  string
	}{
		{token.INT, "123"},
		{token.FLOAT, "45.67"},
		{token.MINUS, "-"},
		{token.INT, "89"},
		{token.MINUS, "-"},
		{token.FLOAT, "12.34"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range numericTests {
		tok := l.NextToken()

		if tok.Type != tt.expectedNumType {
			t.Fatalf("numericTests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedNumType, tok.Type)
		}

		if tok.Literal != tt.expectedNumLit {
			t.Fatalf("numericTests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedNumLit, tok.Literal)
		}
	}
}

func TestNextTokenString(t *testing.T) {
	input := `"hello" 'world' "escaped \"quote\""`

	stringTests := []struct {
		expectedStrType token.TokenType
		expectedStrLit  string
	}{
		{token.STRING, "hello"},
		{token.STRING, "world"},
		{token.STRING, "escaped \"quote\""},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range stringTests {
		tok := l.NextToken()

		if tok.Type != tt.expectedStrType {
			t.Fatalf("stringTests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedStrType, tok.Type)
		}

		if tok.Literal != tt.expectedStrLit {
			t.Fatalf("stringTests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedStrLit, tok.Literal)
		}
	}
}

func TestNextToken_Comments(t *testing.T) {
	input := `// This is a comment
component Position { // Inline comment
	x: number; // Another comment
}`

	commentTests := []struct {
		expectedTokenType token.TokenType
		expectedTokenLit  string
	}{
		{token.COMPONENT, "component"},
		{token.IDENT, "Position"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.COLON, ":"},
		{token.IDENT, "number"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range commentTests {
		tok := l.NextToken()

		if tok.Type != tt.expectedTokenType {
			t.Fatalf("commentTests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedTokenType, tok.Type)
		}

		if tok.Literal != tt.expectedTokenLit {
			t.Fatalf("commentTests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedTokenLit, tok.Literal)
		}
	}
}

func TestNextToken_Operators(t *testing.T) {
	input := `+ - * / = == != < <= > >= && ||`

	operatorTests := []struct {
		expectedTokenType token.TokenType
		expectedTokenLit  string
	}{
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.ASSIGN, "="},
		{token.EQ, "=="},
		{token.NOT_EQ, "!="},
		{token.LT, "<"},
		{token.LTE, "<="},
		{token.GT, ">"},
		{token.GTE, ">="},
		{token.AND, "&&"},
		{token.OR, "||"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range operatorTests {
		tok := l.NextToken()

		if tok.Type != tt.expectedTokenType {
			t.Fatalf("operatorTests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedTokenType, tok.Type)
		}

		if tok.Literal != tt.expectedTokenLit {
			t.Fatalf("operatorTests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedTokenLit, tok.Literal)
		}
	}
}

func TestNextToken_IntFloat(t *testing.T) {
	input := `123 0 3.14`

	tests := []struct {
		expectedTokenType token.TokenType
		expectedTokenLit  string
	}{
		{token.INT, "123"},
		{token.INT, "0"},
		{token.FLOAT, "3.14"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedTokenType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedTokenType, tok.Type)
		}

		if tok.Literal != tt.expectedTokenLit {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedTokenLit, tok.Literal)
		}
	}
}

func TestNextToken_Keywords(t *testing.T) {
	input := `component system relationship true false nil query parameters frequency priority code pair table`
	tests := []struct {
		expectedTokenType token.TokenType
		expectedTokenLit  string
	}{
		{token.COMPONENT, "component"},
		{token.SYSTEM, "system"},
		{token.RELATIONSHIP, "relationship"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.NULL, "nil"},
		{token.QUERY, "query"},
		{token.IDENT, "parameters"}, // parameters is IDENT
		{token.FREQUENCY, "frequency"},
		{token.PRIORITY, "priority"},
		{token.CODE, "code"},
		{token.PAIR, "pair"},
		{token.TABLE, "table"}, // Added table test case here
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedTokenType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedTokenType, tok.Type)
		}

		if tok.Literal != tt.expectedTokenLit {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedTokenLit, tok.Literal)
		}
	}
}

func TestNextToken_TableType(t *testing.T) {
	input := `table<string, any>`
	tests := []struct {
		expectedTokenType token.TokenType
		expectedTokenLit  string
	}{
		{token.TABLE, "table"},
		{token.LT, "<"},
		{token.IDENT, "string"}, // Treat string type name as IDENT for now
		{token.COMMA, ","},
		{token.IDENT, "any"},
		{token.GT, ">"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedTokenType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedTokenType, tok.Type)
		}

		if tok.Literal != tt.expectedTokenLit {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedTokenLit, tok.Literal)
		}
	}
}
