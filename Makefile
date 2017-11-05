
BINARY = api
GOARCH = amd64

# Build the project
all: linux darwin windows

linux:
	GOOS=linux GOARCH=${GOARCH} go build -o ${BINARY}-linux-${GOARCH} api.go

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-darwin-${GOARCH} api.go

windows:
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-windows-${GOARCH}.exe api.go

.PHONY: linux darwin windows
