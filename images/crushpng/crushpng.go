// Wrapper for pngquant - based on https://github.com/yusukebe/go-pngquant
// MIT Licensed
package crushpng

import (
	"bytes"
	"image"
	"image/png"
	"os/exec"
	"strings"
)

// Compress with PNGQuant
func PNGQuant(img image.Image) (image.Image, error) {
	img, err := compress(img, "3")
	if err != nil {
		return nil, err
	}
	return img, nil
}

func compress(input image.Image, speed string) (output image.Image, err error) {
	var w bytes.Buffer
	err = png.Encode(&w, input)
	if err != nil {
		return nil, err
	}

	b := w.Bytes()
	compressed, err := compressBytes(b, speed)
	if err != nil {
		return nil, err
	}

	output, err = png.Decode(bytes.NewReader(compressed))
	return output, err
}

func compressBytes(input []byte, speed string) (output []byte, err error) {
	cmd := exec.Command("pngquant", "-", "--speed", speed, "--quality=75-90")
	cmd.Stdin = strings.NewReader(string(input))
	var o bytes.Buffer
	cmd.Stdout = &o

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	output = o.Bytes()
	return output, nil
}
