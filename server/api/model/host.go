package model

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"time"
	//"github.com/go-sql-driver/mysql"
	"git.oschina.net/k2ops/jarvis/protocol"
)

type Host struct {
	SystemId        int         `json:"systemId"`
	DataCenter      string      `json:"datacenter"`
	Rack            string      `json:"rack"`
	Slot            string      `json:"slot"`
	Tags            hostTags    `json:"tags"`
	Owner           string      `json:"owner"`
	OsExpected      OsInfo      `json:"osExpected"`
	OsDetected      protocol.OsInfo      `json:"osDetected"`
	CpuExpected     CpuInfo     `json:"cpuExpected"`
	CpuDetected     protocol.CpuInfo     `json:"cpuDetected"`
	MemExpected     MemInfo     `json:"memExpected"`
	MemDetected     protocol.MemInfo     `json:"memDetected"`
	DiskExpected    HostDisks   `json:"diskExpected"`
	DiskDetected    protocol.HostDisks   `json:"diskDetected"`
	NetworkExpected NetworkInfo `json:"networkExpected"`
	NetworkDetected protocol.NetworkInfo `json:"networkDetected"`
	Registered      bool        `json:"registered"`
	Connected       bool        `json:"connected"`
	Matched         bool        `json:"matched"`
	Online          bool        `json:"online"`
	HealthStatus    string      `json:"healthStatus"`
	FirstSeenAt     time.Time   `json:"firstSeenAt"`
	LastSeenAt      time.Time   `json:"lastSeenAt"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
	Comments string `json:"comments"`
}

type hostTags []string

func (ht *hostTags) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("hostTags must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, ht)
}

type HostDisks []DiskInfo

func (hd *HostDisks) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("hostDisks must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, hd)
}

type OsInfo struct {
	OsType   string `json:"type"`
	Arch     string `json:"arch"`
	Hostname string `json:"hostname"`
}

func (oi *OsInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("osInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, oi)
}

type CpuInfo struct {
	Socket   int    `json:"socket"`
	Vcpu  int    `json:"vcpu"`
	Model string `json:"model"`
}

func (ci *CpuInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("cpuInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, ci)
}

type MemInfo struct {
	Total uint64 `json:"total"`
}

func (mi *MemInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("memInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, mi)
}

type DiskInfo struct {
	Device   string `json:"device"`
	Capacity uint64 `json:"capacity"`
}

func (di *DiskInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("diskInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, di)
}

type NetworkInfo struct {
	Ip string `json:"ip"`
}

func (ni *NetworkInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("networkInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, ni)
}

func (host *Host) JsonBytes() []byte {
	bytes, err := json.Marshal(host)
	if err != nil {
		log.Error(err.Error())
		return []byte("{}")
	}
	return bytes
}

func (host *Host) JsonString() string {
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
