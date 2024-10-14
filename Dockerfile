FROM golang:1.23.2 as DEV

WORKDIR /app

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod .
COPY . .

RUN go build -o bin ./cmd/api/main.go

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
