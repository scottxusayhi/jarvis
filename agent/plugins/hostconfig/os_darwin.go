// +build darwin

package hostconfig

import (
	"github.com/scottxusayhi/jarvis/protocol"
	"runtime"
	"os"
	"github.com/shirou/gopsutil/host"
)

func OsInfo() (osInfo protocol.OsInfo) {
	osInfo.OsType = runtime.GOOS
	osInfo.Arch = runtime.GOARCH
	osInfo.Hostname, _ = os.Hostname()
	osInfo.Uptime, _ = host.Uptime()
	return
}