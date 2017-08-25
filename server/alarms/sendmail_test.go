package alarms

import (
	"testing"
	"fmt"
)

func TestSendMail(t *testing.T) {
	to := []string{"xudi@k2data.com.cn", "xudi@ict.ac.cn"}
	subj := "This is the email subject"
    //body := "This is an example body.\nWith two lines."
	body2 := `
	<html>
	<h1>Hello</h1>
	</html>
	`
	err := SendMail(to, subj, body2)
	fmt.Println(err)
}

