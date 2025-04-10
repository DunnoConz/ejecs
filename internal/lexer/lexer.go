package lexer

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	NUMBER = "NUMBER"
	STRING = "STRING"

	// Operators
	COLON     = ":"
	SEMICOLON = ";"
	COMMA     = ","
	ASTERISK  = "*"
	DOT       = "."
	PLUS      = "+"
	MINUS     = "-"
	SLASH     = "/"
	ASSIGN    = "="
	PLUSEQ    = "+="
	EQ        = "=="
	NOT_EQ    = "!="
	LT        = "<"
	LTE       = "<="
	GT        = ">"
	GTE       = ">="
	AND       = "&&"
	OR        = "||"

	// Delimiters
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Special characters
	AT       = "@"
	QUESTION = "?"

	// Keywords
	COMPONENT    = "component"
	RELATIONSHIP = "relationship"
	SYSTEM       = "system"
	QUERY        = "query"
	RUN          = "run"
	PAIR         = "pair"
	GET_TARGET   = "getTarget"
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
	column       int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 1}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

	if l.ch == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	tok.Line = l.line
	tok.Column = l.column

	switch l.ch {
	case ':':
		tok = newToken(COLON, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case '*':
		tok = newToken(ASTERISK, l.ch)
	case '.':
		if isDigit(l.peekChar()) {
			l.readChar() // consume the dot
			tok.Literal = "." + l.readNumber()
			tok.Type = NUMBER
			return tok
		}
		tok = newToken(DOT, l.ch)
	case '+':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: PLUSEQ, Literal: "+="}
		} else {
			tok = newToken(PLUS, l.ch)
		}
	case '-':
		tok = newToken(MINUS, l.ch)
	case '/':
		if l.peekChar() == '/' {
			// Skip comment
			l.skipComment()
			return l.NextToken()
		}
		tok = newToken(SLASH, l.ch)
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: EQ, Literal: "=="}
		} else {
			tok = newToken(ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: NOT_EQ, Literal: "!="}
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: LTE, Literal: "<="}
		} else {
			tok = newToken(LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: GTE, Literal: ">="}
		} else {
			tok = newToken(GT, l.ch)
		}
	case '&':
		if l.peekChar() == '&' {
			l.readChar()
			tok = Token{Type: AND, Literal: "&&"}
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	case '|':
		if l.peekChar() == '|' {
			l.readChar()
			tok = Token{Type: OR, Literal: "||"}
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	case '@':
		tok = newToken(AT, l.ch)
	case '?':
		tok = newToken(QUESTION, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = l.lookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			literal := l.readNumber()
			// Check if this is part of an identifier (like 60hz)
			if isLetter(l.ch) {
				literal += l.readIdentifier()
				tok.Type = IDENT
			} else {
				tok.Type = NUMBER
			}
			tok.Literal = literal
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	// First character must be a letter or underscore
	if isLetter(l.ch) {
		l.readChar()
		// Subsequent characters can be letters, digits, or special cases like 'hz'
		for isLetter(l.ch) || isDigit(l.ch) {
			l.readChar()
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	if l.ch == '.' {
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
	}
	// If the next character is a letter, this is part of an identifier
	if isLetter(l.ch) {
		return l.input[position:l.position]
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) lookupIdent(ident string) TokenType {
	switch ident {
	case "component":
		return COMPONENT
	case "relationship":
		return RELATIONSHIP
	case "system":
		return SYSTEM
	case "query":
		return QUERY
	case "run":
		return RUN
	case "pair":
		return PAIR
	case "getTarget":
		return GET_TARGET
	default:
		return IDENT
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}
