package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `component Position {
		x: number;
		y: number;
	}

	@replicated
	component Player {
		name: string?;
		health: number;
	}

	relationship ChildOf {
		child: Entity;
		parent: Entity;
	}

	system Movement {
		query: (Position, pair(ChildOf, *), Velocity);
		frequency: 60hz;
		priority: 1;
		run(pos, vel) {
			pos.x = pos.x + vel.x;
			pos.y = pos.y + vel.y;
		}
	}`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{COMPONENT, "component"},
		{IDENT, "Position"},
		{LBRACE, "{"},
		{IDENT, "x"},
		{COLON, ":"},
		{IDENT, "number"},
		{SEMICOLON, ";"},
		{IDENT, "y"},
		{COLON, ":"},
		{IDENT, "number"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},

		{AT, "@"},
		{IDENT, "replicated"},
		{COMPONENT, "component"},
		{IDENT, "Player"},
		{LBRACE, "{"},
		{IDENT, "name"},
		{COLON, ":"},
		{IDENT, "string"},
		{QUESTION, "?"},
		{SEMICOLON, ";"},
		{IDENT, "health"},
		{COLON, ":"},
		{IDENT, "number"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},

		{RELATIONSHIP, "relationship"},
		{IDENT, "ChildOf"},
		{LBRACE, "{"},
		{IDENT, "child"},
		{COLON, ":"},
		{IDENT, "Entity"},
		{SEMICOLON, ";"},
		{IDENT, "parent"},
		{COLON, ":"},
		{IDENT, "Entity"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},

		{SYSTEM, "system"},
		{IDENT, "Movement"},
		{LBRACE, "{"},
		{QUERY, "query"},
		{COLON, ":"},
		{LPAREN, "("},
		{IDENT, "Position"},
		{COMMA, ","},
		{PAIR, "pair"},
		{LPAREN, "("},
		{IDENT, "ChildOf"},
		{COMMA, ","},
		{ASTERISK, "*"},
		{RPAREN, ")"},
		{COMMA, ","},
		{IDENT, "Velocity"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},
		{IDENT, "frequency"},
		{COLON, ":"},
		{IDENT, "60hz"},
		{SEMICOLON, ";"},
		{IDENT, "priority"},
		{COLON, ":"},
		{IDENT, "1"},
		{SEMICOLON, ";"},
		{RUN, "run"},
		{LPAREN, "("},
		{IDENT, "pos"},
		{COMMA, ","},
		{IDENT, "vel"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{IDENT, "pos"},
		{DOT, "."},
		{IDENT, "x"},
		{ASSIGN, "="},
		{IDENT, "pos"},
		{DOT, "."},
		{IDENT, "x"},
		{PLUS, "+"},
		{IDENT, "vel"},
		{DOT, "."},
		{IDENT, "x"},
		{SEMICOLON, ";"},
		{IDENT, "pos"},
		{DOT, "."},
		{IDENT, "y"},
		{ASSIGN, "="},
		{IDENT, "pos"},
		{DOT, "."},
		{IDENT, "y"},
		{PLUS, "+"},
		{IDENT, "vel"},
		{DOT, "."},
		{IDENT, "y"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{RBRACE, "}"},
		{EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_Numbers(t *testing.T) {
	input := `123 45.67 .89`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{NUMBER, "123"},
		{NUMBER, "45.67"},
		{DOT, "."},
		{NUMBER, "89"},
		{EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_Comments(t *testing.T) {
	input := `// This is a comment
component Position { // Inline comment
	x: number; // Another comment
}`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{COMPONENT, "component"},
		{IDENT, "Position"},
		{LBRACE, "{"},
		{IDENT, "x"},
		{COLON, ":"},
		{IDENT, "number"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_Operators(t *testing.T) {
	input := `+ - * / = == != < <= > >= && ||`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{PLUS, "+"},
		{MINUS, "-"},
		{ASTERISK, "*"},
		{SLASH, "/"},
		{ASSIGN, "="},
		{EQ, "=="},
		{NOT_EQ, "!="},
		{LT, "<"},
		{LTE, "<="},
		{GT, ">"},
		{GTE, ">="},
		{AND, "&&"},
		{OR, "||"},
		{EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
