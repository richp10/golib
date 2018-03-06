package circlecrop

import (
	"image"
	"os/exec"
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

	Convey("Should be able to circlecrop an image withoug error", t, func() {

		img = Go(img)
		// Can't work out how to check this yet..

		Convey("Should pass all MegaChecks", func() {
			cmd := exec.Command("megacheck", "github.com/richp10/golib/images/circlecrop")
			res, _ := cmd.Output()
			So(string(res[:]), ShouldBeEmpty)
		})

	})

}

