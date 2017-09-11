// +build linux

package hostconfig

import (
	"git.oschina.net/k2ops/jarvis/protocol"
	"runtime"
	"os"
	"github.com/shirou/gopsutil/host"
	"git.oschina.net/k2ops/jarvis/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

func OsInfo() (osInfo protocol.OsInfo) {
	osInfo.OsType = runtime.GOOS
	osInfo.Arch = runtime.GOARCH
	osInfo.Hostname, _ = os.Hostname()
	osInfo.Uptime, _ = host.Uptime()
	// distribution
	osInfo.Dist, osInfo.Version = linuxDist()
	return
}

func linuxDist() (dist string, version string) {
	stdout, stderr, err := utils.OsExecLineOut("/usr/bin/lsb_release", "-a")
	if err != nil {
		log.WithError(err).Error(stderr)
		return
	}

	// convert stdout to a property map
	props := make(map[string]string)
	for _, line := range stdout {
		fields := strings.Split(line, ":")
		if len(fields)>=2 {
			props[strings.TrimSpace(fields[0])] = strings.TrimSpace(fields[1])
		}
	}
	log.Debug(props)

	//
	dist = props["Distributor ID"]
	version = props["Release"]
	return
}
