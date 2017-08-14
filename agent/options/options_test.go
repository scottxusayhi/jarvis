package options

import (
	"fmt"
	"testing"
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
)

func TestGetAgentIdFromFile(t *testing.T) {
	AgentIdFile = defaultAgentIdFile
	fmt.Println(ReadAgentIdFromFile())
}

func TestWriteBackAgentIdFile(t *testing.T) {
	AgentIdFile = defaultAgentIdFile
	fmt.Println(WriteBackAgentIdFile("testid"))
}

func TestMisc(t *testing.T) {
	var s string
	fmt.Printf("\"%v\"", s)
}

func TestMkdir(t *testing.T) {
	idFile, err := homedir.Expand("~/.jarvis/agent/id")
	fmt.Println(idFile, err)

	fmt.Println(path.Dir(idFile))

	fmt.Println(os.MkdirAll(path.Dir(idFile), 0700))

}

