# jarvis agent 

jarvis agent for hardware and OS metrics exposed by \*NIX kernels, written
in Go with pluggable metric collectors.

## Building and running 

```bash
    go get git.oschina.net/k2ops/jarvis 
    cd ${GOPATH-$HOME/go}/src/git.oschina.net/k2ops/jarvis/agent
    make && cd build
    ./jarvis-agent --master <master> 
```

## release a deb package and deploy by ansible
In Makefile we use dpkg-buildpackage command to release a deb package. Just
use "make deb-release" to run it.
```bash
    make deb-release
    make deploy
```
By default, we use dns to indicat master of jarvis, to config nameserver of
host by ansible,use:
```bash
    make dns
```
To remove deb package,use
```bash
    make remove
```
