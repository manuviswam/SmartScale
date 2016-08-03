build:
	go build -o out/smartscale main/main.go
	node_modules/.bin/webpack -p