package model

import (
	"testing"
	"encoding/json"
	"fmt"
	"bytes"
)

func TestParseUpdatableFields(t *testing.T) {
	input := map[string]interface{} {
		"foo": "bar",
		"datacenter": "newdatacenter",
		"owner": "owner",
		"rack": "rack",
		"slot": "slot",
		"tags": []string{"tag1", "tag2"},
	}
	inputBytes, err := json.Marshal(input)

	hu, err := ParseUpdatableFields(bytes.NewReader(inputBytes))
	fmt.Println(hu, err)

	fmt.Println(UpdateSqlString(hu))
	fmt.Println(UpdateValues(hu)...)

}

