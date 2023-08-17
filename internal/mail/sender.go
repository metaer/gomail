package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
	"time"
)

func sendDirectlyToRcptMailServer(data []byte, from, rcptTo string) error {
	mxs, err := getMXRecords(rcptTo)
	if err != nil {
		return err
	}
	if len(mxs) == 0 {
		return fmt.Errorf("mx-записи не найдены")
	}

	timeout := 15 * time.Second
	var host string
	var conn net.Conn

	for _, mx := range mxs {
		log.Println("Подключение к mx-серверу:", mx.Host)
		conn, err = net.DialTimeout("tcp", fmt.Sprintf("%s:25", mx.Host), timeout)
		if err == nil { // Если всё в порядке - запоминаем хост и выходим из цикла
			host = mx.Host
			break
		}
	}
	if err != nil { //Если не удалось соединиться ни с одним из mx-серверов
		return fmt.Errorf("не удалось соединиться ни с одним из mx-серверов")
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	err = c.StartTLS(&tls.Config{ServerName: host, MinVersion: tls.VersionTLS12})
	if err != nil {
		return err
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	if err = c.Rcpt(rcptTo); err != nil {
		return err
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	err = c.Quit()
	if err != nil {
		return err
	}

	return nil
}

func getMXRecords(to string) ([]*net.MX, error) {
	var e *mail.Address
	e, err := mail.ParseAddress(to)
	if err != nil {
		return nil, err
	}

	domain := strings.Split(e.Address, "@")[1]

	var mxs []*net.MX
	mxs, err = net.LookupMX(domain)

	if err != nil {
		return nil, err
	}

	return mxs, nil
}
