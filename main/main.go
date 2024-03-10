package main 


import (
    "os"
    "fmt"
    "clex/token"
    "clex/lexer"
)



func main() {
    filename := os.Args[1]

    file, err := os.ReadFile(filename)
    if err != nil {
        panic("Failed to open file")
    }


    fmt.Print(string(file))
    
    l := lexer.New(string(file))
    tok := l.NextToken()

    for tok.Type != token.EOF {
        fmt.Println(fmt.Sprintf("'%s', %s; %d, linia %d", tok.Literal, tok.Type, len(tok.Literal), tok.Line))
        tok = l.NextToken()
    } 

}
