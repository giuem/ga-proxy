FROM golang:alpine3.6 AS BUILD

WORKDIR /go/src/github.com/giuem/ga_proxy

COPY . .

RUN  apk add --no-cache git

RUN go get -v && \
  go build -ldflags "-w -s" -o ga_proxy

FROM alpine:3.6
LABEL maintainer "giuem <i@giuem.com>"

WORKDIR /app

RUN apk --no-cache add ca-certificates curl

COPY --from=BUILD /go/src/github.com/giuem/ga_proxy/ga_proxy .

EXPOSE 80

HEALTHCHECK --interval=1m --timeout=10s --start-period=1s --retries=2 \
  CMD curl -X HEAD -If http://localhost/detect || exit 1

CMD ["./ga_proxy", "-s"]
