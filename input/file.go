package input

import (
	"capital-gains/dtos"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFromFile(filePath string) [][]dtos.Transaction {
	file, _ := os.Open(filePath)
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	rawArrays := strings.Split(string(content), "]")
	var allTransactions [][]dtos.Transaction

	for _, raw := range rawArrays {
		raw = strings.TrimSpace(raw)
		if raw == "" {
			continue
		}

		raw = raw + "]"
		var transactions []dtos.Transaction
		if err := json.Unmarshal([]byte(raw), &transactions); err != nil {
			fmt.Println("Error parsing JSON array:", err)
			panic(err)
		}
		allTransactions = append(allTransactions, transactions)
	}

	return allTransactions
}

func ReadOutput(filepath string) []dtos.OutputData {
	file, _ := os.Open(filepath)
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	var output []dtos.OutputData
	if err := json.Unmarshal(content, &output); err != nil {
		fmt.Println("Error parsing JSON array:", err)
		panic(err)
	}

	return output
}
