FROM golang:1.22.3-alpine3.20 AS builder

COPY . /wb
WORKDIR /wb

RUN apk update && apk upgrade
RUN go mod tidy
RUN go mod download
RUN go build -o ./bin/app cmd/orders/main.go

FROM alpine:3.20

WORKDIR /root/

RUN apk add jq

COPY --from=builder /wb/config/.dev.yaml ./config/.dev.yaml
COPY --from=builder /wb/bin/app .

CMD ["./app", "--stage=dev", "|", "jq"]
