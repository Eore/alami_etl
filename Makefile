appName := processor

build:
	go build -o ${appName} cmd/main.go