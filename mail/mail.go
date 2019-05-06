package mail

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/gotopia/more/config"
	"github.com/pkg/errors"
)

// Send the mail.
func Send(from string, to []string, cc []string, bcc []string, subject string, content string) error {
	return send(from, to, cc, bcc, subject, content)
}

func send(from string, to []string, cc []string, bcc []string, subject string, content string) error {
	msg := msg(from, to, cc, bcc, subject, content)
	scfg := config.Mail.SMTP
	addr := fmt.Sprintf("%v:%v", scfg.Host(), scfg.Port())
	acfg := config.Mail.SMTP.Auth
	var auth smtp.Auth
	switch acfg.Type() {
	case "none":
	case "plain":
		auth = smtp.PlainAuth(acfg.Username(), acfg.Username(), acfg.Password(), scfg.Host())
	default:
		return errors.New("unsupported auth type")
	}
	return errors.Wrapf(smtp.SendMail(addr, auth, from, append(to, append(cc, bcc...)...), msg), "failed to send mail")
}

func msg(from string, to []string, cc []string, bcc []string, subject string, content string) []byte {
	headers := map[string]string{
		"From":    from,
		"To":      strings.Join(to, ","),
		"Cc":      strings.Join(cc, ","),
		"Subject": subject,
	}
	msg := ""
	for k, v := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	const mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += mime
	msg += content
	return []byte(msg)
}
