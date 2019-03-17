# ar150gpio
A Go package specifically for accessing the GPIOs of the GL.iNet AR-150 router running OpenWrt.

## Installation
Just use go get:
```
go get github.com/iketsj/ar150gpio
```

## Example Use
Note that the router only has 4 free GPIOs(1, 14, 16, 17).

```
package main

import (
	gpio "github.com/iketsj/ar150gpio"
	"fmt"
)

func main() {
	led := gpio.NewGPIO()
	led.Initialize(1, gpio.OUT)

	button := gpio.NewGPIO()
	button.Initialize(14, gpio.IN)

	for true {
		if button.Read() == gpio.HIGH {
			led.Write(gpio.HIGH)
			fmt.Println(gpio.HIGH)
		}else {
			led.Write(gpio.LOW)
			fmt.Println(gpio.LOW)
		}
	}
}
```

## Building the source code
Do not forget to set the GOOS and GOARCH environment variable:
```
GOOS=linux GOARCH=mips go build main.go
```

## Transferring the Binary
You could by using scp:
```
scp main root@192.168.8.1:/
```

## Links
* https://oldwiki.archive.openwrt.org/doc/hardware/port.gpio
* https://www.gl-inet.com/products/gl-ar150/

