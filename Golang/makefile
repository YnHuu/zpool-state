BIN="zfsweb"
VERSION="v0.0.1"
BUILDTIME="$(shell date +'%Y-%m-%d_%T')"

clean:
	@rm ${BIN}

build:
	@echo '${BIN} building...'
	@echo "package version" > $(shell pwd)/version/version.go
	@echo 'const Version = ${VERSION}' >> $(shell pwd)/version/version.go
	@echo 'const BuildTime = ${BUILDTIME}' >> $(shell pwd)/version/version.go
	@env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=nomsgpack -ldflags '-s -w' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ${BIN} main.go
	@du -sh ${BIN}

upx: build
	@upx -9 ${BIN}

dev:
	@go run .