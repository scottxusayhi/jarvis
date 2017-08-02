package options

import (
	"flag"
	log "github.com/sirupsen/logrus"
)

var (
	Master     string
	HBInterval int
	Debug bool
	AgentId         string
	AgentIdFile string
)

const (
	defaultMaster = "localhost:2999"
	defaultHBInterval = 30
	defaultDebug = false
	defaultAgentId= ""
	defaultAgentIdFile = "/var/run/jarvis.agent.id"
)

// ENV -> CLI -> default
func LoadCli() {
	flag.StringVar(&Master, "master", "", "Master server address, e.g., 1.2.3.4:2999 (required)")
	flag.IntVar(&HBInterval, "heartbeat-interval", 30, "Heartbeat interval, in seconds.")
	flag.BoolVar(&Debug, "debug", false, "Debug mode enabled. (default false)")
	flag.Parse()
	check()
}


func check() {
	if Master == "" {
		log.Fatal("Missing option --master")
	}
}

func WriteBack() {

}