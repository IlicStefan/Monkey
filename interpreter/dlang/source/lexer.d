module lexer;

import token;

class Lexer {
    private string input;
    private int position; // current position in input (points to current char)
    private int readPosition; // current reading position in input (after current char)
    private char character; // current char under examination

    public this(string input) {
        this.input = input;
    }

    private void readChar() {
        if(readPosition >= input.length) {
            character = '\0';
        }
        else {
            character = input[readPosition];
        }
        position = readPosition;
        readPosition++;
    }

    public Token nextToken() {
        readChar();

        Token token;

        if(character == '\0') {
            token.type = TokenType.EOF;
            return token;
        }

        token.literal = [character];

        switch(character) {
        case '=':
            token.type = TokenType.ASSIGNMENT;
            break;
        case ';':
            token.type = TokenType.SEMICOLON;
            break;
        case '(':
            token.type = TokenType.LEFT_PARENTHESES;
            break;
        case ')':
            token.type = TokenType.RIGHT_PARENTHESES;
            break;
        case ',':
            token.type = TokenType.COMMA;
            break;
        case '+':
            token.type = TokenType.PLUS;
            break;
        case '{':
            token.type = TokenType.LEFT_CURLY_BRACKET;
            break;
        case '}':
            token.type = TokenType.RIGHT_CURLY_BRACKET;
            break;
        default:
            throw new Exception("Lexical error: Invalid character " ~ character);
        }

        return token;
    }
}

unittest {
    string input = "=+(){},;";

    Lexer lexer = new Lexer(input);

    Token[] tests = [
        {TokenType.ASSIGNMENT, "="},
        {TokenType.PLUS, "+"},
        {TokenType.LEFT_PARENTHESES, "("},
        {TokenType.RIGHT_PARENTHESES, ")"},
        {TokenType.LEFT_CURLY_BRACKET, "{"},
        {TokenType.RIGHT_CURLY_BRACKET, "}"},
        {TokenType.COMMA, ","},
        {TokenType.SEMICOLON, ";"},
        {TokenType.EOF, ""},
    ];

    foreach(i, expectedToken; tests) {
        Token token = lexer.nextToken();

        assert(token.type == expectedToken.type);
        assert(token.literal == expectedToken.literal);
    }
}