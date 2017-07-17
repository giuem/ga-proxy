# ga-proxy

Speed up Google Analytics



## Build

1. Get `gox`

```bash
go get github.com/mitchellh/gox
```

2. Build

```bash
gox -ldflags="-s -w"  -output="build/{{.Dir}}{{.OS}}{{.Arch}}"
```

