# qrterm

`qrterm` is terminal command that read qrcode through the ScreenCapture app for osx.

## Installation

```shell
go install github.com/ysugimoto/qrterm/cmd
```

Or download binary from [Releases]() page.

## Usage

Before use this command, make sure your terminal is permitted of `Screen Recording`. Confirm your Security and Privacy setting.

<img width="780" alt="Screen Shot 2022-11-11 at 9 13 25" src="https://user-images.githubusercontent.com/1000401/201233848-4890b6da-dd20-4928-aef6-2eb590966618.png">


### Terminal App

Simply type `qrterm` command in your terminal, then runs ScreenCapture app and capture target QR Code.
After captured, read QR data and open dialog.

<img width="532" alt="Screen Shot 2022-11-11 at 9 18 15" src="https://user-images.githubusercontent.com/1000401/201233873-e9be1e5b-ab42-4ec6-bd83-9c0bf740f1f6.png">

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
    return
  }
  log.Println(data)
}
```

## Contribution

- Fork this repository
- Customize / Fix problem
- Send PR :-)
- Or feel free to create issues for us. We'll look into it

## License

MIT License

