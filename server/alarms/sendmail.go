package alarms

import (
	"net/smtp"
	"fmt"
	log "github.com/sirupsen/logrus"
	"crypto/tls"
	"strings"
)

var (
	smtp_host = "smtp.exmail.qq.com"
	smtp_port = 465
	smtp_from = "ops@k2data.com.cn"
	smtp_password = "K2admin1217"
)

func SendMail(to []string, subj string, body string) (err error) {
	auth := smtp.PlainAuth("", smtp_from, smtp_password, smtp_host)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName: smtp_host,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%v:%v", smtp_host, smtp_port), tlsconfig)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, smtp_host)
	if err != nil {
		log.Error(err.Error())
		return
	}

	err = c.Auth(auth)
	if err != nil {
		log.Error(err.Error())
		return
	}

	err = c.Mail(smtp_from)
	if err != nil {
		log.Error(err.Error())
		return
	}

	for _, toOne := range to {
		err = c.Rcpt(toOne)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}

	wc, err := c.Data()
	if err != nil {
		log.Error(err.Error())
		return
	}

    // Setup headers
    headers := make(map[string]string)
    headers["From"] = smtp_from
    headers["To"] = strings.Join(to, ",")
    headers["Subject"] = subj

    // Setup message
    message := ""
    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

	_, err = wc.Write([]byte(message))
	if err != nil {
		log.Error(err.Error())
		return
	}

	err = wc.Close()
	if err != nil {
		log.Error(err.Error())
		return
	}

	err = c.Quit()
	if err != nil {
		log.Error(err.Error())
		return
	}
	return

}

