package protocol

import (
	"encoding/json"
	"git.oschina.net/k2ops/jarvis/utils"
)


// common json func
func serialize(v interface{}) []byte {
	ret, err := json.Marshal(v)
	if err != nil {
		return []byte("{}")
	}
	return ret
}

// abstract class
type JarvisMessage struct {
	MessageType string `json:"type"`
}
func (m *JarvisMessage) Serialize() []byte {
	return serialize(m)
}
func (m *JarvisMessage) ToJsonString() string {
	return string(m.Serialize())
}


// subclassing
// welcome message
type welcomeMessage struct {
	JarvisMessage
	ClientAddr string `json:"clientAddr"`
	ServerAddr string `json:"serverAddr"`
}
func (m *welcomeMessage) Serialize() []byte {
	return serialize(m)
}
func (m *welcomeMessage) ToJsonString() string {
	return string(m.Serialize())
}
func NewWelcomeMessage(clientAddr string, serverAddr string) *welcomeMessage {
	m := welcomeMessage{}
	m.MessageType = "welcome"
	m.ClientAddr = clientAddr
	m.ServerAddr = serverAddr
	return &m
}
func NewEmptyWelcomeMessage() *welcomeMessage {
	m := welcomeMessage{}
	m.MessageType = "welcome"
	return &m
}

// client register message
type PhysicalDiskInfo struct {
	Device string `json:"device"`
	Total  uint64 `json:"total"`
	Used uint64 `json:"used"`
}
type registerMessage struct {
	JarvisMessage
	UpdatedAt string `json:"updatedAt"`
	OSType string `json:"osType"`
	Arch string `json:"arch"`
	Hostname string `json:"hostname"`
	CPUNum int `json:"cpuNum"`
	MemTotal int `json:"memTotal"`
	UpTime string `json:"upTime"`
	Disks []PhysicalDiskInfo `json:"disks"`
}
func (m *registerMessage) Serialize() []byte {
	return serialize(m)
}
func (m *registerMessage) ToJsonString() string {
	return string(m.Serialize())
}
func NewEmptyRegisterMessage() *registerMessage {
	m := registerMessage{}
	m.MessageType = "register"
	return &m
}

// client heartbeat message
type heartbeatMessage struct {
	JarvisMessage
	UpdatedAt string `json:"updatedAt"`
}
func (m *heartbeatMessage) Serialize() []byte {
	return serialize(m)
}
func (m *heartbeatMessage) ToJsonString() string {
	return string(m.Serialize())
}
func NewHeartbeatMessage() *heartbeatMessage {
	m := heartbeatMessage{}
	m.MessageType = "heartbeat"
	_, m.UpdatedAt = timeutils.ISO8601Now()
	return &m
}

// resource usage message
type resourceUsageMessage struct {
	JarvisMessage
	UpdatedAt string `json:"updatedAt"`
	CPUPercent float32 `json:"cpuPercent"`
	MemUsed uint64 `json:"memUsed"`
	Disks []PhysicalDiskInfo `json:"disks"`
	Network uint `json:"network"`
}
func (m *resourceUsageMessage) Serialize() []byte {
	return serialize(m)
}
func (m *resourceUsageMessage) ToJsonString() string {
	return string(m.Serialize())
}
func NewEmptyResourceUsageMessage () *resourceUsageMessage {
	m := resourceUsageMessage{}
	m.MessageType = "resource-usage"
	return &m
}

// command message
// command-response message
// service status message