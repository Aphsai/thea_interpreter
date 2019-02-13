package ast 

import (
	"thea_interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statements{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "lt"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "variable"},
					Value: "variable",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "another_variable"},
					Value: "another_variable",
				},
			},
		},
	}

	if program.String() !+ "lt variable = another_variable;" {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}
