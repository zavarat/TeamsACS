package gmail

import (
	"crypto/tls"
	"fmt"
	"mime"
	"path"

	"gopkg.in/gomail.v2"

	"github.com/ca17/teamsacs/common"
)

type MailSender struct {
	Server   string
	Port     int
	Tls      bool
	Usernam  string
	Alias    string
	Password string
	Mailtos  []string
}

func (s *MailSender) SendMail(mailTo []string, subject string, body string, files []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(s.Usernam, s.Alias))
	if mailTo == nil || len(mailTo) == 0 {
		if s.Mailtos == nil || len(s.Mailtos) ==0 {
			return fmt.Errorf("Mail receiver not configured")
		}
		m.SetHeader("To", s.Mailtos...)
	}

	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if files != nil && len(files) > 0 {
		m := gomail.NewMessage(
			gomail.SetEncoding(gomail.Base64),
		)
		for _, filename := range files {
			if !common.FileExists(filename) {
				return fmt.Errorf("file %s not exists", filename)
			}
			name := path.Base(filename)
			m.Attach(filename,
				gomail.Rename(name),
				gomail.SetHeader(map[string][]string{
					"Content-Disposition": []string{
						fmt.Sprintf(`attachment; filename="%s"`, mime.BEncoding.Encode("UTF-8", name)),
					},
				}),
			)
		}
	}

	d := gomail.NewDialer(s.Server, s.Port, s.Usernam, s.Password)
	if s.Tls {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return d.DialAndSend(m)
}
