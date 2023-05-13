build:
	go build -o bin/
run:
	go run .
tests:
	go test ./...
clean:
	rm -rf bin/ db_files/