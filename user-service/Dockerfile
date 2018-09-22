FROM golang:1.11.0 as builder

WORKDIR /go/src/github.com/skyarkitekten/go-microservices/user-service

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/skyarkitekten/go-microservices/user-service/user-service .

ENTRYPOINT ["./user-service"]