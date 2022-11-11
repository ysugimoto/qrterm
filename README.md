# qrterm

`qrterm` is terminal command that read qrcode through the ScreenCapture app for osx.

## Installation

```shell
go install github.com/ysugimoto/qrterm/cmd
```

Or download binary from [Releases]() page.

## Usage

Before use this command, make sure your terminal is permitted of `Screen Recording`. Confirm your Security and Privacy setting.


### Terminal App

Simply type `qrterm` command in your terminal, then runs ScreenCapture app and capture target QR Code.
After captured, read QR data and open dialog.

### Program

This package also can use in your Go program.

```go
package main

import (
  "log"
  "github.com/ysugimoto/qrterm"
)

func main() {
  data, err := qrterm.RunProgram()
  if err != nil {
    if _, ok := err.(qrterm.ErrNotFound); ok {
      log.Println("QR Code is not found in captured image.")
      return
    }
    log.Println(err)
  }
}
```

## Contribution

- Fork this repository
- Customize / Fix problem
- Send PR :-)
- Or feel free to create issues for us. We'll look into it

## License

MIT License

