FROM golang:1.23-alpine AS builder
ARG BUILD_PATH
ARG APP_NAME

WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o $APP_NAME $BUILD_PATH

FROM alpine:latest
ARG APP_NAME
ARG APP_PATH=/app
ARG RUN_PATH=$APP_PATH/$APP_NAME

ENV APP_NAME=$APP_NAME \
    RUN_PATH=$RUN_PATH \

RUN apk --no-cache add tzdata && \
    apk --no-cache add ca-certificates

COPY --from=builder /app/$APP_NAME $RUN_PATH

EXPOSE $PORT
ENTRYPOINT ["sh", "-c", "$RUN_PATH"]