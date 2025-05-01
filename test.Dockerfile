FROM golang:1.24 as test-runner
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /app
COPY . .
RUN go mod download
RUN go generate ./...

RUN go mod init testmod || true
RUN go mod tidy

CMD ["go", "test", "-v", "./..."]
