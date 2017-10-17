FROM golang:alpine3.6 AS BUILD

WORKDIR /go/src/github.com/giuem/ga_proxy

COPY . .

RUN  apk add --no-cache git

RUN go get -v && \
  go build -ldflags "-w -s" -o ga_proxy

FROM alpine:3.6
LABEL maintainer "giuem <i@giuem.com>"

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=BUILD /go/src/github.com/giuem/ga_proxy/ga_proxy .

EXPOSE 80

CMD ["./ga_proxy", "-s"]
