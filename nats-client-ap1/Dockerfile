
FROM golang:1.9.0 AS builder

WORKDIR /go/src/github.com/syedomair/api_micro/nats-client-ap1

COPY src .

COPY vendor vendor
#RUN go get -u github.com/golang/dep/cmd/dep
#RUN dep init && dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -o nats-client-ap1 -a -installsuffix cgo main.go 

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/syedomair/api_micro/nats-client-ap1/nats-client-ap1 /nats-client-ap1

EXPOSE 8180

CMD ["./nats-client-ap1"]
