package protocol

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	MSG_HELLO             = "hello"
	MSG_WELCOME           = "welcome"
	MSG_HOST_CONFIG       = "host-config"
	MSG_HEARTBEAT         = "heartbeat"
	MSG_AGENT_ID_REQUEST  = "agent-id-request"
	MSG_AGENT_ID_RESPONSE = "agent-id-response"
	Footer                = '\r'
)

// common json func
func serialize(v interface{}) []byte {
	ret, err := json.Marshal(v)
	if err != nil {
		return append([]byte("{}"), Footer)
	}
	return append(ret, Footer)
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

// hello
type helloMessage struct {
	JarvisMessage
	ClientAddr string `json:"clientAddr"`
	ServerAddr string `json:"serverAddr"`
}

func (m *helloMessage) Serialize() []byte {
	return serialize(m)
}
func (m *helloMessage) ToJsonString() string {
	return string(m.Serialize())
}
func NewHelloMessage() *helloMessage {
	m := helloMessage{}
	m.MessageType = MSG_HELLO
	return &m
}

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

// detected host info
type OsInfo struct {
	OsType   string `json:"type"`
	Arch     string `json:"arch"`
	// (linux) distribution
	Dist string `json:"dist"`
	Version string `json:"version"`
	Hostname string `json:"hostname"`
	Uptime   uint64 `json:"uptime"`
}

func (oi *OsInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("osInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, oi)
}

type CpuInfo struct {
	Socket int    `json:"socket"`
	Vcpu    int `json:"vcpu"`
	Model   string `json:"model"`
}

func (ci *CpuInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("cpuInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, ci)
}

type MemInfo struct {
	Total     uint64 `json:"total"`
	Available uint64 `json:"available"`
	Used      uint64 `json:"used"`
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
	Model string `json:"model"`
	Capacity uint64 `json:"capacity"`
	Used     uint64 `json:"used"`
}

func (di *DiskInfo) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("diskInfo must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, di)
}

type HostDisks []DiskInfo
func (hd *HostDisks) Scan(src interface{}) error {
	byteValue, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("hostDisks must be a []byte, got %T instead", src)
	}
	return json.Unmarshal(byteValue, hd)
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

type HostConfigMessage struct {
	JarvisMessage
	AgentId string `json:"agentId"`
	UpdatedAt    time.Time `json:"updatedAt"`
	OsDetected   OsInfo    `json:"osDetected"`
	CpuDetected  CpuInfo   `json:"cpuDetected"`
	MemDetected  MemInfo   `json:"memDetected"`
	DiskDetected HostDisks `json:"diskDetected"`
	NetworkDetected NetworkInfo `json:"networkDetected"`
}

func (m *HostConfigMessage) Serialize() []byte {
	return serialize(m)
}
func (m *HostConfigMessage) ToJsonString() string {
	return string(m.Serialize())
}
func NewEmptyHostConfigMessage(aid string) *HostConfigMessage {
	m := HostConfigMessage{}
	m.MessageType = MSG_HOST_CONFIG
	m.AgentId = aid
	m.UpdatedAt = time.Now()
	return &m
}

// client heartbeat message
type HeartbeatMessage struct {
	JarvisMessage
	AgentId   string    `json:"agentId"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m *HeartbeatMessage) Serialize() []byte {
	return serialize(m)
}
func (m *HeartbeatMessage) ToJsonString() string {
	return string(m.Serialize())
}
func NewHeartbeatMessage(agentId string) *HeartbeatMessage {
	m := HeartbeatMessage{}
	m.MessageType = MSG_HEARTBEAT
	m.AgentId = agentId
	m.UpdatedAt = time.Now()
	return &m
}

// command message
// command-response message
// service status message

// agent id request
type agentIdRequest struct {
	JarvisMessage
}

func (m *agentIdRequest) Serialize() []byte {
	return serialize(m)
}
func (m *agentIdRequest) ToJsonString() string {
	return string(m.Serialize())
}
func NewAgentIdRequest() *agentIdRequest {
	m := agentIdRequest{}
	m.MessageType = MSG_AGENT_ID_REQUEST
	return &m
}

// agent id response
type AgentIdResponse struct {
	JarvisMessage
	AgentId string
}

func (m *AgentIdResponse) Serialize() []byte {
	return serialize(m)
}
func (m *AgentIdResponse) ToJsonString() string {
	return string(m.Serialize())
}
func NewAgentIdResponse(id string) *AgentIdResponse {
	m := AgentIdResponse{}
	m.MessageType = MSG_AGENT_ID_RESPONSE
	m.AgentId = id
	return &m
}

type jsonObject map[string]interface{}

func MsgType(raw []byte) (string, error) {
	var err error
	msg := jsonObject{}
	err = json.Unmarshal(raw, &msg)
	if err != nil {
		return "", err
	}
	msgType, ok := msg["type"].(string)
	if ok {
		return msgType, nil
	} else {
		return "", errors.New(fmt.Sprintf("msg type: expect string but got %T", msg["type"]))
	}
}
