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

	fmt.Println(NewHeartbeatMessage("temp agent").ToJsonString())

}

func ExampleNewHeartbeatMessage() {
	fmt.Println(NewHeartbeatMessage("temp agent").ToJsonString())
}

func ExampleNewWelcomeMessage() {
	fmt.Println(NewWelcomeMessage("client", "server").ToJsonString())
	// Output: {"type":"welcome","clientAddr":"client","serverAddr":"server"}
}
