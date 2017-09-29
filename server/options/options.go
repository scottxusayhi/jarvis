package options

import (
	"flag"
	log "github.com/sirupsen/logrus"
)


//"root:passw0rd@tcp(localhost:3306)/jarvis?parseTime=true"
var (
	MysqlAddr string = "localhost:3306"
	Debug bool = false
	UseMasterTime bool = false
)

// TODO ENV -> CLI -> default
func LoadCli() {
	flag.StringVar(&MysqlAddr, "mysql-addr", "localhost:3306", "mysql address")
	flag.BoolVar(&Debug, "debug", false, "Debug mode enabled. (default false)")
	flag.BoolVar(&UseMasterTime, "use-master-time", false, "use master time as heartbeat time")
	flag.Parse()
	check()
}


func check() {
	if false {
		log.Fatal("invalid parameters")
	}
}

