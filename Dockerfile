FROM golang:1.21.13-alpine3.20

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GIN_MODE=release

RUN mkdir app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

ARG PORT

ENV PORT=$PORT

COPY . .

RUN go build -o main .

CMD ["./main"]