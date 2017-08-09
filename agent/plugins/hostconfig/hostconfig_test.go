package hostconfig

import (
	"testing"
	"fmt"
)

func TestDetect(t *testing.T) {
	fmt.Println(NewHostConfigMessage().Serialize())
}

