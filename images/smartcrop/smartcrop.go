// Crop an image with intelligent content selection
package smartcrop

import (
	"image"
	_ "image/png"

	"github.com/muesli/smartcrop"
)

func Crop(img image.Image, width int, height int) (image.Image, error) {

	analyzer := smartcrop.NewAnalyzer()
	topCrop, err := analyzer.FindBestCrop(img, width, height)
	if err != nil {
		return nil, err
	}

	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	croppedimg := img.(SubImager).SubImage(topCrop)
	return croppedimg, nil
}
