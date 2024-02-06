build:
	go build -o bin/main ./cmd/playlist-maker/main.go

run:
	go run ./cmd/playlist-maker/main.go

clean:
	go clean
