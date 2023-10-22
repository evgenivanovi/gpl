package search

import (
	slices "github.com/evgenivanovi/gpl/std/slice"
	"github.com/evgenivanovi/gpl/stdx"
)

type SearchOperator string

func (o SearchOperator) String() string {
	return string(o)
}

const Equals SearchOperator = "=="
const NotEquals SearchOperator = "!="
const GreaterThan SearchOperator = ">"
const LessThan SearchOperator = "<"
const GreaterThanEqual SearchOperator = ">="
const LessThanEqual SearchOperator = "<="
const StartsWith SearchOperator = "^%"
const NotStartsWith SearchOperator = "!^%"
const EndsWithCharacter SearchOperator = "$%"
const NotEndsWith SearchOperator = "!$%"
const ContainsAll SearchOperator = "(*)"
const NotContainsAll SearchOperator = "!(*)"
const ContainsAny SearchOperator = "()"

type SearchConditionsOp func(*SearchConditions)

func WithSearchCondition(condition SearchCondition) SearchConditionsOp {
	return func(cond *SearchConditions) {
		cond.AddCondition(condition)
	}
}

func WithSearchConditions(conditions ...SearchCondition) SearchConditionsOp {
	return func(cond *SearchConditions) {
		cond.AddConditions(conditions...)
	}
}

type SearchConditions struct {
	conditions []SearchCondition
}

func NewSearches() *SearchConditions {
	return &SearchConditions{
		conditions: make([]SearchCondition, 0),
	}
}

func (c *SearchConditions) IsEmpty() bool {
	return slices.IsEmpty(c.conditions)
}

func (c *SearchConditions) Conditions() []SearchCondition {
	return c.conditions
}

func (c *SearchConditions) AddCondition(condition SearchCondition) {
	c.conditions = append(c.conditions, condition)
}

func (c *SearchConditions) AddConditions(conditions ...SearchCondition) {
	c.conditions = append(c.conditions, conditions...)
}

type SearchCondition struct {
	key   Key
	value stdx.Value
	op    SearchOperator
}

func NewSearch(
	key Key,
	value stdx.Value,
	op SearchOperator,
) *SearchCondition {
	return &SearchCondition{
		key:   key,
		value: value,
		op:    op,
	}
}

func NewEquality(
	key Key,
	value stdx.Value,
) *SearchCondition {
	return &SearchCondition{
		key:   key,
		value: value,
		op:    Equals,
	}
}

func NewInequality(
	key Key,
	value stdx.Value,
) *SearchCondition {
	return &SearchCondition{
		key:   key,
		value: value,
		op:    NotEquals,
	}
}

func NewNotContainsAll(
	key Key,
	value stdx.Value,
) *SearchCondition {
	return &SearchCondition{
		key:   key,
		value: value,
		op:    NotContainsAll,
	}
}

func NewContainsAny(
	key Key,
	value stdx.Value,
) *SearchCondition {
	return &SearchCondition{
		key:   key,
		value: value,
		op:    ContainsAny,
	}
}

func (c *SearchCondition) Key() Key {
	return c.key
}

func (c *SearchCondition) Value() stdx.Value {
	return c.value
}

func (c *SearchCondition) Op() SearchOperator {
	return c.op
}

func IsEquals(condition SearchCondition) bool {
	return condition.Op() == Equals
}

func IsEqualsOp(op SearchOperator) bool {
	return op == Equals
}

func IsNotEquals(condition SearchCondition) bool {
	return condition.Op() == NotEquals
}

func IsNotEqualsOp(op SearchOperator) bool {
	return op == NotEquals
}

func IsGreaterThan(condition SearchCondition) bool {
	return condition.Op() == GreaterThan
}

func IsGreaterThanOp(op SearchOperator) bool {
	return op == GreaterThan
}

func IsLessThan(condition SearchCondition) bool {
	return condition.Op() == LessThan
}

func IsLessThanOp(op SearchOperator) bool {
	return op == LessThan
}

func IsGreaterThanEqual(condition SearchCondition) bool {
	return condition.Op() == GreaterThanEqual
}

func IsGreaterThanEqualOp(op SearchOperator) bool {
	return op == GreaterThanEqual
}

func IsLessThanEqual(condition SearchCondition) bool {
	return condition.Op() == LessThanEqual
}

func IsLessThanEqualOp(op SearchOperator) bool {
	return op == LessThanEqual
}

func IsStartsWith(condition SearchCondition) bool {
	return condition.Op() == StartsWith
}

func IsStartsWithOp(op SearchOperator) bool {
	return op == StartsWith
}

func IsNotStartsWith(condition SearchCondition) bool {
	return condition.Op() == NotStartsWith
}

func IsNotStartsWithOp(op SearchOperator) bool {
	return op == NotStartsWith
}

func IsEndsWith(condition SearchCondition) bool {
	return condition.Op() == EndsWithCharacter
}

func IsEndsWithOp(op SearchOperator) bool {
	return op == EndsWithCharacter
}

func IsNotEndsWith(condition SearchCondition) bool {
	return condition.Op() == NotEndsWith
}

func IsNotEndsWithOp(op SearchOperator) bool {
	return op == NotEndsWith
}

func IsContainsAll(condition SearchCondition) bool {
	return condition.Op() == ContainsAll
}

func IsContainsAllOp(op SearchOperator) bool {
	return op == ContainsAll
}

func IsNotContainsAll(condition SearchCondition) bool {
	return condition.Op() == NotContainsAll
}

func IsNotContainsAllOp(op SearchOperator) bool {
	return op == NotContainsAll
}

func IsContainsAny(condition SearchCondition) bool {
	return condition.Op() == ContainsAny
}

func IsContainsAnyOp(op SearchOperator) bool {
	return op == ContainsAny
}
