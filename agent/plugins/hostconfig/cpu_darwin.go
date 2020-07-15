// +build darwin

package hostconfig

import (
	"github.com/scottxusayhi/jarvis/protocol"
	"runtime"
)

const (
	CPU_VENDOR_INTEL = "GenuineIntel"
)

func CpuInfo() (cpuInfo protocol.CpuInfo) {
	cpuInfo.Vcpu = runtime.NumCPU()
	return
}
