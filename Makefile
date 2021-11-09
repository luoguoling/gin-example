.PHONY: all build run gotool clean help
BINARY="bluebell"
all: gotool build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./main.go conf/config.yaml
gotool:
	go fmt ./
	go vet ./
clean:
	@if [ -f %{BINARY} ];then rm ${BINARY}; fi

help:
	@echo "make - 格式化 Go 代码，并编译生成二进制文件"
	@echo "make build - 编译Go代码,生成二进制"
	@echo "make run ,直接运行go代码"
	@echo "make clean -移除二进制文件和vim swap files"
	@echo "make gotool-运行go工具'fmt' and 'vet'"
