build:
	rm -rf ./bin/gomail && GOOS=linux GOARCH=amd64 go build -o ./bin/gomail ./cmd/smtp
unlink:
	ssh root@smtp.mailer-demo.ru "unlink /usr/local/bin/gomail"
upload:
	scp ./bin/gomail root@smtp.mailer-demo.ru:/usr/local/bin/gomail
restart:
	ssh root@smtp.mailer-demo.ru "service gomail restart"

update: build unlink upload restart
