package plugins

import (
	"fmt"
	"testing"
)

func TestDetect(t *testing.T) {
	fmt.Println(NewHostConfigMessage().Serialize())
}
