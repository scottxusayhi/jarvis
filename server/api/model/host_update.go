package model

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"strings"
	"fmt"
	"io"
	"bytes"
	"time"
	"github.com/scottxusayhi/jarvis/utils"
	"sort"
)

type HostUpdate struct {
	DataCenter      string      `json:"datacenter,omitempty"`
	Rack            string      `json:"rack,omitempty"`
	Slot            string      `json:"slot,omitempty"`
	Tags            hostTags    `json:"tags,omitempty"`
	Owner           string      `json:"owner,omitempty"`
	OsExpected      OsInfo      `json:"osExpected,omitempty"`
	CpuExpected     CpuInfo     `json:"cpuExpected,omitempty"`
	MemExpected     MemInfo     `json:"memExpected,omitempty"`
	DiskExpected    HostDisks   `json:"diskExpected,omitempty"`
	NetworkExpected NetworkInfo `json:"networkExpected,omitempty"`
	Comments string `json:"comments,omitempty"`
}

func (hu *HostUpdate) JsonBytes() []byte {
	bytes, err := json.Marshal(hu)
	if err != nil {
		log.Error(err.Error())
		return []byte("{}")
	}
	return bytes
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}

//
func ParseUpdatableFields(r io.Reader) (result map[string]interface{}, err error) {
	// first parse to HostUpdate struct for value validation
	var raw bytes.Buffer
	_, err = raw.ReadFrom(r)
	if err != nil {
		return
	}
	hostUpdate := HostUpdate{}
	if err != nil {
		log.Error("read error")
		return
	}
	err = json.Unmarshal(raw.Bytes(), &hostUpdate)
	if err != nil {
		log.Error("bytes->HostUpdate error")
		return
	}
	// struct to bytes
	bytes, err := json.Marshal(&hostUpdate)
	if err != nil {
		log.Error("HostUpdate->bytes error")
		return
	}
	// and bytes to map
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		log.Error("bytes->map error")
		return
	}

	// raw map
	rawMap := make(map[string]interface{})
	json.Unmarshal(raw.Bytes(), &rawMap)
	inputKeys := make([]string, 1)
	for k, _ := range rawMap {
		inputKeys = append(inputKeys, k)
	}
	//

	for k, _ := range result {
		if !contains(inputKeys, k) {
			delete(result, k)
		}
	}
	log.WithFields(log.Fields{
		"value": result,
	}).Info("update host payload")
	return

}

func UpdateSqlString(m map[string]interface{}) (s string) {
	// guarantee key order for each iterations to make correct sql update
	keys := make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fields := make([]string, 0)
	for _, k := range keys {
		fields = append(fields, fmt.Sprintf("%v=?", k))
	}
	fields = append(fields, "updatedAt=?")
	return strings.Join(fields, ", ")
}

func UpdateValues(m map[string]interface{}) (result []interface{}) {
	// guarantee key order for each iterations to make correct sql update
	keys := make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)


	for _, k := range keys {
		result = append(result, utils.SafeMarshalJson(m[k]))
	}
	// updatedAt
	result = append(result, time.Now())
	return result
}

