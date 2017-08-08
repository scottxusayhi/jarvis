package backend

import (
	"fmt"
	"testing"
)

func TestNewQuery(t *testing.T) {
	q := Query{
		"datacenter": "goldwind",
		"rack":       "01",
		"slot":       "010203",
	}
	fmt.Println(q.SqlString())
}
