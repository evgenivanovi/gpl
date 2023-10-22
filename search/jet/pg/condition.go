package pg

import (
	"reflect"

	"github.com/evgenivanovi/gpl/search"
	"github.com/go-jet/jet/v2/postgres"
)

/* __________________________________________________ */

func VisitCondition(
	column postgres.Column,
	condition search.SearchCondition,
) postgres.BoolExpression {

	if search.IsEquals(condition) {
		return visitEqualsCondition(column, condition)
	}

	if search.IsNotEquals(condition) {
		return visitIsNotEqualsCondition(column, condition)
	}

	if search.IsContainsAny(condition) {
		return visitContainsAnyCondition(column, condition)
	}

	if search.IsNotContainsAll(condition) {
		return visitNotContainsAllCondition(column, condition)
	}

	return nil

}

/* __________________________________________________ */

func visitEqualsCondition(
	column postgres.Column,
	condition search.SearchCondition,
) postgres.BoolExpression {

	/* __________________________________________________ */

	isNilValue := condition.Value().IsNil()

	if isNilValue {
		col := column.(postgres.ColumnBool)
		return col.IS_NULL()
	}

	/* __________________________________________________ */

	isBooleanColumn := reflect.TypeOf(column).
		Implements(reflect.TypeOf((*postgres.ColumnBool)(nil)).Elem())

	isBooleanValue := condition.Value().IsBool()

	if isBooleanColumn && isBooleanValue {
		col := column.(postgres.ColumnBool)
		value := condition.Value().GetBool()
		return col.EQ(postgres.Bool(value))
	}

	/* __________________________________________________ */

	isStringColumn := reflect.TypeOf(column).
		Implements(reflect.TypeOf((*postgres.ColumnString)(nil)).Elem())

	isStringValue := condition.Value().IsString()

	if isStringColumn && isStringValue {
		col := column.(postgres.ColumnString)
		value := condition.Value().GetString()
		return col.EQ(postgres.String(value))
	}

	/* __________________________________________________ */

	isIntegerColumn := reflect.TypeOf(column).
		Implements(reflect.TypeOf((*postgres.ColumnInteger)(nil)).Elem())

	if isIntegerColumn {

		col := column.(postgres.ColumnInteger)

		if condition.Value().IsInt64() {
			value := condition.Value().GetInt64()
			return col.EQ(postgres.Int64(value))
		}

		if condition.Value().IsUint64() {
			value := condition.Value().GetUint64()
			return col.EQ(postgres.Uint64(value))
		}

	}

	/* __________________________________________________ */

	panic("Invalid invocation of method or invalid parameters.")

}

func visitIsNotEqualsCondition(
	column postgres.Column,
	condition search.SearchCondition,
) postgres.BoolExpression {

	/* __________________________________________________ */

	isNilValue := condition.Value().IsNil()

	if isNilValue {
		col := column.(postgres.ColumnBool)
		return col.IS_NOT_NULL()
	}

	/* __________________________________________________ */

	isBooleanColumn := reflect.TypeOf(column).
		Implements(reflect.TypeOf((*postgres.ColumnBool)(nil)).Elem())

	isBooleanValue := condition.Value().IsBool()

	if isBooleanColumn && isBooleanValue {
		col := column.(postgres.ColumnBool)
		value := condition.Value().GetBool()
		return col.NOT_EQ(postgres.Bool(value))
	}

	/* __________________________________________________ */

	isStringColumn := reflect.TypeOf(column).
		Implements(reflect.TypeOf((*postgres.ColumnString)(nil)).Elem())

	isStringValue := condition.Value().IsString()

	if isStringColumn && isStringValue {
		col := column.(postgres.ColumnString)
		value := condition.Value().GetString()
		return col.NOT_EQ(postgres.String(value))
	}

	/* __________________________________________________ */

	isIntegerColumn := reflect.TypeOf(column).
		Implements(reflect.TypeOf((*postgres.ColumnInteger)(nil)).Elem())

	if isIntegerColumn {

		col := column.(postgres.ColumnInteger)

		if condition.Value().IsInt64() {
			value := condition.Value().GetInt64()
			return col.NOT_EQ(postgres.Int64(value))
		}

		if condition.Value().IsUint64() {
			value := condition.Value().GetUint64()
			return col.NOT_EQ(postgres.Uint64(value))
		}

	}

	/* __________________________________________________ */

	panic("Invalid invocation of method or invalid parameters.")

}

/* __________________________________________________ */

func visitContainsAnyCondition(
	column postgres.Column,
	condition search.SearchCondition,
) postgres.BoolExpression {

	/* __________________________________________________ */

	isStringColumn := reflect.TypeOf(column).
		Implements(reflect.TypeOf((*postgres.ColumnString)(nil)).Elem())

	isStringsValue := condition.Value().IsStrings()

	if isStringColumn && isStringsValue {
		col := column.(postgres.ColumnString)
		values := condition.Value().GetStrings()

		exp := make([]postgres.Expression, 0)
		for _, value := range values {
			exp = append(exp, postgres.String(value))
		}
		return col.IN(exp...)
	}

	/* __________________________________________________ */

	isIntegerColumn := reflect.TypeOf(column).
		Implements(reflect.TypeOf((*postgres.ColumnInteger)(nil)).Elem())

	if isIntegerColumn {

		col := column.(postgres.ColumnInteger)

		if condition.Value().IsInts64() {
			values := condition.Value().GetInts64()

			exp := make([]postgres.Expression, 0)
			for _, value := range values {
				exp = append(exp, postgres.Int64(value))
			}
			return col.IN(exp...)
		}

		if condition.Value().IsUints64() {
			values := condition.Value().GetUints64()

			exp := make([]postgres.Expression, 0)
			for _, value := range values {
				exp = append(exp, postgres.Uint64(value))
			}
			return col.IN(exp...)
		}

	}

	/* __________________________________________________ */

	panic("Invalid invocation of method or invalid parameters.")

}

func visitNotContainsAllCondition(
	column postgres.Column,
	condition search.SearchCondition,
) postgres.BoolExpression {

	/* __________________________________________________ */

	isStringColumn := reflect.TypeOf(column).
		Implements(reflect.TypeOf((*postgres.ColumnString)(nil)).Elem())

	isStringsValue := condition.Value().IsStrings()

	if isStringColumn && isStringsValue {
		col := column.(postgres.ColumnString)
		values := condition.Value().GetStrings()

		exp := make([]postgres.Expression, 0)
		for _, value := range values {
			exp = append(exp, postgres.String(value))
		}
		return col.NOT_IN(exp...)
	}

	/* __________________________________________________ */

	isIntegerColumn := reflect.TypeOf(column).
		Implements(reflect.TypeOf((*postgres.ColumnInteger)(nil)).Elem())

	if isIntegerColumn {

		col := column.(postgres.ColumnInteger)

		if condition.Value().IsInts64() {
			values := condition.Value().GetInts64()

			exp := make([]postgres.Expression, 0)
			for _, value := range values {
				exp = append(exp, postgres.Int64(value))
			}
			return col.NOT_IN(exp...)
		}

		if condition.Value().IsUints64() {
			values := condition.Value().GetUints64()

			exp := make([]postgres.Expression, 0)
			for _, value := range values {
				exp = append(exp, postgres.Uint64(value))
			}
			return col.NOT_IN(exp...)
		}

	}

	/* __________________________________________________ */

	panic("Invalid invocation of method or invalid parameters.")

}

/* __________________________________________________ */
