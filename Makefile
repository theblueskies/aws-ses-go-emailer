linuxbuild:
	rm -rf bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/emailer .

build:
	rm -rf bin
	go build -o bin/emailer .

localrun:
	heroku local web
