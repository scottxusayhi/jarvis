package options

import (
	"bytes"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/mitchellh/go-homedir"
	"path"
)

var (
	Master      string
	HBInterval  int
	Debug       bool
	AgentIdFile string
	UseMasterTime bool

	defaultAgentIdFile, _ = homedir.Expand("~/.jarvis/agent/id")
)

const (
	//defaultMaster = "localhost:2999"
	defaultHBInterval  = 30
	defaultDebug       = false
	defaultUseMasterTime = false
)

// TODO ENV -> CLI -> default
func LoadCli() {
	AgentIdFile = defaultAgentIdFile
	prepareIdFile(AgentIdFile)
	flag.StringVar(&Master, "master", "", "Master server address, e.g., 1.2.3.4:2999 (required)")
	flag.IntVar(&HBInterval, "heartbeat-interval", defaultHBInterval, "Heartbeat interval, in seconds.")
	flag.BoolVar(&Debug, "debug", defaultDebug, "Debug mode enabled. (default false)")
	flag.BoolVar(&UseMasterTime, "use-master-time", defaultUseMasterTime, "Use timestamp when message received rather than collected")
	flag.Parse()
	check()
}

func prepareIdFile(idFile string) {
	_, err := os.Stat(idFile)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(idFile), 0700)
		if err != nil {
			log.Error(err.Error())
			return
		}
		file, err := os.Create(idFile)
		if err != nil {
			log.Error(err.Error())
			return
		}
		defer file.Close()
	}
}

func check() {
	if Master == "" {
		log.Fatal("Missing option --master")
	}

	_, err := os.Stat(AgentIdFile)
	if os.IsNotExist(err) {
		log.Fatal("Can not read nor create agent id file " + AgentIdFile)
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

func ReadAgentIdFromFile() (string, error) {
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
	return string(bytes.Trim(id, "\x00")), nil
}
