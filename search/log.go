package search

import (
	"context"
	"log/slog"

	"github.com/evgenivanovi/gpl/stdx"
	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
)

type SpecificationModel struct {
	SearchConditions *SearchConditionsModel `json:"search,omitempty"`
	SliceConditions  *SliceConditionsModel  `json:"slice,omitempty"`
	OrderConditions  *OrderConditionsModel  `json:"order,omitempty"`
}

type SearchConditionModel struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Kind  string `json:"type,omitempty"`
	Op    string `json:"op,omitempty"`
}

type SearchConditionsModel struct {
	Conditions []SearchConditionModel `json:"conditions,omitempty"`
}

type SliceConditionsModel struct {
	Limit  int64 `json:"limit,omitempty"`
	Offset int64 `json:"offset,omitempty"`
}

type OrderConditionsModel struct {
	Conditions []OrderConditionModel `json:"conditions,omitempty"`
}

type OrderConditionModel struct {
	Key string `json:"key,omitempty"`
	Op  string `json:"op,omitempty"`
}

func Log(spec Specification) {
	slogx.Log().Debug(
		"Specification query", LogAttr(spec),
	)
}

func LogCtx(ctx context.Context, spec Specification) {
	slogx.FromCtx(ctx).Debug(
		"Specification query", LogAttr(spec),
	)
}

func LogAttr(spec Specification) slog.Attr {
	return slog.Any("specification", specificationModel(spec))
}

func specificationModel(spec Specification) *SpecificationModel {
	if spec == nil {
		return nil
	}

	return &SpecificationModel{
		SearchConditions: searchConditionsModel(spec.SearchConditions()),
		SliceConditions:  sliceConditionsModel(spec.SliceConditions()),
		OrderConditions:  orderConditionsModel(spec.OrderConditions()),
	}
}

func searchConditionsModel(conditions *SearchConditions) *SearchConditionsModel {
	if conditions == nil {
		return nil
	}
	return &SearchConditionsModel{
		Conditions: searchConditionModels(conditions.Conditions()),
	}
}

func searchConditionModels(conditions []SearchCondition) []SearchConditionModel {
	res := make([]SearchConditionModel, 0)
	for _, condition := range conditions {
		res = append(res, searchConditionModel(condition))
	}
	return res
}

func searchConditionModel(condition SearchCondition) SearchConditionModel {
	return SearchConditionModel{
		Key:   condition.Key().String(),
		Value: condition.Value().String(),
		Kind:  stdx.ValueTypeName(condition.Value()),
		Op:    condition.Op().String(),
	}
}

func sliceConditionsModel(conditions *SliceCondition) *SliceConditionsModel {
	if conditions == nil || !conditions.Chunked() {
		return nil
	}
	return &SliceConditionsModel{
		Limit:  conditions.Chunk().limit,
		Offset: conditions.Chunk().offset,
	}
}

func orderConditionsModel(conditions *OrderConditions) *OrderConditionsModel {
	if conditions == nil {
		return nil
	}
	return &OrderConditionsModel{
		Conditions: orderConditionModels(conditions.Conditions()),
	}
}

func orderConditionModels(conditions []OrderCondition) []OrderConditionModel {
	res := make([]OrderConditionModel, 0)
	for _, condition := range conditions {
		res = append(res, orderConditionModel(condition))
	}
	return res
}

func orderConditionModel(condition OrderCondition) OrderConditionModel {
	return OrderConditionModel{
		Key: condition.Key().String(),
		Op:  condition.Op().String(),
	}
}
