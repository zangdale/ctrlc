

build:
	go build .
	cp ./ctrlc ./minigui/ctrlc.app/Contents/MacOS/

buildx:
	GOOS="darwin" GOARCH="amd64" go build .
	cp ./ctrlc ./minigui/ctrlc.app/Contents/MacOS/
	