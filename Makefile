
#include .env
PROJECT_NAME	 := $(shell basename "$(PWD)")
TIMESTAMP    	 := $(shell /bin/date "+%F %T")
VERSION			 := 1.0.0

GOFILES_NOVENDOR := $(shell go list ./... | grep -v /vendor/)

# Use linker flags to provide version/build settings
# LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
LDFLAGS		:= -s -w
#LDFLAGS		:= -s -w \
#			   -X 'main.BuildVersion=$(VERSION)' \
#			   -X 'main.BuildGitBranch=$(shell git describe --all)' \
#			   -X 'main.BuildGitRev=$(shell git rev-list --count HEAD)' \
#			   -X 'main.BuildGitCommit=$(shell git rev-parse HEAD)' \
#			   -X 'main.BuildDate=$(shell /bin/date "+%F %T")'

# Make is verbose in Linux. Make it silent.
# MAKEFLAGS += --silent

help: Makefile
	@echo " 选择要执行的命令 "$(PROJECT_NAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## ts: 输出时间戳
ts:
	@echo "$(TIMESTAMP)"

## fmt: 格式化项目文件
fmt:
	@go fmt $(GOFILES_NOVENDOR)
#	@go fmt $(CURDIR)/...
	@go mod tidy

## clean: 清理
clean:
	@rm -rf $(CURDIR)/bin/* &> /dev/null

## commit: 快速提交
commit: clean fmt
	git add .
	git commit -m "quick save @$(TIMESTAMP)"
	git push

## run: 执行当前项目
run:
	@go run main.go

## build: 编译当前项目
build:
	@go build -x -race -ldflags "$(LDFLAGS)" -o ./bin/$(PROJECT_NAME) .

## buildall: 编译当前项目，生成全平台binary
buildall:
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/bin/$(PROJECT_NAME)-linux-amd64-$(VERSION)
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/bin/$(PROJECT_NAME)-linux-macos-$(VERSION)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/bin/$(PROJECT_NAME)-linux-windows-$(VERSION).exe


.PHONY: fmt help