package dbconnect_test

import (
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMegaChecks(t *testing.T) {

	Convey("Should pass megachecks", t, func() {

		cmd := exec.Command("megacheck")
		res, _ := cmd.Output()
		So(string(res[:]), ShouldBeEmpty)

	})
}
