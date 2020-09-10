package randomavatar_test

import (
	"os/exec"
	"testing"

	"github.com/richp10/golib/images/randomavatar"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRandomAvatar(t *testing.T) {

	gender := "Male"
	username := "bob"

	Convey("Create Random Male Avatar", t, func() {
		_, err := randomavatar.Create(gender, username)
		So(err, ShouldBeNil)

		Convey("Create Random Female Avatar", func() {
			_, err := randomavatar.Create("Female", username)
			So(err, ShouldBeNil)

		})
		// Todo maybe add tests to check valid image..

		Convey("Should pass all MegaChecks", func() {
			cmd := exec.Command("megacheck", "github.com/richp10/golib/images/randomavatar")
			res, _ := cmd.Output()
			So(string(res[:]), ShouldBeEmpty)
		})

	})
}
