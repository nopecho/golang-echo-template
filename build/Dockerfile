FROM golang:1.23-alpine AS builder
ARG BUILD_PATH
ARG APP_NAME=server

WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o $APP_NAME $BUILD_PATH

FROM alpine:latest
ARG APP_NAME=server
ENV APP_NAME=$APP_NAME

RUN apk --no-cache add tzdata && \
    apk --no-cache add ca-certificates

COPY --from=builder /app /app

EXPOSE $PORT
ENTRYPOINT ["sh", "-c", "/app/$APP_NAME"]