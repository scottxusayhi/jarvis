package tcp

import (
	"net"
	"strings"
	"errors"
	"fmt"
)

var allAgent map[string]JarvisAgent

type JarvisAgent struct {
	Datacenter string
	Rack string
	Slot string
	Conn net.Conn
}

func (a *JarvisAgent) UpdatePosition (dc string, rack string, slot string) {
	oldKey := getKey(a.Datacenter, a.Rack, a.Slot)
	a.Datacenter = dc
	a.Rack = rack
	a.Slot = slot

	allAgent[getKey(a.Datacenter, a.Rack, a.Slot)] = allAgent[oldKey]
	delete(allAgent, oldKey)
}

func getKey(part ...string) string {
	return strings.Join(part, "")
}

func RegisterNewAgent (dc string, rack string, slot string, conn net.Conn) {
	allAgent[getKey(dc, rack, slot)] = JarvisAgent{
		Datacenter: dc,
		Rack: rack,
		Slot: slot,
		Conn: conn,
	}
}

func GetAgent(dc string, rack string, slot string) (*JarvisAgent, error) {
	result, ok := allAgent[getKey(dc, rack, slot)]
	if !ok {
		return nil, errors.New("no such agent found. " + fmt.Sprintf("datacenter=%v rack=%v slot=%v", dc, rack, slot))
	}
	return &result, nil
}

func AgentExists(dc string, rack string, slot string) bool {
	_, ok := allAgent[getKey(dc, rack, slot)]
	return ok
}

