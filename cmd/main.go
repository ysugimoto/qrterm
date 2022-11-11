package main

import (
	"github.com/ysugimoto/qrterm"
	"io"
	"os"
)

func main() {
	if err := qrterm.RunApp(); err != nil {
		if _, ok := err.(qrterm.ErrNotFound); ok {
			// nolint
			io.WriteString(os.Stderr, "QR Code is not found in captured image.\n")
		} else {
			// nolint
			io.WriteString(os.Stderr, err.Error()+"\n")
		}
		os.Exit(1)
	}
}
