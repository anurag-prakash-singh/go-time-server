
all:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o time-server

