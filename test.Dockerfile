
FROM golang:1.19 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app", "go test -tags=e2e -v ./..."]
