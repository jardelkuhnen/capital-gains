package calculator

import (
	"capital-gains/constants"
	"capital-gains/dtos"
	"capital-gains/weightedaverage"
)

func Calculate(inputData []dtos.Transaction) []dtos.OutputData {
	outputData := make([]dtos.OutputData, 0)

	storedData := dtos.StoreData{}

	for _, op := range inputData {

		if op.Operation == constants.BUY {
			average := weightedaverage.CalculateWeightedAverage(storedData.ActualQuantity, op.Quantity, storedData.ActualAverage, op.UnitCost)
			storedData.ActualQuantity += op.Quantity
			storedData.ActualAverage = average

			outputData = append(outputData, dtos.OutputData{Tax: 0})
			continue
		}

		if op.Operation == constants.SELL && !ShouldPayTaxes(&op) {
			if profit := weightedaverage.CalculateProfit(storedData.ActualAverage, op.UnitCost, op.Quantity); profit < 0 {
				storedData.AccumulatedLoss += profit
			}

			storedData.ActualQuantity -= op.Quantity
			outputData = append(outputData, dtos.OutputData{Tax: 0})
			continue
		}

		profit := weightedaverage.CalculateProfit(storedData.ActualAverage, op.UnitCost, op.Quantity)
		if profit < 0 {
			storedData.AccumulatedLoss += profit
		} else {
			actualLoss := storedData.AccumulatedLoss

			if storedData.AccumulatedLoss < 0 {
				storedData.AccumulatedLoss += profit
			}
			profit += actualLoss
		}

		storedData.ActualQuantity -= op.Quantity
		taxValue := CalculateTaxValue(profit)
		outputData = append(outputData, dtos.OutputData{Tax: taxValue})

	}

	return outputData
}
