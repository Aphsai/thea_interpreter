package parser

import (
	"testing"
	"thea_interpreter/ast"
	"thea_interpeter/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `lt x = 5; lt y = 10; lt foobar = 838383;`

	l := lexer.New(input)
	p := New(l)

	program := p.parseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d"
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i];
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "lt" {
		t.Errorf("s.TokenLiteral not 'let.' got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s.' got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s.' got=%s", name, letStmt.Name)
		return false
	}
	return true
		
}
