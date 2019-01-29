package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	ASTERISK = "*"
	SLASH = "/"

	BANG = "!"
	NOT_EQ = "!="
	EQ = "=="
	LT = "<"
	GT = ">"
	
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
	TRUE = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"lt": LET,
	"if": IF,
	"el": ELSE,
	"rt": RETURN,
	"tr": TRUE,
	"fl": FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
