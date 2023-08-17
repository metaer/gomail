package main

import (
	"crypto/tls"
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	gosmtp "github.com/emersion/go-smtp"

	"gomail/internal/config"
)

func main() {
	cfg, err := config.NewDevClientConfigFromYaml("dev_config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	c, err := gosmtp.Dial("localhost:10587")
	if err != nil {
		log.Fatal(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer c.Close()

	// #nosec - в отладочной команде линтер не должен ругаться на InsecureSkipVerify
	if err = c.StartTLS(&tls.Config{MinVersion: tls.VersionTLS12, InsecureSkipVerify: true}); err != nil {
		log.Fatal(err)
	}

	err = c.Auth(sasl.NewPlainClient("", cfg.Username, cfg.Password))
	if err != nil {
		log.Fatal(err)
	}

	from := "admin@" + cfg.FromDomain

	to := []string{cfg.DevRcptEmail}
	msg := strings.NewReader(
		"To: " + cfg.DevRcptEmail + "\r\n" +
			"From: " + from + "\r\n" +
			"Subject: тестовое письмо\r\n" +
			"\r\n" +
			"Тело письма\r\n")
	err = c.SendMail(from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
