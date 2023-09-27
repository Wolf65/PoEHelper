Name=PoE Helper

.PHONY: all
all: run

.PHONY: run
run:
	clear
	go run .

.PHONY: build
build:
	clear
	GOOS=windows GOARCH=amd64 go build -ldflags='-s -w -extldflags="-static"' -o "_out/PoEHelper.exe"

.PHONY: debug
debug:
	clear
	GOOS=windows GOARCH=amd64 go build -gcflags=all="-N -l" -ldflags='-extldflags="-static"' -o "_out/PoEHelperDebug.exe"

.PHONY: gdb
gdb:
	clear
	gdb -silent ./_out/PoEHelperDebug.exe