// Convenience wrapper for our image manipulation tools
package imagetools

import (
	"image"

	"github.com/richp10/golib/images/circlecrop"
	"github.com/richp10/golib/images/crushpng"
	"github.com/richp10/golib/images/randomavatar"
	"github.com/richp10/golib/images/smartcrop"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

var (
	Img image.Image
)

/*
func Load(filename string) {
	var err error
	Img, err = imgio.Open(filename)
	if err != nil {
		panic(err)
	}
}
*/

func SavePNG(target string) error {
	if err := imgio.Save(target, Img, imgio.PNGEncoder()); err != nil {
		return err
	}
	return nil
}

/*
func SavePNGTemp() (file *os.File, err error) {
	file, err = ioutil.TempFile(os.TempDir(), "prefix")
	if err != nil {
		return nil, err
	}

	if err := imgio.Save(file.Name(), Img, imgio.PNGEncoder()); err != nil {
		return nil, err
	}
	return file, nil
}
*/

func PNGQuant() error {
	var err error
	Img, err = crushpng.PNGQuant(Img)
	if err != nil {
		return err
	}
	return nil
}

func CircleCrop() {
	Img = circlecrop.Go(Img)
}

func CreateRandomAvatar(gender string, username string) error {
	var err error
	Img, err = randomavatar.Create(gender, username)
	if err != nil {
		return err
	}
	return nil
}

func SmartCrop(width int, height int) error {
	var err error
	Img, err = smartcrop.Crop(Img, width, height)
	if err != nil {
		return err
	}
	return nil
}

// Resize to square
func ResizeSquare(size int) {
	Img = transform.Resize(Img, size, size, transform.Linear)
}

func Resize(width int, height int) {
	Img = transform.Resize(Img, width, height, transform.Linear)
}

func Set(i image.Image) {
	Img = i
}

func Get() image.Image {
	return Img
}
