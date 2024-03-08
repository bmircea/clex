
package token 

type TokenType string


type Token struct {
    Type TokenType
    Literal string
}


const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
    
    // Identifiers
    IDENTIFIER = "IDENTIFIER"
    INT = "INT"
    
    // Operators
    EQUAL = "="
    PLUS = "+"


    // Delimiters
    COMMA = ","
    SEMICOLON = ";"

    LPAR = "("
    RPAR = ")"
    LSQUIRLY = "{"
    RSQUIRLY = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    
)

var keywords = map[string] TokenType{
    "fn": FUNCTION,
}

func LookupIdentifier(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }

    return IDENTIFIER
}


