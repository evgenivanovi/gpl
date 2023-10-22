package search

import slices "github.com/evgenivanovi/gpl/std/slice"

/* __________________________________________________ */

type OrderOperator string

func (o OrderOperator) String() string {
	return string(o)
}

const ASC OrderOperator = "ASC"
const DESC OrderOperator = "DESC"

/* __________________________________________________ */

type OrderConditionsOp func(*OrderConditions)

func WithOrderCondition(condition OrderCondition) OrderConditionsOp {
	return func(cond *OrderConditions) {
		cond.AddCondition(condition)
	}
}

func WithOrderConditions(conditions ...OrderCondition) OrderConditionsOp {
	return func(cond *OrderConditions) {
		cond.AddConditions(conditions...)
	}
}

/* __________________________________________________ */

type OrderConditions struct {
	conditions []OrderCondition
}

func NewOrders() *OrderConditions {
	return &OrderConditions{
		conditions: make([]OrderCondition, 0),
	}
}

func (c *OrderConditions) IsEmpty() bool {
	return slices.IsEmpty(c.conditions)
}

func (c *OrderConditions) Conditions() []OrderCondition {
	return c.conditions
}

func (c *OrderConditions) AddCondition(condition OrderCondition) {
	c.conditions = append(c.conditions, condition)
}

func (c *OrderConditions) AddConditions(conditions ...OrderCondition) {
	c.conditions = append(c.conditions, conditions...)
}

/* __________________________________________________ */

type OrderConditionOp func(options *OrderCondition)

func WithASC(key Key) OrderConditionOp {
	return func(cond *OrderCondition) {
		cond.key = key
		cond.op = ASC
	}
}

func WithDESC(key Key) OrderConditionOp {
	return func(cond *OrderCondition) {
		cond.key = key
		cond.op = DESC
	}
}

/* __________________________________________________ */

type OrderCondition struct {
	key Key
	op  OrderOperator
}

func AscOrder(key Key) OrderCondition {
	condition := &OrderCondition{}
	WithASC(key)(condition)
	return *condition
}

func DescOrder(key Key) OrderCondition {
	condition := &OrderCondition{}
	WithDESC(key)(condition)
	return *condition
}

func (c *OrderCondition) Key() Key {
	return c.key
}

func (c *OrderCondition) Op() OrderOperator {
	return c.op
}

func (c *OrderCondition) ASC() bool {
	return c.op == ASC
}

func (c *OrderCondition) DESC() bool {
	return c.op == DESC
}

/* __________________________________________________ */
