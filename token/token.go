
package token 

type TokenType string


type Token struct {
    Type TokenType
    Literal string
    Line int
}


const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
    
    COMMENT = "COMMENT"
    // Identifiers
    IDENTIFIER = "IDENTIFIER"
    INT = "INT"
    FLOATING = "FLOATING"
    CHAR = "CHAR"
    
    // Operators
    EQUAL = "="
    PLUS = "+"
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"
    LINE_COMMENT = "//"
    MODULO = "%"
    INCREMENT = "++"
    DECREMENT = "--"
    POINT = "."
    LESS_THAN = "<"
    GREATER_THAN = ">"
    LTOE = "<="
    GTOE = ">="
    AMPER = "&"
    BIT_OR = "|"
    BIT_XOR = "^"
    BIT_NOT = "~"
    BIT_LS = "<<"
    BIT_RS = ">>"
    HASH = "#"    
    LOG_AND = "&&"
    LOG_OR = "||"
    QUOTE = "\""

    PLUS_ASSIGN = "+="
    MINUS_ASSIGN = "-="
    MULT_ASSIGN = "*="
    DIV_ASSIGN = "/="
    MOD_ASSIGN = "%="
    AND_ASSIGN = "&="
    OR_ASSIGN = "|="
    XOR_ASSIGN = "^="

    IS_EQ = "=="
    IS_NOT_EQ = "!="
    
    // Delimiters
    COMMA = ","
    SEMICOLON = ";"

    ARROW = "->"

    LPAR = "("
    RPAR = ")"
    LBRACK = "["
    RBRACK = "]"
    LSQUIRLY = "{"
    RSQUIRLY = "}"

    // Keywords
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
    AUTO = "AUTO"
    BREAK = "BREAK"
    CASE = "CASE"
    CONST = "CONST"
    CONTINUE = "CONTINUE"
    DEFAULT = "DEFAULT"
    DO = "DO"
    EXTERN = "EXTERN"
    FOR = "FOR"
    GOTO = "GOTO"
    INLINE = "INLINE"
    REGISTER = "REGISTER"
    RESTRICT = "RESTRICT"
    SIZEOF = "SIZEOF"
    STATIC = "STATIC"
    STATIC_ASSERT = "STATIC_ASSERT"
    SWITCH = "SWITCH"
    TYPEDEF = "TYPEDEF"
    UNION = "UNION"
    UNSIGNED = "UNSIGNED"
    VOID = "VOID"
    VOLATILE = "VOLATILE"
    WHILE = "WHILE"
    ENUM_TYPE = "ENUM"
    SIGNED_TYPE = "SIGNED"
    STRUCT_TYPE = "STRUCT"
    INT_TYPE = "INT_TYPE"
    LONG_TYPE = "LONG_TYPE"
    SHORT_TYPE = "SHORT_TYPE"
    FLOAT_TYPE = "FLOAT_TYPE"
    DOUBLE_TYPE = "DOUBLE_TYPE"
    CHAR_TYPE = "CHAR_TYPE"
)

var keywords = map[string] TokenType{
    "auto": AUTO,
    "break": BREAK,
    "case": CASE,
    "char": CHAR_TYPE,
    "const": CONST,
    "continue": CONTINUE,
    "default": DEFAULT,
    "do": DO,
    "double": DOUBLE_TYPE,
    "enum": ENUM_TYPE,
    "extern": EXTERN,
    "float": FLOAT_TYPE,
    "for": FOR,
    "goto": GOTO,
    "inline": INLINE,
    "int": INT_TYPE,
    "long": LONG_TYPE,
    "register": REGISTER,
    "restrict": RESTRICT,
    "short": SHORT_TYPE,
    "signed": SIGNED_TYPE,
    "sizeof": SIZEOF,
    "static": STATIC,
    "static_assert": STATIC_ASSERT,
    "struct": STRUCT_TYPE,
    "switch": SWITCH,
    "typedef": TYPEDEF,
    "union": UNION,
    "unsigned": UNSIGNED,
    "void": VOID,
    "volatile": VOLATILE,
    "while": WHILE,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

func LookupIdentifier(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }

    return IDENTIFIER
}

func GetCategoryName(ttype TokenType, lit string) string {
    // switch
    switch ttype {
        case EQUAL, PLUS, MINUS, BANG, ASTERISK, SLASH, LINE_COMMENT, MODULO, INCREMENT, DECREMENT, POINT, LESS_THAN, 
            GREATER_THAN, LTOE, GTOE, AMPER, BIT_OR, BIT_XOR, BIT_NOT, BIT_LS, BIT_RS, HASH, LOG_AND, LOG_OR, QUOTE, 
            PLUS_ASSIGN, MINUS_ASSIGN, MULT_ASSIGN, DIV_ASSIGN, MOD_ASSIGN, AND_ASSIGN, OR_ASSIGN, XOR_ASSIGN, 
            IS_EQ, IS_NOT_EQ, ARROW:
            return "operator"
        case COMMA, SEMICOLON, LPAR, RPAR, LSQUIRLY, RSQUIRLY, LBRACK, RBRACK:
            return "delimiter"
        case ILLEGAL:
            return "error"
        case INT:
            return "int"
        case FLOATING:
            return "floating point value"
        case CHAR:
            return "char"
        case EOF:
            return "eof"
        case COMMENT:
            return "comment"
    }

    if LookupIdentifier(lit) == IDENTIFIER {
        return "identifier"
    } else {
        return "key_word"
    }
}
