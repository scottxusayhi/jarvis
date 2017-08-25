package mysql

import (
	"testing"
	"fmt"
)

func TestJarvisMysqlBackend_GetActiveAlarms(t *testing.T) {
	backend, _ := GetBackend()
	fmt.Println(backend.GetActiveAlarms())


}

