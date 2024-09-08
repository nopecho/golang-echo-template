FROM golang:1.23-alpine AS builder

ARG BUILD_PATH
WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o server $BUILD_PATH

FROM alpine:latest

RUN apk --no-cache add tzdata && \
    apk --no-cache add ca-certificates

COPY --from=builder /app /app

EXPOSE $PORT
ENTRYPOINT ["/app/server"]