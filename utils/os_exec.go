package utils

import (
	"os/exec"
	"bytes"
	log "github.com/sirupsen/logrus"
	"strings"
)

func OsExec(name string, args ...string) (stdout bytes.Buffer, stderr bytes.Buffer, err error) {
	// locate binary
	bin, err := exec.LookPath(name)
	if err != nil {
		return
	}
	// run
	cmd := exec.Command(bin, args...)
	log.Debug(cmd.Args)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	log.Debug(stdout.String())
	log.Debug(stderr.String())
	return
}

func OsExecLineOut(name string, args ...string) (outlines []string, errlines []string, err error) {
	stdout, stderr, err := OsExec(name, args...)
	return strings.Split(stdout.String(), "\n"), strings.Split(stderr.String(), "\n"), err
}
