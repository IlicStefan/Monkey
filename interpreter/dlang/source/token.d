module token;

struct Token {
    TokenType type;
    string literal;
}

enum TokenType {
    ILEGAL = "ILEGAL",
    EOF    = "EOF",

    IDENTIFIER = "IDENTIFIER",
    INT        = "INT",

    ASSIGNMENT = "=",
    PLUS       = "+",

    COMMA     = ",",
    SEMICOLON = ";",

    LEFT_PARENTHESES    = "(",
    RIGHT_PARENTHESES   = ")",
    LEFT_CURLY_BRACKET  = "{",
    RIGHT_CURLY_BRACKET = "}",

    FUNCTION = "FUNCTION",
    LET      = "LET",
}