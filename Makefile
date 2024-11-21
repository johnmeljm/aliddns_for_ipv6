
APP=aliddns

.PHONY: build
build: clean
	GOARCH=amd64 GOOS=linux go build -o ./bin/${APP}_linux main.go gconfig.go aliddns.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${APP}_win.exe main.go gconfig.go aliddns.go
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${APP}_darwin_amd64 main.go gconfig.go aliddns.go
	GOARCH=arm64 GOOS=darwin go build -o ./bin/${APP}_darwin_arm64 main.go gconfig.go aliddns.go

.PHONY: run
run:
	go run -race main.go gconfig.go aliddns.go

.PHONY: clean
clean:
	go clean
	rm -rf ./bin
