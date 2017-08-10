package utils

import (
	"testing"
	"fmt"
)

func TestOsExec(t *testing.T) {
	fmt.Println(OsExec("/bin/ls", "-a"))
	fmt.Println(OsExecLineOut("/bin/ls", "-a"))
	// errors
	fmt.Println(OsExec("/bin/unknown", "-a"))
	fmt.Println(OsExecLineOut("/bin/unknown", "-a"))
}

