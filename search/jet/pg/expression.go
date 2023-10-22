package pg

import (
	"github.com/evgenivanovi/gpl/search"
	"github.com/go-jet/jet/v2/postgres"
)

/* __________________________________________________ */

func TrueExpression() postgres.BoolExpression {
	return postgres.Bool(true).IS_TRUE()
}

func FalseExpression() postgres.BoolExpression {
	return postgres.Bool(true).IS_FALSE()
}

/* __________________________________________________ */

func SearchExpression(
	spec search.Specification,
	mapping map[search.Key]postgres.Column,
) postgres.BoolExpression {

	expression := TrueExpression()

	for _, condition := range spec.SearchConditions().Conditions() {

		column := mapping[condition.Key()]
		if column == nil {
			continue
		}

		exp := VisitCondition(column, condition)
		if exp == nil {
			continue
		}

		expression = expression.AND(exp)

	}

	return expression

}

func OrderExpression(
	spec search.Specification,
	mapping map[search.Key]postgres.Column,
) []postgres.OrderByClause {

	expressions := make([]postgres.OrderByClause, 0)

	for _, condition := range spec.OrderConditions().Conditions() {

		column := mapping[condition.Key()]
		if column == nil {
			continue
		}

		if condition.ASC() {
			expressions = append(expressions, column.ASC())
		}

		if condition.DESC() {
			expressions = append(expressions, column.DESC())
		}

	}

	return expressions

}

/* __________________________________________________ */
