FROM golang:1.25-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o collector-service ./cmd/collector-service

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/collector-service .
EXPOSE 9090
ENTRYPOINT ["./collector-service"]
