// +build linux

package hostconfig

import (
	"git.oschina.net/k2ops/jarvis/protocol"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"bytes"
	"strings"
	"strconv"
)

func PhysicalDisks() (disks protocol.HostDisks, err error) {
	lsblk, err := exec.LookPath("/bin/lsblk")
	if err != nil {
		return
	}

	cmd := exec.Command(lsblk, "-bdno", "NAME,MODEL,SIZE")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}
	log.Debug(cmd.Args)

	for _, line := range strings.Split(out.String(), "\n") {
		log.Debug(line)
		fields := strings.Fields(line)
		if len(fields)>=3 {
			size, err := strconv.ParseUint(fields[2], 10, 64)
			if err != nil {
				size = 0
			}
			disks = append(disks, protocol.DiskInfo{
				Device: "/dev/"+ fields[0],
				Model: fields[1],
				Capacity: size,
			})
		}
	}

	return disks, nil
}