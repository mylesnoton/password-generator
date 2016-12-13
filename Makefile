install:
	go build
	mv passgen /usr/local/bin
test:
	go test --cpu 4 --count 100 ./...