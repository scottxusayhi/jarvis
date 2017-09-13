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
		"owner": "xudi,shenhongyuan",
		"order": "+networkDetected->'$.ip'",
	}
	fmt.Println(q.SqlStringWhere())
	fmt.Println(SqlStringOrder(q))
	pi := PageInfo(q)
	fmt.Println((&pi).SqlString())
}
