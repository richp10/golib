package sendemail

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/go-mail/mail"
	"github.com/matcornic/hermes/v2"
	"github.com/richp10/golib/email/emailfactory"
	"github.com/richp10/golib/email/smtpconnect"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Send email using hermes Email struct.
func Send(from string, to string, subject string, body hermes.Email, conn mail.SendCloser, logo string) error {
	var err error

	// If connection has not been set, create it now..
	// We allow it to be passed so we can re-use the
	// connection if Send is used within a loop
	// such as sending newsletters
	if conn == nil {
		conn, err = smtpconnect.InitSMTP()
		if err != nil {
			log.Error("sendemail.Send: " + err.Error())
			return err
		}
	}

	if from == "" {
		from = viper.GetString("SMTPFrom")
	}

	// Convert hermes email into html and plain text..
	emailBody, emailText, _ := emailfactory.Make(body, logo)

	// Set X-Entity-Ref-ID Header to random value prevents Gmail Threading
	rand.Seed(time.Now().Unix())

	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetHeader("X-Entity-Ref-ID", strconv.Itoa(rand.Intn(10000)))
	m.SetBody("text/plain", emailText)
	m.AddAlternative("text/html", emailBody)

	return mail.Send(conn, m)
}
