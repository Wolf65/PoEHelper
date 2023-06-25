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
	go build -o "out/PoEHelper.exe"

.PHONY: build1
build1:
	clear
	go build -ldflags='-extldflags="-static"' -o "out/PoEHelper1.exe"

.PHONY: build2
build2:
	clear
	go build -ldflags "-linkmode 'external' -extldflags '-static'" -o "out/PoEHelper2.exe"
	
#.PHONY: ldd
#ldd: ldd PoEHelper.exe