package parser

import (
	"fmt"

	"github.com/AvicennaJr/Nuru/ast"
)

func (p *Parser) parseAssignmentExpression(exp ast.Expression) ast.Expression {
	switch node := exp.(type) {
	// temporarily making let keyword optional
	case *ast.Identifier:
		e := &ast.Assign{
			Token: p.curToken,
			Name:  exp.(*ast.Identifier),
		}
		precedence := p.curPrecedence()
		p.nextToken()
		e.Value = p.parseExpression(precedence)
		return e

	case *ast.IndexExpression:
	default:
		if node != nil {
			msg := fmt.Sprintf("Mstari %d:Tulitegemea kupata kitambulishi au array, badala yake tumepata: %s", p.curToken.Line, node.TokenLiteral())
			p.errors = append(p.errors, msg)
		} else {
			msg := fmt.Sprintf("Mstari %d: Umekosea mkuu", p.curToken.Line)
			p.errors = append(p.errors, msg)
		}
		return nil
	}

	ae := &ast.AssignmentExpression{Token: p.curToken, Left: exp}

	p.nextToken()

	ae.Value = p.parseExpression(LOWEST)

	return ae
}
