package backend

import (
	"testing"
	"fmt"
)

func TestNewQuery(t *testing.T) {
	q := Query {
		"datacenter": "goldwind",
		"rack": "01",
		"slot": "010203",
	}
	fmt.Println(q.String())
}

