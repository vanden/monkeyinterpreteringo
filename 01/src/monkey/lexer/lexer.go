package lexer

type Lexer struct {
	input        string
	position     int  // points to current char of input
	readPosition int  // position after current char
	ch           byte // char under inspection
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
