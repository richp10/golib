package imagetools

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"gopkg.in/h2non/filetype.v1"

	. "github.com/smartystreets/goconvey/convey"
)

func TestImageTools(t *testing.T) {

	gender := "Male"
	username := "bob"
	file, _ := ioutil.TempFile(os.TempDir(), "prefix")
	defer os.Remove(file.Name())

	Convey("Should be able to create Random Avatar", t, func() {
		err := CreateRandomAvatar(gender, username)
		So(err, ShouldBeNil)

		Convey("Should be able to Resize the image", func() {
			beforewidth := img.Bounds().Max.X
			ResizeSquare(100)
			afterwidth := img.Bounds().Max.X

			So(beforewidth, ShouldNotEqual, 100)

			So(afterwidth, ShouldEqual, 100)
		})

		// I don't know how to programatically check if circlecrop
		// works..

		Convey("Should be able to save image to disk as png", func() {
			err = SavePNG(file.Name())
			So(err, ShouldBeNil)
		})

		Convey("Should be able reduce image size with pngquant", func() {
			// Get size before crush
			file, err := os.Open(file.Name())
			So(err, ShouldBeNil)
			fi, err := file.Stat()
			So(err, ShouldBeNil)
			before := fi.Size()
			file.Close()

			PNGQuant()
			err = SavePNG(file.Name())
			So(err, ShouldBeNil)
			file, err = os.Open(file.Name())
			So(err, ShouldBeNil)
			fi, err = file.Stat()
			So(err, ShouldBeNil)
			after := fi.Size()
			file.Close()

			compare := before == after
			So(compare, ShouldBeFalse)
		})

		// Could not think of a good assertion - at least this checks
		// the code does not panic
		Convey("Should be able to circlecrop without error", func() {
			CircleCrop()
		})

		Convey("Should be able to find pngquant in the path", func() {
			_, err := exec.LookPath("pngquant")
			So(err, ShouldBeNil)
		})

		Convey("Should be able to find pngout in the path", func() {
			_, err := exec.LookPath("pngout")
			So(err, ShouldBeNil)
		})

		Convey("Should be able to smartcrop without error", func() {
			err = SmartCrop(50, 50)
			So(err, ShouldBeNil)
		})

		Convey("Created image should be a PNG", func() {
			buf, _ := ioutil.ReadFile(file.Name())
			isimg := filetype.IsImage(buf)
			So(isimg, ShouldBeTrue)
			ispng := filetype.IsMIMESupported("image/png")
			So(ispng, ShouldBeTrue)
		})

		Convey("Should pass all MegaChecks", func() {
			cmd := exec.Command("megacheck", "github.com/richp10/golib/images/smartcrop")
			res, _ := cmd.Output()
			So(string(res[:]), ShouldBeEmpty)
		})

	})

}
