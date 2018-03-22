
FROM golang:1.9.0 AS builder

COPY . /go/src/github.com/syedomair/api_micro/nats-client-ap1/src
WORKDIR /go/src/github.com/syedomair/api_micro/nats-client-ap1/src
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -o nats-client-ap1 -a -installsuffix cgo ./main.go

FROM scratch

FROM alpine:latest

RUN apk --no-cache add ca-certificates



COPY --from=builder /go/src/github.com/syedomair/api_micro/nats-client-ap1/src/nats-client-ap1 /nats-client-ap1

EXPOSE 8180

ENTRYPOINT ["./nats-client-ap1"]

