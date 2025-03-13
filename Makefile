build:
	@CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/arithmo main.go

buildwin:
	@CGO_ENABLED=0 GOOS=windows go build -ldflags "-s -w" -o bin/arithmo.exe main.go

run:
	@go run main.go
