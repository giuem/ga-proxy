# ga-proxy

Speed up Google Analytics



## Get Start

### 1. Install 

Download binary from [release](https://github.com/giuem/ga-proxy/releases) or build yourself.

### 2. Run

```
./ga_proxy [arguments]
```

options:

```
-d, -debug output debug info.
-s, -skip_ssl skip ssl verify. (some envirinment don't have certificate)
-l, -listen listen address, default is :80
```




## Build

1. Get requirement

   ```bash
   go get github.com/mitchellh/gox github.com/tomasen/realip github.com/satori/go.uuid 
   ```


2. Build

   ```bash
   gox -os="!freebsd !netbsd !openbsd" -ldflags="-s -w" -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}"
   ```


