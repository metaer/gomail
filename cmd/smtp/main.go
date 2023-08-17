package main

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/emersion/go-smtp"

	"gomail/internal/config"
	"gomail/internal/mail"
)

func main() {
	cfg, err := config.NewServerConfigFromYaml(config.ConfigFilePath)
	if err != nil {
		log.Fatal(err)
	}
	tlsCfg, err := getTlsConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s := smtp.NewServer(mail.NewBackend())
	s.Addr = ":" + cfg.Port
	s.Domain = cfg.SmtpServerDomain
	s.ReadTimeout = 15 * time.Second
	s.WriteTimeout = 15 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 1
	s.AllowInsecureAuth = false
	s.TLSConfig = tlsCfg
	s.EnableREQUIRETLS = true

	log.Println("Сервер стартовал на", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getTlsConfig(cfg *config.ServerConfig) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(cfg.TlsCert, cfg.TlsKey)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}
	return tlsConfig, nil
}
