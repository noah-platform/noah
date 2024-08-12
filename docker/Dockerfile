FROM golang:1.22-alpine AS builder

ARG APP_DIR

ENV CGO_ENABLED=0

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR $APP_DIR

RUN go generate ./...

RUN go build -o main .

####################################################################################################

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/$APP_DIR/main .

CMD ["./main"]
