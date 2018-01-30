package utils

// modify from https://github.com/giuem/realip/blob/master/realip.go

import (
  "errors"
  "net"
  "strings"

  "github.com/valyala/fasthttp"
)

var cidrs []*net.IPNet

func init() {
  maxCidrBlocks := []string{
    "127.0.0.1/8",    // localhost
    "10.0.0.0/8",     // 24-bit block
    "172.16.0.0/12",  // 20-bit block
    "192.168.0.0/16", // 16-bit block
    "169.254.0.0/16", // link local address
    "::1/128",        // localhost IPv6
    "fc00::/7",       // unique local address IPv6
    "fe80::/10",      // link local address IPv6
  }

  cidrs = make([]*net.IPNet, len(maxCidrBlocks))
  for i, maxCidrBlock := range maxCidrBlocks {
    _, cidr, _ := net.ParseCIDR(maxCidrBlock)
    cidrs[i] = cidr
  }
}
func isPrivateAddress(address string) (bool, error) {
  ipAddress := net.ParseIP(address)
  if ipAddress == nil {
    return false, errors.New("address is not valid")
  }

  for i := range cidrs {
    if cidrs[i].Contains(ipAddress) {
      return true, nil
    }
  }

  return false, nil
}

// Realip return realip ip
func Realip(ctx *fasthttp.RequestCtx) string {
  if xffb := ctx.Request.Header.Peek("X-Forwarded-For"); len(xffb) > 0 {
    xff := string(xffb)
    for _, address := range strings.Split(xff, ",") {
      address = strings.TrimSpace(address)
      isPrivate, err := isPrivateAddress(address)
      if !isPrivate && err == nil {
        return address
      }
    }
  }

  if xripb := ctx.Request.Header.Peek("X-Real-Ip"); len(xripb) > 0 {
    return string(xripb)
  }

  return ctx.RemoteIP().String()
}
