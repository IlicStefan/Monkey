package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGNMENT, "="},
		{token.PLUS, "+"},
		{token.LEFT_PARENTHESES, "("},
		{token.RIGHT_PARENTHESES, ")"},
		{token.LEFT_CURLY_BRACKET, "{"},
		{token.RIGHT_CURLY_BRACKET, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tokenType := range tests {
		token := lexer.NextToken()

		if token.Type != tokenType.expectedType {
			t.Fatalf("tests[%d] - token type wrong. Expected=%q, got %q",
				i, tokenType.expectedType, token.Type)
		}

		if token.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected=%q, got=%q",
				i, tokenType.expectedLiteral, token.Type)
		}
	}
}
