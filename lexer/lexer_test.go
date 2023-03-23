// lexer/lexer_test.go

package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ILLEGAL, "ILLEGAL"},
		{token.EOF, "EOF"},

		// Identifiers + literals
		{token.IDENT, "IDENT"},
		{token.INT, "INT"},

		// Operators
		{token.ASSIGN, "="},
		{token.PLUS, "+"},

		// Delimiters
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},

		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},

		// Keywords
		{token.FUNCTION, "FUNCTION"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedToken, tok.Type)
		}

		if tok.Type != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
