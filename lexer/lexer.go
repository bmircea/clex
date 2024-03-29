package lexer

import (
    "clex/token"
)

type Lexer struct {
    input string
    position int // points to current char
    readPosition int // after current char (peek)
    line int // line no
    ch byte // current char
    stringOpen bool // is string open
}

func New(input string) *Lexer {
    l := &Lexer{input: input, line: 1, stringOpen: false}
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }

    l.position = l.readPosition
    l.readPosition += 1
}

func (l* Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    } else {
        return l.input[l.readPosition]
    }
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token
    l.skipWhitespace()

    switch l.ch {
        case '=':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.IS_EQ, Literal: literal, Line: l.line}
            } else {
                tok = newToken(token.EQUAL, l.ch, l.line)
            }
        case '!':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.IS_NOT_EQ, Literal: literal, Line: l.line}
            } else {
                tok = newToken(token.BANG, l.ch, l.line)
            }
        case '-':
            if l.peekChar() == '>' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.ARROW, Literal: literal, Line: l.line}
            } else if l.peekChar() == '='{
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.MINUS_ASSIGN, Literal: literal, Line: l.line}
            } else if l.peekChar() == '-' {
                ch:= l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.DECREMENT, Literal: literal, Line: l.line}
            } else {
                tok = newToken(token.MINUS, l.ch, l.line)
            }
        case '/':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.DIV_ASSIGN, Literal: literal, Line: l.line}
            } else if l.peekChar() == '*' {
                literal := l.readComment(true)
                
                tok = token.Token{Type: token.COMMENT, Literal: literal, Line: l.line}

            } else  if l.peekChar() == '/'{
                literal := l.readComment(false)
                tok = token.Token{Type: token.COMMENT, Literal: literal, Line: l.line}
            } else {
                tok = newToken(token.SLASH, l.ch, l.line)
            }
        case '*':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.MULT_ASSIGN, Literal: literal, Line: l.line}
            } else {
                tok = newToken(token.ASTERISK, l.ch, l.line)
            }
        case '<':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.LTOE, Literal: literal, Line: l.line}
            } else if l.peekChar() == '<' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.BIT_LS, Literal: literal, Line: l.line}
            } else {
                tok = newToken(token.LESS_THAN, l.ch, l.line)
            }
        case '>':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.GTOE, Literal: literal, Line: l.line}
            } else if l.peekChar() == '>' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.BIT_RS, Literal: literal, Line: l.line}
            } else {
                tok = newToken(token.GREATER_THAN, l.ch, l.line)
            }
        case ';':
            tok = newToken(token.SEMICOLON, l.ch, l.line)
        case '(':
            tok = newToken(token.LPAR, l.ch, l.line)
        case ')':
            tok = newToken(token.RPAR, l.ch, l.line)
        case ',':
            tok = newToken(token.COMMA, l.ch, l.line)
        case '+':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.PLUS_ASSIGN, Literal: literal, Line: l.line}
            } else if l.peekChar() == '+' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.INCREMENT, Literal: literal, Line: l.line}
            } else {
                tok = newToken(token.PLUS, l.ch, l.line)
            }
        case '{':
            tok = newToken(token.LSQUIRLY, l.ch, l.line)
        case '}':
            tok = newToken(token.RSQUIRLY, l.ch, l.line)
        case '[':
            tok = newToken(token.LBRACK, l.ch, l.line)
        case ']':
            tok = newToken(token.RBRACK, l.ch, l.line)
        case '^':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.XOR_ASSIGN, Literal: literal, Line: l.line}
            } else {
                tok = newToken(token.BIT_XOR, l.ch, l.line)
            }
        case '|':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.OR_ASSIGN, Literal: literal, Line: l.line}
            } else if l.peekChar() == '|' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.LOG_OR, Literal: literal, Line: l.line}
            } else {
                tok = newToken(token.BIT_OR, l.ch, l.line)
            }
        case '&':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.AND_ASSIGN, Literal: literal, Line: l.line}
            } else if l.peekChar() == '&' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.LOG_AND, Literal: literal, Line: l.line}
            } else {
                tok = newToken(token.AMPER, l.ch, l.line)
            }
        case '%':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)

                tok = token.Token{Type: token.MOD_ASSIGN, Literal: literal, Line: l.line}

            } else {
                tok = newToken(token.MODULO, l.ch, l.line)
            }
        case '"':
            l.stringOpen = !l.stringOpen
            tok = newToken(token.QUOTE, l.ch, l.line)
        case '#':
            tok = newToken(token.HASH, l.ch, l.line)
        case '~':
            tok = newToken(token.BIT_NOT, l.ch, l.line)
        case '.':
            tok = newToken(token.POINT, l.ch, l.line)
        case 0:
            tok.Type = token.EOF
            tok.Literal = ""
            tok.Line = l.line
        default:
            if isLetter(l.ch) {
                if l.stringOpen {
                    tok = newToken(token.CHAR, l.ch, l.line)
                } else {
                    //fmt.Println(fmt.Sprintf("Looking up for %s \n", string(l.ch)))
                    tok.Literal = l.readIdentifier()
                    //fmt.Println(fmt.Sprintf("Identifier is %s \n", tok.Literal))
                    tok.Type = token.LookupIdentifier(tok.Literal)
                    tok.Line = l.line
                    return tok
                }
            } else if isDigit(l.ch){
                tok.Type = token.INT
                tok.Type, tok.Literal = l.readNumber()
                tok.Line = l.line
                return tok
            } else {
                tok = newToken(token.ILLEGAL, l.ch, l.line)
            }
    }

    l.readChar()
    return tok
}


func newToken(tokenType token.TokenType, ch byte, l int) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch), Line: l}
}


func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
}

func (l *Lexer) readComment(multiline bool) string {
    position := l.position // /
    
    if multiline {
        for !(l.ch == '*' && l.peekChar() == '/') {
            l.readChar()
        }

        l.readChar()
    } else {
        for l.ch != '\n' && l.ch != '\r' {
            l.readChar()
        }
    }

    return l.input[position:l.position]

}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}


func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        if l.ch == '\n' {
            l.line++
        } else if l.ch == '\r' && l.peekChar() == '\n'{
            l.readChar()
            l.line++
        } 

        l.readChar()
    }
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}


func (l *Lexer) readNumber() (token.TokenType, string) {
    position := l.position
    toktype := token.INT
    countDecimalPoints := 0
    for isDigit(l.ch) || (l.ch == '.' && countDecimalPoints == 0) {
        if l.ch == '.' {
            toktype = token.FLOATING
            countDecimalPoints = 1
        }
        l.readChar()
    }

    return token.TokenType(toktype), l.input[position:l.position]
        
}
