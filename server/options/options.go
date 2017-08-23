package options

import (
	"flag"
	log "github.com/sirupsen/logrus"
)


//"root:passw0rd@tcp(localhost:3306)/jarvis?parseTime=true"
var (
	MysqlAddr string
	Debug bool
)

// TODO ENV -> CLI -> default
func LoadCli() {
	flag.StringVar(&MysqlAddr, "mysql-addr", "localhost:3306", "mysql address")
	flag.BoolVar(&Debug, "debug", false, "Debug mode enabled. (default false)")
	flag.Parse()
	check()
}


func check() {
	if false {
		log.Fatal("invalid parameters")
	}
}

