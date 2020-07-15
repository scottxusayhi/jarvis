package mysql

import (
	"testing"
	"fmt"
	"github.com/scottxusayhi/jarvis/server/backend"
)

func TestJarvisMysqlBackend_ListDatacenters(t *testing.T) {
	backend, _ := GetBackend()
	fmt.Println(backend.ListDatacenters())
}

func TestJarvisMysqlBackend_SearchHost2(t *testing.T) {
	query := backend.Query{
		"registered": "1",
		"matched": "1",
	}

	backend, _ := GetBackend()
	fmt.Println(backend.SearchHost(query))
}


