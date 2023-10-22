package search

type Key string

func (k Key) String() string {
	return string(k)
}

type Specification interface {
	SearchConditions() *SearchConditions
	SliceConditions() *SliceCondition
	OrderConditions() *OrderConditions
}

type SpecificationTemplate struct {
	slices   *SliceCondition
	orders   *OrderConditions
	searches *SearchConditions
}

func NewSpecificationTemplate() *SpecificationTemplate {
	return &SpecificationTemplate{
		slices:   NewSlice(),
		orders:   NewOrders(),
		searches: NewSearches(),
	}
}

func (s *SpecificationTemplate) SearchConditions() *SearchConditions {
	return s.searches
}

func (s *SpecificationTemplate) WithSearch(ops ...SearchCondition) *SpecificationTemplate {
	WithSearchConditions(ops...)(s.searches)
	return s
}

func (s *SpecificationTemplate) SliceConditions() *SliceCondition {
	return s.slices
}

func (s *SpecificationTemplate) WithSlice(ops ...SliceConditionOp) *SpecificationTemplate {
	for _, op := range ops {
		op(s.slices)
	}
	return s
}

func (s *SpecificationTemplate) OrderConditions() *OrderConditions {
	return s.orders
}

func (s *SpecificationTemplate) WithOrder(ops ...OrderCondition) *SpecificationTemplate {
	WithOrderConditions(ops...)(s.orders)
	return s
}
