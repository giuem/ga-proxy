FROM alpine:3.6
LABEL maintainer "giuem <i@giuem.com>"

WORKDIR /app

RUN apk --no-cache add curl wget && \
    curl -sS https://api.github.com/repos/giuem/ga-proxy/releases/latest | grep -Eo 'https(.+?)linux_386' | wget -nv -i - -O ga_proxy && \
    chmod +x ga_proxy && \
    apk del curl wget

EXPOSE 80

CMD ["./ga_proxy", "-s"]
