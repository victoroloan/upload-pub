FROM golang:alpine

RUN apk update && apk add --no-cache git

ENV GIN_MODE=release
WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o upload

ENTRYPOINT ["/app/upload"]
