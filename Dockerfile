FROM golang:1.19 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]
