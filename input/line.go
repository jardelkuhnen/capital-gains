package input

import (
	"bufio"
	"capital-gains/dtos"
	"encoding/json"
	"os"
)

func ReadFromStdin() []dtos.Transaction {
	var result []dtos.Transaction
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		var data dtos.Transaction
		_ = json.Unmarshal(line, &data)
		result = append(result, data)

		if string(line) == "" {
			break
		}
	}

	return result
}
