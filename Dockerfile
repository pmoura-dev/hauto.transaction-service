# Builder image
FROM golang:1.21-alpine3.18 AS Builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /transaction-service main.go

# Runtime image
FROM alpine:3.18 AS runtime

COPY --from=builder /transaction-service /transaction-service

ENTRYPOINT [ "/transaction-service" ]