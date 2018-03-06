// Create random avatar image from gender and username
package randomavatar

import (
	"image"
	"strings"

	"github.com/o1egl/govatar"
)

func Create(gender string, username string) (image.Image, error) {
	var img image.Image
	var err error

	if strings.ToLower(gender) == "male" {
		img, err = govatar.GenerateFromUsername(govatar.MALE, username)
	} else if strings.ToLower(gender) == "female" {
		img, err = govatar.GenerateFromUsername(govatar.FEMALE, username)
	} else {
		img, err = govatar.GenerateFromUsername(govatar.FEMALE, username)
	}
	if err != nil {
		return nil, err
	}

	return img, nil
}
