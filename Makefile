APP_VERSION     := $(shell git describe --abbrev=0)
APP_NAME        := gotest
BUILD_VERSION   := $(shell git log -1 --oneline)
BUILD_TIME      := $(shell date "+%FT%T%z")
GIT_REVISION    := $(shell git rev-parse --short HEAD)
GIT_BRANCH      := $(shell git name-rev --name-only HEAD)
GO_VERSION      := $(shell go version)
SOURCE          := main.go
TARGET_DIR      := /usr/datacenter/gotest
GOOS            := $(shell go env GOOS)



all:
	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	GOOS=linux GOARCH=amd64							\
	go build -ldflags                           \
	"                                           \
	-X 'main.AppName=${APP_NAME}' \
	-X 'main.AppVersion=${APP_VERSION}' \
	-X 'main.BuildTime=${BUILD_TIME}' \
	-X 'main.GitRevision=${GIT_REVISION}' \
	-X 'main.GitBranch=${GIT_BRANCH}' \
	-X 'main.GoVersion=${GO_VERSION}' \
	-w -s                               \
	"                                           \
	-o ${APP_NAME} ${SOURCE}
mac:
	GOARCH=amd64						\
	go build -ldflags                           \
	"                                           \
	-X 'main.AppName=${APP_NAME}' \
	-X 'main.AppVersion=${APP_VERSION}' \
	-X 'main.BuildVersion=${BUILD_VERSION}' \
	-X 'main.BuildTime=${BUILD_TIME}' \
	-X 'main.GitRevision=${GIT_REVISION}' \
	-X 'main.GitBranch=${GIT_BRANCH}' \
	-X 'main.GoVersion=${GO_VERSION}' \
	-w -s                               \
	"                                           \
	-o ${APP_NAME} ${SOURCE}

clean:
	rm ${APP_NAME} -f

install:
	mkdir -p ${TARGET_DIR}
	cp ${APP_NAME} ${TARGET_DIR} -f
