package mysql

import (
	"testing"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func TestJarvisMysqlBackend_SearchHost(t *testing.T) {
	backend, _ := GetBackend()
	query := make(map[string]string)
	hosts, _ := backend.SearchHost(query)
	fmt.Println(hosts)
}

func TestJarvisMysqlBackend_GetOneHost(t *testing.T) {
	backend, _ := GetBackend()
	host, _ := backend.GetOneHost("goldwind", "01", "010203", "kmx-1")
	log.Info(host.JsonString())
}


