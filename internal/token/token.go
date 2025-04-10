package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"

	// Basic types
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"
	BOOL   = "BOOL"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
	LTE      = "<="
	GTE      = ">="
	PLUSEQ   = "+="
	AND      = "&&"
	OR       = "||"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	DOT       = "."
	AT        = "@"
	QUESTION  = "?"

	// Keywords
	COMPONENT    = "component"
	RELATIONSHIP = "relationship"
	SYSTEM       = "system"
	QUERY        = "query"
	RUN          = "run"
	PAIR         = "pair"
	GET_TARGET   = "getTarget"
	USING        = "using"
	FREQUENCY    = "frequency"
	PRIORITY     = "priority"
	CODE         = "code"
	FUNCTION     = "function"
	LET          = "let"
	TRUE         = "true"
	FALSE        = "false"
	IF           = "if"
	ELSE         = "else"
	RETURN       = "return"
	FOR          = "for"
	IN           = "in"
	WHILE        = "while"
	BREAK        = "break"
	CONTINUE     = "continue"
	NULL         = "null"
	TABLE        = "table"
)

// Complex types supported by the language
type ComplexType string

const (
	Vector2        ComplexType = "Vector2"
	Vector3        ComplexType = "Vector3"
	CFrame         ComplexType = "CFrame"
	Color3         ComplexType = "Color3"
	ColorSequence  ComplexType = "ColorSequence"
	NumberRange    ComplexType = "NumberRange"
	NumberSequence ComplexType = "NumberSequence"
	UDim           ComplexType = "UDim"
	UDim2          ComplexType = "UDim2"
	Ray            ComplexType = "Ray"
	Region3        ComplexType = "Region3"
	Region3Int16   ComplexType = "Region3Int16"
	Rect           ComplexType = "Rect"
	Instance       ComplexType = "Instance"
	EnumItem       ComplexType = "EnumItem"
	BrickColor     ComplexType = "BrickColor"
)

// IsComplexType checks if a string represents a complex type
func IsComplexType(s string) bool {
	_, ok := map[string]ComplexType{
		"Vector2":        Vector2,
		"Vector3":        Vector3,
		"CFrame":         CFrame,
		"Color3":         Color3,
		"ColorSequence":  ColorSequence,
		"NumberRange":    NumberRange,
		"NumberSequence": NumberSequence,
		"UDim":           UDim,
		"UDim2":          UDim2,
		"Ray":            Ray,
		"Region3":        Region3,
		"Region3Int16":   Region3Int16,
		"Rect":           Rect,
		"Instance":       Instance,
		"EnumItem":       EnumItem,
		"BrickColor":     BrickColor,
	}[s]
	return ok
}

// IsKeyword checks if a string is a language keyword
func IsKeyword(s string) bool {
	switch s {
	case "component", "relationship", "system", "query",
		"run", "pair", "getTarget", "using", "code",
		"function", "let", "true", "false", "if",
		"else", "return", "for", "in", "while",
		"break", "continue", "null":
		return true
	}
	return false
}

// New creates a new Token
func New(tokenType TokenType, literal string, line, column int) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
		Line:    line,
		Column:  column,
	}
}
