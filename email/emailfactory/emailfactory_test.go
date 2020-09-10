package emailfactory_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/richp10/golib/email/emailfactory"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/richp10/golib/env"

	"github.com/matcornic/hermes/v2"
	"github.com/spf13/viper"
)

func TestSMTPConnection(t *testing.T) {

	Convey("Should use email factory to create HTML and text versions of email", t, func() {
		email := hermes.Email{
			Body: hermes.Body{
				Title: "TEST EMAIL",
				Intros: []string{
					`THIS IS A TEST EMAIL`,
				},
				Actions: []hermes.Action{
					{
						Button: hermes.Button{
							Color: "#DC4D2F", // Red Button
							Text:  "Big Red Button",
							Link:  "https://example.org",
						},
					},
				},
				Signature: "Regards",
			},
		}

		_, _, err := emailfactory.Make(email, "")
		So(err, ShouldBeNil)

		Convey("Should pass all MegaChecks", func() {
			cmd := exec.Command("megacheck", "github.com/richp10/golib/email/emailfactory")
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
		env.Load("")
	}
	m.Run()
	code := m.Run()
	os.Exit(code)
}
