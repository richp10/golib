package smtpconnect_test

import (
	"os"
	"os/exec"
	"testing"
	"github.com/richp10/golib/env"
	"github.com/spf13/viper"
	"github.com/richp10/golib/email/smtpconnect"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSMTPConnection(t *testing.T) {

	Convey("Should create valid SMTP Connection without error", t, func() {

		Convey("Should be reading configuration values", func() {
			port := viper.GetInt("SMTPPort")
			So(port, ShouldEqual, 587)
		})

		Convey("Should create SMTP connection", func() {
			_, err := smtpconnect.InitSMTP()
			So(err, ShouldBeNil)
		})

		Convey("Should pass all MegaChecks", func() {
			cmd := exec.Command("megacheck", "github.com/richp10/golib/email/smtpconnect.go")
			res, _ := cmd.Output()
			So(string(res[:]), ShouldBeEmpty)
		})

	})
}

func TestMegaChecks(t *testing.T) {

	Convey("Should pass megachecks", t, func() {

		cmd := exec.Command("megacheck")
		res, _ := cmd.Output()
		So(string(res[:]), ShouldBeEmpty)

	})
}

func TestMain(m *testing.M) {
	if viper.GetString("TEST") != "asdfadsfasdfasdf" {
		env.Load()
	}
	m.Run()
	code := m.Run()
	os.Exit(code)
}
