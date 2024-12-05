
# CapitalGains
Application used to calculate income tax on capital gains.
Each buy or sell operation is provided on a line, containing the type of operation (buy or sell), the number of shares, the unit price of the shares, and the operation date.

Based on the provided input, the application calculates the income tax on capital gains or losses and displays the result.

## Installing Dependencies
At the root of the project, run the following command:
```bash
  go mod tidy 
```

## Running Locally
#### Using a File
At the root of the project, there is a file named `input.txt` containing an example input.

To run the project using this file, execute the following command from the project root:
```bash
  go run main.go < input.txt
```


#### Using Input from stdin
Run the project with the following command and provide each input line by line.

When a blank line is entered, the input reading is finalized, and the operations are calculat
```bash
  go run main.go
```


## Tests
The application uses [Testfy](https://github.com/stretchr/testify) library to better organize and execute tests.


#### Execution
The `_testdata` directory contains all the use cases presented in the challenge.

To run all the tests, execute the following command from the project root:
```bash
  go test -v ./...
```

## Using Docker

Exists a Dockerfile in the root of the project that can be used to build a Docker image of the application.
To build a image use the following command:

```bash
  docker build -t capital-gains .
```

#### Using Input from stdin
Use the following command:
```bash
  docker run -it capital-gains
```
#### Using a File
Use the following command:
```bash
  docker run -i capital-gains < input.txt
```