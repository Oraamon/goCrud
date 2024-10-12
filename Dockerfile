FROM golang:1.23.2

WORKDIR /app

COPY go.mod .
COPY . .

RUN go build -o bin ./cmd/api/main.go

ENTRYPOINT ["go", "run", "./cmd/api/main.go"]
