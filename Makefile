build:
	go build -o bin/
run:
	go run .
tests:
	go test ./...
clean:
	rm -rf bin/ muscurdi_db/