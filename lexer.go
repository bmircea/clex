package lexer

import (
	"fmt"
	"lexer/token"
)

type Lexer struct {
    input string
    position int // points to current char
    readPosition int // after current char (peek)
    line int // line no
    ch byte // current char
}

func New(input string) *Lexer {
    l := &Lexer{input: input, line: 1}
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
            tok = newToken(token.MINUS, l.ch, l.line)
        case '/':
            tok = newToken(token.SLASH, l.ch, l.line)
        case '*':
            tok = newToken(token.ASTERISK, l.ch, l.line)
        case '<':
            tok = newToken(token.LESS_THAN, l.ch, l.line)
        case '>':
            tok = newToken(token.GREATER_THAN, l.ch, l.line)
        case ';':
            tok = newToken(token.SEMICOLON, l.ch, l.line)
        case '(':
            tok = newToken(token.LPAR, l.ch, l.line)
        case ')':
            tok = newToken(token.RPAR, l.ch, l.line)
        case ',':
            tok = newToken(token.COMMA, l.ch, l.line)
        case '+':
            tok = newToken(token.PLUS, l.ch, l.line)
        case '{':
            tok = newToken(token.LSQUIRLY, l.ch, l.line)
        case '}':
            tok = newToken(token.RSQUIRLY, l.ch, l.line)
        case 0:
            tok.Type = token.EOF
            tok.Literal = ""
            tok.Line = l.line
        default:
            if isLetter(l.ch) {
                fmt.Println(fmt.Sprintf("Looking up for %s \n", string(l.ch)))
                tok.Literal = l.readIdentifier()
                fmt.Println(fmt.Sprintf("Identifier is %s \n", tok.Literal))
                tok.Type = token.LookupIdentifier(tok.Literal)
                tok.Line = l.line
                return tok
            } else if isDigit(l.ch){
                tok.Type = token.INT
                tok.Literal = l.readNumber()
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


func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
        
}
