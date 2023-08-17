package mail

import "github.com/emersion/go-smtp"

type Backend struct{}

func NewBackend() smtp.Backend {
	return &Backend{}
}

func (bkd *Backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}
