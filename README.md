# ssid

![reviewdog](https://github.com/shmokmt/ssid/workflows/reviewdog/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/shmokmt/ssid?tab=doc)](https://pkg.go.dev/github.com/shmokmt/ssid?tab=doc)

A tiny go package to get the SSID of the Wi-Fi you are connected to.

## Runtime

- Go 1.14 or later
- macOS（darwin）

## Installation

```
go get github.com/shmokmt/ssid
```

## Usage

```go
package main

import (
    "github.com/shmokmt/ssid"
)

func main() {
    fmt.Println(ssid.String())
}
```

## References

- [yelinaung/wifi-name](https://github.com/yelinaung/wifi-name)

## LICENSE

MIT
