package calculator

import "capital-gains/dtos"

func CalculateTaxValue(value float64) float64 {
	if value <= 0 {
		return 0
	}

	return value * 0.20
}

func ShouldPayTaxes(op *dtos.Transaction) bool {
	return op.Quantity*op.UnitCost > 20000
}
