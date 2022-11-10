package main

import (
	"log"
	"bytes"

	"image/png"
	"os/exec"
	"strings"

	"github.com/ncruces/zenity"
	"golang.design/x/clipboard"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func init() {
	if err := clipboard.Init(); err != nil {
		panic(err)
	}
}

func main() {
	if err := exec.Command("screencapture", "-S", "-i", "-c", "-t", "png").Run(); err != nil {
		log.Fatalln("exec error", err)
	}
	img := clipboard.Read(clipboard.FmtImage)
	if img == nil {
		log.Fatalln("clipboard reader error, image is nil")
	}

	p, err := png.Decode(bytes.NewReader(img))
	if err != nil {
		log.Fatalln("image decode error", err)
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(p)
	if err != nil {
		log.Fatalln("convert bitmap error", err)
	}
	reader := qrcode.NewQRCodeReader()
	ret, err := reader.Decode(bmp, nil)
	if err != nil {
		if _, ok :=  err.(gozxing.NotFoundException); ok {
			println("QR Code not found in capture")
			return
		}
		log.Fatalln("qr decode error", err)
	}

	if strings.HasPrefix(ret.String(), "http://") || strings.HasPrefix(ret.String(), "https://") {
		err = zenity.Question(ret.String(),
			zenity.Title("Open Browser"),
			zenity.NoIcon,
			zenity.OKLabel("Open"),
		)

		if err == nil {
			exec.Command("open", ret.String()).Run()
		}
	} else {
		clipboard.Write(clipboard.FmtText, []byte(ret.String()))
		zenity.Notify("Saved to clipboard\n" + ret.String(),
			zenity.Title("Text Found"),
			zenity.InfoIcon)
	}
}
