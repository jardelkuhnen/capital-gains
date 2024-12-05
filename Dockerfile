FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

ENTRYPOINT ["go", "run", "main.go"]