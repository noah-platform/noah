FROM golang:1.25-alpine AS builder

ARG APP_DIR

ENV CGO_ENABLED=0

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/swaggo/swag/v2/cmd/swag@v2.0.0-rc4
RUN go install github.com/actgardner/gogen-avro/v10/cmd/...@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

WORKDIR $APP_DIR

RUN go generate ./...
RUN go build -o main .

####################################################################################################

FROM alpine:3.20

ARG APP_DIR

WORKDIR /app

COPY --from=builder /app/$APP_DIR/main .
COPY --from=builder /app/$APP_DIR/generated/docs ./generated/docs

CMD ["./main"]
