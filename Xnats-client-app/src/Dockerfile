FROM golang:1.9.0-alpine3.6 AS builder
COPY . /go/src/github.com/syedomair/api_micro/nats-client-app/src
WORKDIR /go/src/github.com/syedomair/api_micro/nats-client-app/src
RUN apk add --update git
RUN go get -u github.com/nats-io/go-nats
RUN go get -u github.com/nats-io/nuid
RUN CGO_ENABLED=0 go build -o nats-client-app -v -a ./main.go

FROM scratch
COPY --from=builder /go/src/github.com/syedomair/api_micro/nats-client-app/src/nats-client-app /nats-client-app
ENTRYPOINT ["/nats-client-app"]

