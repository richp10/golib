package crushpng

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/richp10/golib/images/randomavatar"

	"github.com/anthonynsimon/bild/imgio"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCrushPNG(t *testing.T) {

	img, _ := randomavatar.Create("Male", "bob")
	file, _ := ioutil.TempFile(os.TempDir(), "prefix")
	defer os.Remove(file.Name())
	err := imgio.Save(file.Name(), img, imgio.PNGEncoder())
	if err != nil {
		panic(err)
	}

	Convey("Should be able reduce image size with pngquant", t, func() {
		// Get size before crush
		file, err := os.Open(file.Name())
		So(err, ShouldBeNil)
		fi, err := file.Stat()
		So(err, ShouldBeNil)
		before := fi.Size()
		file.Close()

		img, err := PNGQuant(img)
		So(err, ShouldBeNil)

		err = imgio.Save(file.Name(), img, imgio.PNGEncoder())
		So(err, ShouldBeNil)
		file, err = os.Open(file.Name())
		So(err, ShouldBeNil)
		fi, err = file.Stat()
		So(err, ShouldBeNil)
		after := fi.Size()
		file.Close()

		compare := before == after
		So(compare, ShouldBeFalse)

		Convey("Should pass all MegaChecks", func() {
			cmd := exec.Command("megacheck", "github.com/richp10/golib/images/crushpng")
			res, _ := cmd.Output()
			So(string(res[:]), ShouldBeEmpty)
		})

	})

}

