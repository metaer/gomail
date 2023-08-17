package main

import (
	"log"
	"net/smtp"
)

const (
	sendTo             = "example@example.com" // Заменить перед выполнением на свой email
	smtpServerUsername = "username"            // Заменить перед выполнением на реальный username для аутентификации на smtp сервере
	smtpServerPassword = "password"            // Заменить перед выполнением на реальный пароль для аутентификации на smtp сервере
)

func main() {
	/*
		Отправителя тоже можно заменить на своего, но в этом случае надо добавить к своему домену txt-запись SPF вида:
		v=spf1 include:spf.mailer-demo.ru
		Иначе увеличивается риск попадания в спам
	*/
	from := "admin@mailer-demo.ru"

	message := []byte("Subject: Тестовая тема\r\n" +
		"To: " + sendTo + "\r\n" +
		"From: " + from + "\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
		"\r\n" +
		"Тело сообщения")

	smtpHost := "smtp.mailer-demo.ru"
	smtpPort := "587"

	auth := smtp.PlainAuth("", smtpServerUsername, smtpServerPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{sendTo}, message)
	if err != nil {
		log.Fatal(err)
	}
}
