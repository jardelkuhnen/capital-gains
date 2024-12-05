package main

import (
	"capital-gains/calculator"
	"capital-gains/input"
	"capital-gains/output"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		inputData := input.ReadFromFile(os.Args[1])

		for _, d := range inputData {
			output.PrintOutput(calculator.Calculate(d))
		}

		return
	}

	output.PrintOutput(calculator.Calculate(input.ReadFromStdin()))
}
