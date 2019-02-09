FROM golang:1.11-alpine AS BUILD

WORKDIR /src

# module
RUN apk --no-cache add git ca-certificates tzdata && update-ca-certificates
COPY go.mod go.sum ./
RUN go mod download

# build
COPY ga ga
COPY server server
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ga-proxy

FROM alpine:3.8
LABEL maintainer "giuem <giuemcom+docker@gmail.com>"
EXPOSE 80
ENV IP=0.0.0.0
ENV PORT=80
ENV GIN_MODE=release

COPY --from=BUILD /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=BUILD /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILD /src/ga-proxy /ga-proxy

HEALTHCHECK --interval=1m --timeout=10s --start-period=1s --retries=2 \
  CMD /ga-proxy ping

CMD ["/ga-proxy"]
