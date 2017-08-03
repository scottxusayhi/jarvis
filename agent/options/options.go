package options

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	Master     string
	HBInterval int
	Debug bool
	AgentIdFile string
)

const (
	//defaultMaster = "localhost:2999"
	defaultHBInterval = 30
	defaultDebug = false
	defaultAgentIdFile = "./jarvis.agent.id"
)

// TODO ENV -> CLI -> default
func LoadCli() {
	flag.StringVar(&Master, "master", "", "Master server address, e.g., 1.2.3.4:2999 (required)")
	flag.IntVar(&HBInterval, "heartbeat-interval", defaultHBInterval, "Heartbeat interval, in seconds.")
	flag.BoolVar(&Debug, "debug", defaultDebug, "Debug mode enabled. (default false)")
	flag.Parse()
	check()
}


func check() {
	if Master == "" {
		log.Fatal("Missing option --master")
	}
}

func WriteBackAgentIdFile(id string) error {
	var idFile *os.File
	var err error
	defer idFile.Close()

	os.Remove(AgentIdFile)

	_, err = os.Stat(AgentIdFile)
	if os.IsNotExist(err) {
		idFile, err = os.Create(AgentIdFile)
		if err != nil {
			return err
		}
	} else {
		idFile, err = os.OpenFile(AgentIdFile, os.O_RDWR, 0755)
		if err != nil {
			return err
		}
	}
	idFile.WriteString(id)
	return nil
}

func GetAgentIdFromFile() (string, error) {
	var idFile *os.File
	var err error
	defer idFile.Close()

	idFile, err = os.Open(AgentIdFile)
	if err != nil {
		return "", err
	}
	id := make([]byte, 128)
	_, err = idFile.Read(id)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(id), "\n"), nil
}
