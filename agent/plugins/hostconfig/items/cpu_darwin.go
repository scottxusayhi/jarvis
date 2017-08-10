// +build darwin

package items

import (
	"git.oschina.net/k2ops/jarvis/protocol"
	"runtime"
)

const (
	CPU_VENDOR_INTEL = "GenuineIntel"
)

func CpuInfo() (cpuInfo protocol.CpuInfo) {
	cpuInfo.Vcpu = runtime.NumCPU()
	return
}
