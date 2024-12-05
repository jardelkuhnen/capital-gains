package calculator

import (
	"capital-gains/input"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type CalculatorSuite struct {
	suite.Suite
}

func Test_Calculator(t *testing.T) {
	suite.Run(t, new(CalculatorSuite))
}

func (s *CalculatorSuite) Test_Calculate() {

	scenarios := []struct {
		name           string
		inputFile      string
		expectedOutput string
	}{
		{
			name:           "test_case_1",
			inputFile:      "_testdata/test_case_1_input.txt",
			expectedOutput: "_testdata/test_case_1_output.txt",
		},
		{
			name:           "test_case_2",
			inputFile:      "_testdata/test_case_2_input.txt",
			expectedOutput: "_testdata/test_case_2_output.txt",
		},
		{
			name:           "test_case_3",
			inputFile:      "_testdata/test_case_3_input.txt",
			expectedOutput: "_testdata/test_case_3_output.txt",
		},
		{
			name:           "test_case_4",
			inputFile:      "_testdata/test_case_4_input.txt",
			expectedOutput: "_testdata/test_case_4_output.txt",
		},
		{
			name:           "test_case_5",
			inputFile:      "_testdata/test_case_5_input.txt",
			expectedOutput: "_testdata/test_case_5_output.txt",
		},
		{
			name:           "test_case_6",
			inputFile:      "_testdata/test_case_6_input.txt",
			expectedOutput: "_testdata/test_case_6_output.txt",
		},
		{
			name:           "test_case_7",
			inputFile:      "_testdata/test_case_7_input.txt",
			expectedOutput: "_testdata/test_case_7_output.txt",
		},
		{
			name:           "test_case_8",
			inputFile:      "_testdata/test_case_8_input.txt",
			expectedOutput: "_testdata/test_case_8_output.txt",
		},
	}

	for _, scenario := range scenarios {
		s.Run(scenario.name, func() {
			directory, _ := os.Getwd()
			inputData := input.ReadFromFile(fmt.Sprintf("%s/%s", directory, scenario.inputFile))
			expectedOutputData := input.ReadOutput(scenario.expectedOutput)

			for _, d := range inputData {
				result := Calculate(d)

				outputJson, _ := json.Marshal(result)
				expectedJson, _ := json.Marshal(expectedOutputData)

				if string(expectedJson) != string(outputJson) {
					fmt.Println("Scenario: " + scenario.name)
					fmt.Println("----------------------------")
				}
				assert.Equal(s.T(), string(expectedJson), string(outputJson), fmt.Sprintf("Scenario {%s}", scenario.name))
			}

		})
	}

}
