package dtos

import "capital-gains/constants"

type Transaction struct {
	Operation constants.Operation `json:"operation"`
	UnitCost  float64             `json:"unit-cost"`
	Quantity  float64             `json:"quantity"`
}
