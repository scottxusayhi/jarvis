// +build linux

package items

import (
	"git.oschina.net/k2ops/jarvis/protocol"
	log "github.com/sirupsen/logrus"
	"strings"
	"strconv"
	"git.oschina.net/k2ops/jarvis/utils"
	"runtime"
	"fmt"
)

const (
	CPU_VENDOR_INTEL = "GenuineIntel"
)

func CpuInfo() (cpuInfo protocol.CpuInfo) {
	cpuInfo.Vcpu = runtime.NumCPU()
	stdout, stderr, err := utils.OsExecLineOut("/usr/bin/lscpu")
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
	// construct cpuInfo
	vcpu, _ := strconv.ParseInt(props["CPU(s)"], 10, 8)
	socket, _ := strconv.ParseInt(props["Socket(s)"], 10, 8)
	cpuInfo.Vcpu = int(vcpu)
	cpuInfo.Socket = int(socket)
	cpuInfo.Model = fmt.Sprintf("%v %v family %v model %v stepping %v",
		props["Vendor ID"],
		props["Architecture"],
		props["CPU family"],
		props["Model"],
		props["Stepping"],
	)
	return
}
