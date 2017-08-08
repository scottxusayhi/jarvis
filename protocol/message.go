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
	MSG_REGISTER          = "register"
	MSG_HEARTBEAT         = "heartbeat"
	MSG_RESOURCE_USAGE    = "resource-usage"
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

// client register message
type PhysicalDiskInfo struct {
	Device string `json:"device"`
	Total  uint64 `json:"total"`
	Used   uint64 `json:"used"`
}
type registerMessage struct {
	JarvisMessage
	UpdatedAt string             `json:"updatedAt"`
	OSType    string             `json:"osType"`
	Arch      string             `json:"arch"`
	Hostname  string             `json:"hostname"`
	CPUNum    int                `json:"cpuNum"`
	MemTotal  int                `json:"memTotal"`
	UpTime    string             `json:"upTime"`
	Disks     []PhysicalDiskInfo `json:"disks"`
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

// resource usage message
type resourceUsageMessage struct {
	JarvisMessage
	UpdatedAt  string             `json:"updatedAt"`
	CPUPercent float32            `json:"cpuPercent"`
	MemUsed    uint64             `json:"memUsed"`
	Disks      []PhysicalDiskInfo `json:"disks"`
	Network    uint               `json:"network"`
}

func (m *resourceUsageMessage) Serialize() []byte {
	return serialize(m)
}
func (m *resourceUsageMessage) ToJsonString() string {
	return string(m.Serialize())
}
func NewEmptyResourceUsageMessage() *resourceUsageMessage {
	m := resourceUsageMessage{}
	m.MessageType = "resource-usage"
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
