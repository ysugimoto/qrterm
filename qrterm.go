package qrterm

import (
	"bytes"
	"fmt"

	"image/png"
	"os/exec"
	"strings"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/ncruces/zenity"
	"golang.design/x/clipboard"
)

// re-expose NotFoundException that QR code is not found in captured image
type ErrNotFound = gozxing.NotFoundException

// Must initialize clipboard
func init() {
	if err := clipboard.Init(); err != nil {
		panic(err)
	}
}

// Run as application. Display dialog and open browser
func RunApp() error {
	result, err := RunProgram()
	if err != nil {
		return err
	}

	if strings.HasPrefix(result, "http://") || strings.HasPrefix(result, "https://") {
		// If read data starts with http[s]://, dialog to open URL with browser
		if err = zenity.Question(result, zenity.Title("Open Browser"), zenity.NoIcon, zenity.OKLabel("Open")); err == nil {
			if err := exec.Command("open", result).Run(); err != nil {
				return fmt.Errorf("Failed to open default browser: %s", err)
			}
		}
	} else {
		// Otherwise, write to clipboard
		clipboard.Write(clipboard.FmtText, []byte(result))
		if err := zenity.Notify("Saved to clipboard\n"+result, zenity.Title("Text Found"), zenity.InfoIcon); err != nil {
			return fmt.Errorf("Failed to open dialog: %s", err)
		}
	}

	return nil
}

// Run capture program and return read data
func RunProgram() (string, error) {
	fmt.Println("Drag cursor and capture target QR Code...")
	if err := exec.Command("screencapture", "-S", "-i", "-c", "-t", "png").Run(); err != nil {
		return "", fmt.Errorf("Failed to execute screencapture: %w", err)
	}
	img := clipboard.Read(clipboard.FmtImage)
	if img == nil {
		return "", fmt.Errorf("Failed to read from clipboard")
	}

	p, err := png.Decode(bytes.NewReader(img))
	if err != nil {
		return "", fmt.Errorf("Failed to decode png image: %w", err)
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(p)
	if err != nil {
		return "", fmt.Errorf("Failed to create bitmap image: %w", err)
	}
	reader := qrcode.NewQRCodeReader()
	ret, err := reader.Decode(bmp, nil)
	if err != nil {
		return "", err
	}
	return ret.String(), nil
}
