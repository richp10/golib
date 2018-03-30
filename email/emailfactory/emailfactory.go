// Uses Hermes to create the plain text and html version of
// the email - ready to be sent using sendemail package
package emailfactory

import (
	"github.com/matcornic/hermes"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Make(email hermes.Email, logo string) (body string, text string, err error) {

	if logo == "" {
		logo = viper.GetString("EmailLogo")
	}

	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Header & footer of e-mails
			Name:      viper.GetString("EmailProjName"),
			Link:      viper.GetString("EmailLink"),
			Logo:      logo,
			Copyright: viper.GetString("EmailCopyright"),
		},
	}
	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		log.Error("emailfactory.Make html: " + err.Error())
		return "", "", err
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	emailText, err := h.GeneratePlainText(email)
	if err != nil {
		log.Error("emailfactory.Make plain: " + err.Error())
		return "", "", err

	}

	return emailBody, emailText, nil
}
