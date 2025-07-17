FROM golang:tip-alpine3.21

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "app.go"]