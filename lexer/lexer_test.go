package lexer

import (
	"clex/token"
	"testing"
)

func TestNextToken(t *testing.T) {
    input := `=+(){},;
    !-/*5;
    5 < 10 > 5;return true;return false; if;
    10 == 10;
    10 != 9;
    3.75;
    `
    

    tests := []struct {
        expectedT token.TokenType 
        expectedL string
        expectedLine int
    } {
        {token.EQUAL, "=", 1},
        {token.PLUS, "+", 1},
        {token.LPAR, "(", 1},
        {token.RPAR, ")", 1},
        {token.LSQUIRLY, "{", 1},
        {token.RSQUIRLY, "}", 1},
        {token.COMMA, ",", 1},
        {token.SEMICOLON, ";", 1},
        {token.BANG, "!", 2},
        {token.MINUS, "-", 2},
        {token.SLASH, "/", 2},
        {token.ASTERISK, "*", 2},
        {token.INT, "5", 2},
        {token.SEMICOLON, ";", 2},
        {token.INT, "5", 3},
        {token.LESS_THAN, "<", 3},
        {token.INT, "10", 3},
        {token.GREATER_THAN, ">", 3},
        {token.INT, "5", 3},
        {token.SEMICOLON, ";", 3},
        {token.RETURN, "return", 3},
        {token.TRUE, "true", 3},
        {token.SEMICOLON, ";", 3},
        {token.RETURN, "return", 3},
        {token.FALSE, "false", 3},
        {token.SEMICOLON, ";", 3},
        {token.IF, "if", 3},
        {token.SEMICOLON, ";", 3},
        {token.INT, "10", 4},
        {token.IS_EQ, "==", 4},
        {token.INT, "10", 4},
        {token.SEMICOLON, ";", 4},
        {token.INT, "10", 5},
        {token.IS_NOT_EQ, "!=", 5},
        {token.INT, "9", 5},
        {token.SEMICOLON, ";", 5},
        {token.FLOATING, "3.75", 6},
        {token.SEMICOLON, ";", 6},
    }

    l := New(input)

    for i, tt := range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedT {
            t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedT, tok.Type)
        }

        if tok.Literal != tt.expectedL {
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedL, tok.Literal)
        }

        if tok.Line != tt.expectedLine {
        
            t.Logf("tests[%d] - line number wrong. expected=%q, got %q", i, tt.expectedLine, tok.Line)
        }
    }
}
