package model

import (
	"testing"
	"fmt"
)

func TestAccess(t *testing.T) {
	host := Host{}
	host.Match = true
	fmt.Println(host.Match)
	fmt.Println(host.Json())
}

