package mail

import (
	"io"
	"log"

	"github.com/emersion/go-smtp"

	"gomail/internal/config"
)

type Session struct {
	auth   bool
	from   string
	rcptTo string
	cfg    *config.ServerConfig
}

func (s *Session) AuthPlain(username, password string) error {
	var err error
	if s.cfg == nil {
		s.cfg, err = config.NewServerConfigFromYaml(config.ConfigFilePath)
		if err != nil {
			return err
		}
	}

	users := s.cfg.Users

	for _, user := range users {
		if username == user.Username && password == user.Password {
			s.auth = true
			return nil
		}
	}

	return smtp.ErrAuthFailed
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	if !s.auth {
		return smtp.ErrAuthRequired
	}
	log.Println("Mail from:", from)
	s.from = from
	return nil
}

func (s *Session) Rcpt(to string) error {
	if !s.auth {
		return smtp.ErrAuthRequired
	}
	log.Println("Rcpt to:", to)
	s.rcptTo = to
	return nil
}

func (s *Session) Data(r io.Reader) error {
	var err error
	if s.cfg == nil {
		s.cfg, err = config.NewServerConfigFromYaml(config.ConfigFilePath)
		if err != nil {
			return err
		}
	}
	if !s.auth {
		return smtp.ErrAuthRequired
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if s.cfg.AddDkimSignature {
		data, err = sign(data, s.cfg.DkimPrivateKeyFile, s.cfg.DkimSelector, s.cfg.FromDomain)
		if err != nil {
			return err
		}
	}
	log.Println("Data:", string(data))
	log.Println("Успешно принято к отправке")
	err = sendDirectlyToRcptMailServer(data, s.from, s.rcptTo)
	if err != nil {
		return err
	}
	log.Println("Успешно отправлено")
	return nil
}

func (s *Session) Reset() {
	s.auth = false
	s.from = ""
	s.rcptTo = ""
}

func (s *Session) Logout() error {
	s.auth = false
	return nil
}
