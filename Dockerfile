FROM golang:1.11-alpine AS BUILD

WORKDIR /src

RUN apk --no-cache add git ca-certificates tzdata && update-ca-certificates

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ga_proxy

FROM alpine:3.8
LABEL maintainer "giuem <i@giuem.com>"

COPY --from=BUILD /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=BUILD /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILD /src/ga_proxy /ga_proxy

EXPOSE 80

# HEALTHCHECK --interval=1m --timeout=10s --start-period=1s --retries=2 \
#   CMD curl -X HEAD -If http://localhost/detect || exit 1

CMD ["/ga_proxy", "-s"]
