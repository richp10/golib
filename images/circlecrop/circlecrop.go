// Circle crop a square image with transparency
// From https://blog.golang.org/go-imagedraw-package
// Golang BSD Licensed code snippet
package circlecrop

import (
	"image"
	"image/color"
	"image/draw"
)

type circle struct {
	p image.Point
	r int
}

func Go(img image.Image) image.Image {
	p := image.Point{
		X: img.Bounds().Max.X / 2,
		Y: img.Bounds().Max.Y / 2,
	}
	r := img.Bounds().Max.X / 2

	dst := image.NewRGBA(img.Bounds())
	draw.DrawMask(dst, img.Bounds(), img, image.ZP, &circle{p, r}, image.ZP, draw.Over)
	return dst
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		//noinspection GoStructInitializationWithoutFieldNames
		return color.Alpha{255}
	}
	//noinspection GoStructInitializationWithoutFieldNames
	return color.Alpha{0}
}
