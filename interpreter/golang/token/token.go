package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILEGAL = "ILEGAL"
	EOF    = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	ASSIGNMENT = "="
	PLUS       = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LEFT_PARENTHESES    = "("
	RIGHT_PARENTHESES   = ")"
	LEFT_CURLY_BRACKET  = "{"
	RIGHT_CURLY_BRACKET = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
