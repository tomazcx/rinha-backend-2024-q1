FROM golang:1.22.0

WORKDIR /usr/local/app

COPY . .

RUN go mod tidy
RUN go build -o ./main ./cmd/api/main.go 

CMD ["./main"]
