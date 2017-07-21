package model

import (
	"time"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
)

type Host struct {
	DataCenter string `json:"datacenter"`
	Rack string `json:"rack"`
	Slot string `json:"slot"`
	Hostname string `json:"hostname"`
	Tags []string `json:"tags"`
	Owner string `json:"owner"`
	OsExpected osInfo `json:"osExpected"`
	OsDetected osInfo `json:"osDetected"`
	CpuExpected cpuInfo `json:"cpuExpected"`
	CpuDetected cpuInfo `json:"cpuDetected"`
	MemExpected memInfo `json:"memExpected"`
	MemDetected memInfo `json:"memDetected"`
	DiskExpected []diskInfo `json:"diskExpected"`
	DiskDetected []diskInfo `json:"diskDetected"`
	NetworkExpected networkInfo `json:"networkExpected"`
	NetworkDetected networkInfo `json:"networkDetected"`
	Registered bool `json:"registered"`
	Connected bool `json:"connected"`
	Match bool `json:"match"`
	Online bool `json:"online"`
	HealthStatus string `json:"healthStatus"`
	FirstSeenAt time.Time `json:"firstSeenAt"`
	LastSeenAt time.Time `json:"lastSeenAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type osInfo struct {
	OsType string `json:"type"`
	Arch string `json:"arch"`
	Hostname string `json:"hostname"`
}

type cpuInfo struct {
	Cpu int `json:"cpu"`
	Vcpu int `json:"vcpu"`
	Model string `json:"model"`
}

type memInfo struct {
	TotalMem uint64 `json:"totalMem"`
}

type diskInfo struct {
	Device string `json:"device"`
	Capacity uint64 `json:"capacity"`
}

type networkInfo struct {

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