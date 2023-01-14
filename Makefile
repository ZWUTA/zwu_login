BINARY=zwu

build:
		go build -o ${BINARY}

install:
		go install

debug:
		go clean
		# build for linux_amd64
		GOOS=linux GOARCH=amd64 go build -o ${BINARY}_linux_amd64
		# build for linux_arm64
		GOOS=linux GOARCH=arm64 go build -o ${BINARY}_linux_arm64
		# build for windows_amd64
		GOOS=windows GOARCH=amd64 go build -o ${BINARY}_windows_amd64.exe
		# build for windows_arm64
		GOOS=windows GOARCH=arm64 go build -o ${BINARY}_windows_arm64.exe
		# build for darwin_amd64
		GOOS=darwin GOARCH=amd64 go build -o ${BINARY}_darwin_amd64
		# build for darwin_arm64
		GOOS=darwin GOARCH=arm64 go build -o ${BINARY}_darwin_arm64

release:
		go clean
		# build for linux_amd64
		GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ${BINARY}_linux_amd64 && upx -9 ${BINARY}_linux_amd64
		# build for linux_arm64
		GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o ${BINARY}_linux_arm64 && upx -9 ${BINARY}_linux_arm64
		# build for windows_amd64
		GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ${BINARY}_windows_amd64.exe && upx -9 ${BINARY}_windows_amd64.exe
		# build for windows_arm64
		GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" -o ${BINARY}_windows_arm64.exe && upx -9 ${BINARY}_windows_arm64.exe
		# build for darwin_amd64
		GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ${BINARY}_darwin_amd64 && upx -9 ${BINARY}_darwin_amd64
		# build for darwin_arm64
		GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o ${BINARY}_darwin_arm64 && upx -9 ${BINARY}_darwin_arm64

clean:
		go clean

.PHONY:  clean build