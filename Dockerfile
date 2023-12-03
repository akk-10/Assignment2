# Build Stage
FROM golang:alpine as builder

RUN apk update && \
    apk add --no-cache git bash build-base postgresql-client

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o myapp ./cmd/api

# Final Stage
FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/myapp .

EXPOSE 4000
CMD ["./myapp"]
