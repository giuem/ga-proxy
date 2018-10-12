FROM golang:1.11-alpine3.8 AS BUILD

WORKDIR /src

RUN apk --no-cache add git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags "-w -s" -o ga_proxy

FROM alpine:3.8
LABEL maintainer "giuem <i@giuem.com>"

WORKDIR /app

RUN apk --no-cache add ca-certificates curl

COPY --from=BUILD /src/ga_proxy .

EXPOSE 80

HEALTHCHECK --interval=1m --timeout=10s --start-period=1s --retries=2 \
  CMD curl -X HEAD -If http://localhost/detect || exit 1

CMD ["./ga_proxy", "-s"]
