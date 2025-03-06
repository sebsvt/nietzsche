FROM golang:1.23-bullseye

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o nietzsche cmd/main.go

CMD ["/app/nietzsche"]
