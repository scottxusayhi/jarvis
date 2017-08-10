// +build darwin

package items

import (
	"git.oschina.net/k2ops/jarvis/protocol"
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