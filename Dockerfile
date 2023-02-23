FROM golang:1.19-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache make

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o app ./internal/main.go ./internal/router.go

EXPOSE 8000

ENV TZ EUROPE/MOSCOW

CMD ["./app"]