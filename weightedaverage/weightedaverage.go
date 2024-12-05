package weightedaverage

func CalculateWeightedAverage(currentQuantity, purchaseQuantity, currentAveragePrice, purchasePrice float64) float64 {
	currentTotalCost := currentQuantity * currentAveragePrice
	purchaseTotalCost := purchaseQuantity * purchasePrice
	totalQuantity := currentQuantity + purchaseQuantity

	if totalQuantity == 0 {
		return 0
	}

	return (currentTotalCost + purchaseTotalCost) / totalQuantity
}

func CalculateProfit(averagePrice, sellingPrice, quantitySold float64) float64 {
	return (sellingPrice - averagePrice) * quantitySold
}
