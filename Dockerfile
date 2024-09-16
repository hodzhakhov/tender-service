FROM mirror.gcr.io/golang:1.22-alpine as builder

WORKDIR /src
COPY . .

RUN go build -o app ./cmd/main.go

EXPOSE 8080

CMD ["./app"]
