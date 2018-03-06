package smtpconnect

import (
	"strconv"

	"github.com/go-mail/mail"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Open and return an SMTP Connection
func InitSMTP() (mail.SendCloser, error) {

	dialer := mail.NewDialer(
		viper.GetString("SMTPHost"),
		viper.GetInt("SMTPPort"),
		viper.GetString("SMTPUser"),
		viper.GetString("SMTPPass"),
	)

	conn, err := dialer.Dial()
	if err != nil {
		log.Info("SMTPPort="+strconv.Itoa(viper.GetInt("SMTPPort")))
		log.Info("SMTPHost="+viper.GetString("SMTPHost"))
		log.Info("SMTPUser="+viper.GetString("SMTPUser"))
		log.Info("SMTPUser="+viper.GetString("SMTPPass"))

		log.Error("smtpconnect.InitSMTP: "+err.Error())
		return nil, err
	}
	return conn, nil
}
