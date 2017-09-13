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
		"order": "+networkDetected->'$.ip'",
	}
	fmt.Println(q.SqlStringWhere())
	fmt.Println(SqlStringOrder(q))
}
