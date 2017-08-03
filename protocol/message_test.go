package protocol

import (
	"testing"
	//"fmt"
	"fmt"
)

func TestMessages(t *testing.T) {
	m := JarvisMessage{}
	m.MessageType = "testType"
	fmt.Println(m.ToJsonString())

	fmt.Println(NewWelcomeMessage("client", "server").ToJsonString())

	fmt.Println(NewEmptyRegisterMessage().ToJsonString())

	fmt.Println(NewHeartbeatMessage().ToJsonString())

	fmt.Println(NewEmptyResourceUsageMessage().ToJsonString())

}

func ExampleNewEmptyRegisterMessage() {
	fmt.Println(NewEmptyRegisterMessage().ToJsonString())
	// Output: {"type":"register","updatedAt":"","osType":"","arch":"","hostname":"","cpuNum":0,"memTotal":0,"upTime":"","disks":null}
}

func ExampleNewEmptyResourceUsageMessage() {
	fmt.Println(NewEmptyResourceUsageMessage().ToJsonString())
	// Output: {"type":"resource-usage","updatedAt":"","cpuPercent":0,"memUsed":0,"disks":null,"network":0}
}

func ExampleNewHeartbeatMessage() {
	fmt.Println(NewHeartbeatMessage().ToJsonString())
}

func ExampleNewWelcomeMessage() {
	fmt.Println(NewWelcomeMessage("client", "server").ToJsonString())
	// Output: {"type":"welcome","clientAddr":"client","serverAddr":"server"}
}



