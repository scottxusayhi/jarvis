package mysql

import (
	"testing"
	"fmt"
)

func TestJarvisMysqlBackend_GetEmailAlarmRecipients(t *testing.T) {
	backend, _ := GetBackend()
	fmt.Println(backend.GetEmailAlarmRecipients())
}

