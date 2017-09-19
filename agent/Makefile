GO     ?= GO15VENDOREXPERIMENT=1 go
GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
GOARCH := $(shell $(GO) env GOARCH)

STATICCHECK ?= $(GOPATH)/bin/staticcheck
pkgs         = $(shell $(GO) list ./... | grep -v /vendor/)

PREFIX                  ?= $(shell pwd)
BIN_DIR                 ?= $(shell pwd)
MACH                    ?= $(shell uname -m)

ifeq ($(OS),Windows_NT)
    OS_detected := Windows
else
    OS_detected := $(shell uname -s)
endif


all: format vet staticcheck build

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

staticcheck: $(STATICCHECK)
	@echo ">> running staticcheck"
	@$(STATICCHECK) $(pkgs)

build: $(PROMU)
	@echo ">> building binaries"
	@$(GO) build -o build/jarvis_agent_linux_amd64.bin

install: 
	@echo ">> install binaries"
	install -D build/*  /usr/sbin/jarvis-agent 
	install -D -m 664 etc/default/* /etc/default/jarvis-agent
	install -D etc/init.d/* /etc/init.d/jarvis-agent

deb-release:
	@echo ">> release deb package"
	tar -czf ../jarvis+agent-0.1.tar.gz  ./
	mkdir -p releases/jarvis+agent-0.1 && mv ../jarvis+agent-0.1.tar.gz releases/ 
	cd releases/jarvis+agent-0.1 && dh_make -s -e ludun@k2data.com.cn -f ../jarvis+agent-0.1.tar.gz && dpkg-buildpackage -rfakeroot -uc -us
	
clean:
	-rm -rf build
	-rm -rf releases
	-rm -rf description-pak
	
$(GOPATH)/bin/staticcheck:
	@GOOS= GOARCH= $(GO) get -u honnef.co/go/tools/cmd/staticcheck


.PHONY: all style format build vet install staticcheck

# Declaring the binaries at their default locations as PHONY targets is a hack
# to ensure the latest version is downloaded on every make execution.
# If this is not desired, copy/symlink these binaries to a different path and
# set the respective environment variables.
.PHONY: $(GOPATH)/bin/staticcheck
