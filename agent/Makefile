GO     ?= GO15VENDOREXPERIMENT=1 go
GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
GOARCH := $(shell $(GO) env GOARCH)

STATICCHECK ?= $(GOPATH)/bin/staticcheck
pkgs         = $(shell $(GO) list ./... | grep -v /vendor/)

PREFIX                  ?= $(shell pwd)
BIN_DIR                 ?= $(shell pwd)
MACH                    ?= $(shell uname -m)
VERSION			?= 0.1

ifeq ($(OS),Windows_NT)
    OS_detected := Windows
else
    OS_detected := $(shell uname -s)
endif


all: build install

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
	rm -rf debian/jarvis-agent/usr
	install -D build/*  debian/jarvis-agent/usr/sbin/jarvis-agent
	install -D -m 664 etc/default/* debian/jarvis-agent/etc/default/jarvis-agent
	install -D etc/init/* debian/jarvis-agent/etc/init/jarvis-agent.conf
	install -D etc/init.d/jarvis-debain debian/jarvis-agent/etc/init.d/jarvis-agent


dh_make:
	dh_make -s -e ludun@k2data.com.cn --native

deb-release:
	@echo ">> release deb package"
	mkdir -p ../releases/jarvis-agent-$(VERSION) && cp -R ../agent/* ../releases/jarvis-agent-$(VERSION)/
	cd ../releases/jarvis-agent-$(VERSION) && dpkg-buildpackage -rfakeroot -uc -us

rpm-release:
	@echo ">> release rpm package"
	cp -R ../agent/* ~/rpmbuild/BUILD/
	install -D -m 664 etc/default/* ~/rpmbuild/SOURCES/jarvis-agent-default
	install -D -m 664 etc/init.d/jarvis-rpm ~/rpmbuild/SOURCES/jarvis-agent-init
	rpmbuild -bb jarvis-agent.spec


clean:
	-rm -rf build
	-rm -rf ../releases

deploy:
	@echo ">> deploy deb package by ansible"
	-ansible all -m copy -a 'src=../releases/jarvis-agent_0.1_amd64.deb dest=/tmp/'
	-ansible all -m command -a 'dpkg -i /tmp/jarvis-agent_0.1_amd64.deb' -s
	-ansible all -m command -a 'service jarvis-agent start' -s
	-ansible all -m command -a 'rm -f /tmp/jarvis-agent_0.1_amd64.deb'

remove:
	ansible all -m command -a 'service jarvis-agent stop' -s
	ansible all -m command -a 'dpkg -P jarvis-agent' -s

dns:
	ansible all -m raw -a 'sed -i  "1a nameserver 10.10.90.230"  /etc/resolv.conf' -s

log:
	ansible all -m raw -a 'tail -10 /var/log/upstart/jarvis-agent.log' -s

ps:
	ansible all -m raw -a 'ps -ef|grep jarvis|grep -v sh' -s

$(GOPATH)/bin/staticcheck:
	@GOOS= GOARCH= $(GO) get -u honnef.co/go/tools/cmd/staticcheck


.PHONY: all style format build vet install staticcheck

# Declaring the binaries at their default locations as PHONY targets is a hack
# to ensure the latest version is downloaded on every make execution.
# If this is not desired, copy/symlink these binaries to a different path and
# set the respective environment variables.
.PHONY: $(GOPATH)/bin/staticcheck
