package model

import (
	"time"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"fmt"
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
func (host *Host) Json () string {
	bytes, err := json.Marshal(host)
	if err != nil {
		log.Error(err.Error())
		return "{}"
	}
	return string(bytes)
}

type osInfo struct {
	osType string `json:"type"`
	arch string `json:"arch"`
	hostname string `json:"hostname"`
}

type cpuInfo struct {
	cpu int `json:"cpu"`
	vcpu int `json:"vcpu"`
	model string `json:"model"`
}

type memInfo struct {
	totalMem uint64 `json:"totalMem"`
}

type diskInfo struct {
	device string `json:"device"`
	capacity uint64 `json:"capacity"`
}

type networkInfo struct {

}