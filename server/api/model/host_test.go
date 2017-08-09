package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"testing"
)

func TestHost_Json(t *testing.T) {
	host := Host{}
	host.Matched = true
	fmt.Println(host.Matched)
	fmt.Println(host.JsonString())
}

func TestHost_Parse(t *testing.T) {
	host, err := ParseHost(strings.NewReader("sds"))
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println(host.JsonString())
}
