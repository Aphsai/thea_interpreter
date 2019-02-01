package parser

import (
	"thea_interpreter/ast"
	"thea_interpreter/lexer"
	"thea_interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil

}

function parseProgram() {
	program = newProgramASTNode()
	advanceTokens()

	for (currentToken() != EOF_TOKEN) {
		statement = null
		if (currentToken() == LET_TOKEN) {
			statement = parseLetStatement()
		} else if (currentToken() == RETURN_TOKEN) {
			statement = parseReturnStatement()
		} else if (currentToken() == IF_TOKEN) {
			statement = parseIfStatement()
		}
		if (statement != null) {
			program.Statements.push(statement)
		}
		advanceTokens()
	}

	return program;
}

function parseLetStatement() {
	advanceTokens()
	identifier = parseIdentifier()
	advanceTokens()
	if currentToken() != EQUAL_TOKEN {
		parseError("No equal sign.")
		return null
	}
	advanceTokens()

	value = parseExpressions()
	variableStatement = newVariableStatementASTNode()
	variableStatement.identifier = identifier
	variableStatement.value = value
	return variableStatement
}

function parseIdentifier() {
	identifier = newIdentifierASTNode()
	identifier.token = currentToken()
	return identifier
}

function parseExpressions() {
	if (currentToken() == INTEGER_TOKEN) {
		if (nextToken() == PLUS_TOKEN) {
			return parseOperatorExpression()
		} else if (nextToken() == SEMICOLON_TOKEN) {
			return parseIntegerLiteral()
		}
	} else if (currentToken() == LEFT_PAREN) {
		return parseGroupedExpression()
	}
}

function parseOperatorExpression() {
	operatorExpression = newOperatorExpression()
	operatorExpression.left = parseIntegerLiteral()
	operatorExpression.operator = currentToken()
	operatorExpression.right = parseExpression()

	return operatorExpression()
}
