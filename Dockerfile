# Download mod
FROM golang:1.23-alpine AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Development Stage
## Includes tools for hot-reloading & debugging
FROM base AS dev
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY . .
EXPOSE 8080
EXPOSE 2345
CMD ["air", "-c", ".air.toml"]

# Production Stage
## Optimized build
FROM base AS prod
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/main ./cmd/api
## Runtime
FROM alpine:latest AS runtime
WORKDIR /root
COPY --from=prod /app/main .
EXPOSE 8080
CMD ["./main"]
