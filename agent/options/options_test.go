package options

import (
	"testing"
	"fmt"
)

func TestGetAgentIdFromFile(t *testing.T) {
	AgentIdFile=defaultAgentIdFile
	fmt.Println(GetAgentIdFromFile())
}

func TestWriteBackAgentIdFile(t *testing.T) {
	AgentIdFile=defaultAgentIdFile
	fmt.Println(WriteBackAgentIdFile("testid"))
}

func TestMisc(t *testing.T) {
	var s string
	fmt.Printf("\"%v\"", s)
}


