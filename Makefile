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
	GOOS=windows GOARCH=amd64 go build -ldflags='-extldflags="-static"' -o "out/PoEHelper.exe"

.PHONY: buildZip
buildZip:
	clear
	GOOS=windows GOARCH=amd64 go build -ldflags='-s -w -extldflags="-static"' -o "out/PoEHelperZip"

#.PHONY: ldd
#ldd: ldd PoEHelper.exe