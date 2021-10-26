GO_VERSION=$(shell go version)
GIT_COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME=$(shell date)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOBINPATH=$(shell which go)
GO_PATH=$(abspath ../../../..)

BUILD_COMMAND="$(GOBINPATH) build -v -ldflags \"-s -w -X 'github.com/lxyzhangqing/gpu-memory-monitor/version.goVersion=$(GO_VERSION)' -X 'github.com/lxyzhangqing/gpu-memory-monitor/version.gitCommit=$(GIT_COMMIT)' -X 'github.com/lxyzhangqing/gpu-memory-monitor/version.branch=$(BRANCH)' -X 'github.com/lxyzhangqing/gpu-memory-monitor/version.buildTime=$(BUILD_TIME)' -X 'github.com/lxyzhangqing/gpu-memory-monitor/version.oSArch=$(GOOS)/$(GOARCH)'\" ."


all:	export GOPATH=$(GO_PATH)
all:
	@echo "begin build gpu-memory-monitor ................"
	@echo ""
	@echo "Go version: " $(GO_VERSION)
	@echo "Git commit: " $(GIT_COMMIT)
	@echo "Branch:     " $(BRANCH)
	@echo "Build time: " $(BUILD_TIME)
	@echo "OS/Arch:    " $(GOOS)/$(GOARCH)
	@echo "GOPATH:     " $(GO_PATH)

	@echo $(BUILD_COMMAND)
	@bash -c $(BUILD_COMMAND)

clean:
	rm -f gpu-memoroy-monitor

