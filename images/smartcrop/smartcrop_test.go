package smartcrop

import (
	"image"
	"testing"

	"github.com/richp10/golib/images/randomavatar"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	err error
	img image.Image
)

func TestCircleCrop(t *testing.T) {

	img, err = randomavatar.Create("Male", "bob")
	if err != nil {
		panic(err)
	}

	Convey("Should be able to smartcrop without error", t, func() {
		_, err = Crop(img, 50, 50)
		So(err, ShouldBeNil)
	})

}
