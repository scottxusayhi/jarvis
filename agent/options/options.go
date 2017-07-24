package options

import (
	"os"
	"flag"
	log "github.com/sirupsen/logrus"
)

var (
	std = LoadCliFlags()
)

type JarvisClientOptions struct {
	Master     string
	Namespace  string
	HBInterval int
	Debug bool
	Id         string
}
// ENV -> CLI -> default
func LoadCliFlags() *JarvisClientOptions {
	o := JarvisClientOptions{}
	flag.StringVar(&o.Master, "master", "", "Master server address, e.g., 1.2.3.4:2999 (required)")
	flag.StringVar(&o.Namespace, "namespace", "default", "Agent namespace, will be used to generated agent ID.")
	flag.IntVar(&o.HBInterval, "heartbeat-interval", 30, "Heartbeat interval, in seconds.")
	flag.BoolVar(&o.Debug, "debug", false, "Debug mode enabled. (default false)")
	flag.Parse()
	hn, err := os.Hostname()
	if err!=nil {
		hn = "unknown"
	}
	o.Id = o.Namespace+"/"+hn
	return &o
}


func Check() {
	if std.Master == "" {
		log.Fatal("Missing option --master")
	}
}

func Flags() *JarvisClientOptions {
	return std
}