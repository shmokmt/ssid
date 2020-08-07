# ssid

A simple go package to get the SSID of the Wi-Fi you are connected to

## Installation

```go
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
