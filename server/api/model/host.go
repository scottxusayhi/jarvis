package model

import (
	"time"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"fmt"
	//"github.com/go-sql-driver/mysql"
)

type Host struct {
	SystemId int `json:"systemId"`
	DataCenter string `json:"datacenter"`
	Rack string `json:"rack"`
	Slot string `json:"slot"`
	Tags hostTags `json:"tags"`
	Owner string `json:"owner"`
	OsExpected osInfo `json:"osExpected"`
	OsDetected osInfo `json:"osDetected"`
	CpuExpected cpuInfo `json:"cpuExpected"`
	CpuDetected cpuInfo `json:"cpuDetected"`
	MemExpected memInfo `json:"memExpected"`
	MemDetected memInfo `json:"memDetected"`
	DiskExpected hostDisks `json:"diskExpected"`
	DiskDetected hostDisks `json:"diskDetected"`
	NetworkExpected networkInfo `json:"networkExpected"`
	NetworkDetected networkInfo `json:"networkDetected"`
	Registered bool `json:"registered"`
	Connected bool `json:"connected"`
	Matched bool `json:"matched"`
	Online bool `json:"online"`
	HealthStatus string `json:"healthStatus"`
	FirstSeenAt time.Time `json:"firstSeenAt"`
	LastSeenAt time.Time `json:"lastSeenAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type hostTags []string
func (ht *hostTags) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("hostTags must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, ht)
}

type hostDisks []diskInfo
func (hd *hostDisks) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("hostDisks must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, hd)
}

type osInfo struct {
	OsType string `json:"type"`
	Arch string `json:"arch"`
	Hostname string `json:"hostname"`
}
func (oi *osInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("osInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, oi)
}

type cpuInfo struct {
	Cpu int `json:"cpu"`
	Vcpu int `json:"vcpu"`
	Model string `json:"model"`
}
func (ci *cpuInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("cpuInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, ci)
}

type memInfo struct {
	TotalMem uint64 `json:"totalMem"`
}

func (mi *memInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("memInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, mi)
}

type diskInfo struct {
	Device string `json:"device"`
	Capacity uint64 `json:"capacity"`
}

func (di *diskInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("diskInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, di)
}

type networkInfo struct {

}
func (ni *networkInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("networkInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, ni)
}

func (host *Host) JsonBytes () []byte {
	bytes, err := json.Marshal(host)
	if err != nil {
		log.Error(err.Error())
		return []byte("{}")
	}
	return bytes
}

func (host *Host) JsonString () string {
	return string(host.JsonBytes())
}
func ParseHost(r io.Reader) (Host, error) {
	decoder := json.NewDecoder(r)
	newHost := Host{}
	err := decoder.Decode(&newHost)
	if err != nil {
		log.Error(err.Error())
		return newHost, err
	}
	return newHost, nil
}