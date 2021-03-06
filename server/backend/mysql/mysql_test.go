package mysql

import (
	"fmt"
	"github.com/scottxusayhi/jarvis/server/api/model"
	"github.com/scottxusayhi/jarvis/server/backend"
	"github.com/scottxusayhi/jarvis/utils"
	log "github.com/sirupsen/logrus"
	"strings"
	"testing"
)

func TestJarvisMysqlBackend_CreateHost(t *testing.T) {
	var err error
	// generate test data
	jsonStr := "{\"datacenter\":\"k2\",\"rack\":\"01\",\"slot\":\"010203\",\"tags\":[\"tag1\",\"tag2\"],\"owner\":\"cluster100\",\"osExpected\":{\"type\":\"Linux\",\"arch\":\"amd64\"},\"cpuExpected\":{\"cpu\":2,\"vcpu\":12,\"model\":\"\"},\"memExpected\":{\"totalMem\":128000000000},\"diskExpected\":[{\"device\":\"/dev/disk0\",\"capacity\":2000000000000}],\"networkExpected\":{}}"
	host, err := model.ParseHost(strings.NewReader(jsonStr))
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println(host.JsonString())

	// create
	backend, _ := GetBackend()
	err = backend.CreateHost(host)
	if err != nil {
		log.Error(err.Error())
	}
}

func TestJarvisMysqlBackend_GetOneHost(t *testing.T) {
	backend, _ := GetBackend()
	host, _ := backend.GetOneHost("goldwind", "01", "010203")
	log.Info(host.JsonString())
}

func TestJarvisMysqlBackend_SearchHost(t *testing.T) {
	query := backend.Query{
		"datacenter": "goldwind",
	}
	fmt.Println(query.SqlStringWhere())
	backend, err := GetBackend()
	if err != nil {
		log.Error(err.Error())
	}
	hosts, _, err := backend.SearchHost(query)
	if err != nil {
		log.Error(err.Error())
	}
	for _, host := range hosts {
		fmt.Println(host.JsonString())
	}
}

func TestJarvisMysqlBackend_PreserveId(t *testing.T) {
	backend, _ := GetBackend()
	fmt.Println(backend.PreserveId())
	fmt.Println(backend.PreserveId())
}

func TestJarvisMysqlBackend_UpdateConnectionInfo(t *testing.T) {
	backend, _ := GetBackend()
	newId, _ := backend.PreserveId()
	fmt.Println(backend.UpdateConnectionInfo(newId))
}

func TestMisc(t *testing.T) {
	fmt.Println(utils.UnknownDataCenter())
	fmt.Println(utils.UnknownRack())
	fmt.Println(utils.UnknownSlot())
}

func TestJarvisMysqlBackend_MarkOffline(t *testing.T) {
	backend, _ := GetBackend()
	fmt.Println(backend.MarkOffline(38))
}
