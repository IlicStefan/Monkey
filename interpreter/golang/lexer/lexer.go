package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	character    byte // current char under examination
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var aToken token.Token

	switch lexer.character {
	case '=':
		aToken = newToken(token.ASSIGNMENT, lexer.character)
	case ';':
		aToken = newToken(token.SEMICOLON, lexer.character)
	case '(':
		aToken= newToken(token.LEFT_PARENTHESES, lexer.character)
	case ')':
		aToken = newToken(token.RIGHT_PARENTHESES, lexer.character)
	case ',':
		aToken = newToken(token.COMMA, lexer.character)
	case '+':
		aToken= newToken(token.PLUS, lexer.character)
	case '{':
		aToken = newToken(token.LEFT_CURLY_BRACKET, lexer.character)
	case '}':
		aToken = newToken(token.RIGHT_CURLY_BRACKET, lexer.character)
	case 0:
		aToken.Literal = ""
		aToken.Type = token.EOF
	}

	lexer.readChar()
	return aToken
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(character),
	}
}
