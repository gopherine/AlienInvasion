clean:
	rm -rvf ./bin

build: clean
	go build -o bin/alien-invasion ./cmd/main.go