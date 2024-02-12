FROM golang:1.22.0 AS builder

WORKDIR /usr/local/app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ./main ./cmd/api/main.go 

FROM scratch
COPY --from=builder /usr/local/app/main /usr/local/app/.env ./
CMD ["./main"]
