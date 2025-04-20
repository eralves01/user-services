FROM golang:1.24

RUN go install github.com/air-verse/air@latest

WORKDIR /app

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

RUN go mod tidy

CMD ["air"]