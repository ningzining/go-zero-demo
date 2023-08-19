FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main greet/greet.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/greet/etc/greet-api.yaml /app/etc/greet-api.yaml

EXPOSE 8888

CMD ["/app/main"]