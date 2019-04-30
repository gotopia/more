package mail

import "net/smtp"

func Send() error {
	if err := smtp.SendMail(smtpAddr, nil, from, to, []byte(msg)); err != nil {
		zap.L().Error(fmt.Sprintf("%v", err))
		return err
	}
}
