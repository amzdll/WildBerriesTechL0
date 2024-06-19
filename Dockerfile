FROM golang:1.22.3-alpine3.20
#  AS builder

COPY . /project/wb
WORKDIR /project/wb

RUN apk update && apk upgrade && apk add --no-cache alpine-sdk
RUN apk add zsh postgresql-client jq


RUN go mod tidy
RUN go mod download
CMD zsh
#CMD go run cmd/main.go

#FROM alpine:3.20
#
#WORKDIR /root/
#COPY --from=builder /shop/bin/app .
#CMD ["./app"]