FROM golang:1.18-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY app/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM scratch

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]