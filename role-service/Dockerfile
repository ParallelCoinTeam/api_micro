
FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/syedomair/api_micro/role-service

COPY src .
COPY proto proto
COPY vendor vendor
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -o role-service -a -installsuffix cgo main.go repository.go handler.go database.go 

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/syedomair/api_micro/role-service .

EXPOSE 8180

CMD ["./role-service"]
