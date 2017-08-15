package options

import (
	"fmt"
	"testing"
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
	"math"
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

func TestDivide(t *testing.T) {
	fmt.Println(float32(2)/3)
	fmt.Println(math.Ceil(float64(2)/3))

	fmt.Println(5%2)

	x, y := 6, 2
	if x%y >0 {
		fmt.Println(x/y+1)
	} else {
		fmt.Println(x/y)
	}
}


