package sendemail

import (
	"os"
	"os/exec"
	"testing"

	"github.com/richp10/golib/email/emailfactory"
	"github.com/richp10/golib/email/smtpconnect"
	"github.com/richp10/golib/env"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/spf13/viper"

	"github.com/matcornic/hermes"
)

func TestSMTPConnection(t *testing.T) {

	Convey("Should send test email", t, func() {
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

		Convey("Should send test email using supplied connection", func() {
			to := "richpt10@gmail.com"
			_, _, err := emailfactory.Make(email)
			So(err, ShouldBeNil)

			err = Send("", to, "TEST EMAIL", email, nil)
			So(err, ShouldBeNil)

			conn, err := smtpconnect.InitSMTP()
			So(err, ShouldBeNil)
			err = Send("", to, "TEST EMAIL 2", email, conn)
			So(err, ShouldBeNil)

		})

		Convey("Should pass all MegaChecks", func() {
			cmd := exec.Command("megacheck", "github.com/richp10/golib/email/sendemail")
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
