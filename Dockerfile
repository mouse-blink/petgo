# First stage: build and generate Swagger docs
FROM golang:1.24 AS builder
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /app
COPY . .
RUN go mod download
RUN go generate ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/app /app

ENTRYPOINT ["/app"]
