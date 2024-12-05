package output

import (
	"capital-gains/dtos"
	"encoding/json"
	"fmt"
)

func PrintOutput(output []dtos.OutputData) {
	bytes, _ := json.Marshal(output)
	fmt.Println(string(bytes))
}
