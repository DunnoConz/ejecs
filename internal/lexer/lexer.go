package lexer

import (
	"strings"

	"github.com/ejecs/ejecs/internal/token"
)

// No local const needed as all required tokens exist in token package

var keywords = map[string]token.TokenType{
	"component":    token.COMPONENT,
	"system":       token.SYSTEM,
	"relationship": token.RELATIONSHIP,
	"true":         token.TRUE,
	"false":        token.FALSE,
	"nil":          token.NULL,
	"query":        token.QUERY,
	"parameters":   token.IDENT, // Treat as IDENT, requires parser handling
	"frequency":    token.FREQUENCY,
	"priority":     token.PRIORITY,
	"code":         token.CODE,
	"pair":         token.PAIR,
	"table":        token.TABLE,
	// "any" is treated as IDENT by lookupIdent
	// Roblox types are treated as IDENT by lookupIdent
	"Instance": token.IDENT,
	"Vector2":  token.IDENT,
	"Vector3":  token.IDENT,
	"CFrame":   token.IDENT,
	"Color3":   token.IDENT,
	"Enum":     token.IDENT,
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
		l.ch = 0 // NUL character signifies EOF
	} else {
		l.ch = l.input[l.readPosition]
	}

	if l.ch == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	startLine := l.line
	startColumn := l.column

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.EQ, literal, startLine, startColumn)
		} else {
			tok = token.New(token.ASSIGN, string(l.ch), startLine, startColumn)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.NOT_EQ, literal, startLine, startColumn)
		} else {
			tok = token.New(token.BANG, string(l.ch), startLine, startColumn)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.LTE, literal, startLine, startColumn)
		} else {
			tok = token.New(token.LT, string(l.ch), startLine, startColumn)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.GTE, literal, startLine, startColumn)
		} else {
			tok = token.New(token.GT, string(l.ch), startLine, startColumn)
		}
	case '+':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.PLUSEQ, literal, startLine, startColumn)
		} else {
			tok = token.New(token.PLUS, string(l.ch), startLine, startColumn)
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.AND, literal, startLine, startColumn)
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch), startLine, startColumn)
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(token.OR, literal, startLine, startColumn)
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch), startLine, startColumn)
		}
	case '/':
		if l.peekChar() == '/' {
			l.readComment()
			return l.NextToken()
		} else {
			tok = token.New(token.SLASH, string(l.ch), startLine, startColumn)
		}
	case '-':
		tok = token.New(token.MINUS, string(l.ch), startLine, startColumn)
	case '*':
		tok = token.New(token.ASTERISK, string(l.ch), startLine, startColumn)
	case '.':
		tok = token.New(token.DOT, string(l.ch), startLine, startColumn)
	case ',':
		tok = token.New(token.COMMA, string(l.ch), startLine, startColumn)
	case ';':
		tok = token.New(token.SEMICOLON, string(l.ch), startLine, startColumn)
	case ':':
		tok = token.New(token.COLON, string(l.ch), startLine, startColumn)
	case '(':
		tok = token.New(token.LPAREN, string(l.ch), startLine, startColumn)
	case ')':
		tok = token.New(token.RPAREN, string(l.ch), startLine, startColumn)
	case '{':
		tok = token.New(token.LBRACE, string(l.ch), startLine, startColumn)
	case '}':
		tok = token.New(token.RBRACE, string(l.ch), startLine, startColumn)
	case '@':
		tok = token.New(token.AT, string(l.ch), startLine, startColumn)
	case '?':
		tok = token.New(token.QUESTION, string(l.ch), startLine, startColumn)
	case '[':
		tok = token.New(token.LBRACKET, string(l.ch), startLine, startColumn) // Use token.LBRACKET
	case ']':
		tok = token.New(token.RBRACKET, string(l.ch), startLine, startColumn) // Use token.RBRACKET
	case '"', '\'':
		tok.Literal = l.readString(l.ch)
		tok.Type = token.STRING
		tok.Line = startLine
		tok.Column = startColumn
		if l.ch == '"' || l.ch == '\'' { // Check if readString stopped at a quote
			l.readChar() // Consume the closing quote
		}
		return tok
	case 0:
		tok = token.New(token.EOF, "", startLine, startColumn)
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tokType := lookupIdent(literal)
			tok = token.New(tokType, literal, startLine, startColumn)
			return tok
		} else if isDigit(l.ch) {
			literal := l.readNumber()
			var numTokType token.TokenType = token.INT
			if strings.Contains(literal, ".") {
				numTokType = token.FLOAT
			}
			tok = token.New(token.TokenType(numTokType), literal, startLine, startColumn)
			return tok
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch), startLine, startColumn)
		}
	}

	l.readChar()
	return tok
}

func lookupIdent(ident string) token.TokenType {
	// Restore original lookup logic
	if tokType, ok := keywords[ident]; ok {
		return tokType // Returns token.TokenType from map
	}
	// Treat unknown identifiers (like Roblox types not explicitly mapped, 'any', etc.) as IDENT
	return token.IDENT
}

// ... Helper functions (skipWhitespace, readIdentifier, readNumber, readString, readComment, peekChar, isLetter, isDigit) ...

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) || (l.ch == '.' && isDigit(l.peekChar())) { // Allow dot only if followed by digit
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString(quote byte) string {
	var value strings.Builder
	for {
		l.readChar()
		if l.ch == '\\' { // Handle escape sequence
			l.readChar() // Read the character after backslash
			switch l.ch {
			case 'n':
				value.WriteByte('\n')
			case 't':
				value.WriteByte('\t')
			case 'r':
				value.WriteByte('\r')
			case '\'':
				value.WriteByte('\'')
			case '"':
				value.WriteByte('"')
			case '\\':
				value.WriteByte('\\')
			default:
				// Optional: Report error for invalid escape or just include backslash + char
				value.WriteByte('\\')
				value.WriteByte(l.ch)
			}
		} else if l.ch == quote || l.ch == 0 { // End on matching quote or EOF
			break
		} else {
			value.WriteByte(l.ch) // Append regular character
		}
	}
	return value.String()
}

func (l *Lexer) readComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
