

build:
	go build -x .
	cp ./ctrlc ./minigui/

buildx:
	GOOS="darwin" GOARCH="amd64" go build .
	cp ./ctrlc ./minigui/ctrlc.app/Contents/MacOS/
	