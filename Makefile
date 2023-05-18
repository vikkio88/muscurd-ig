build:
	go build -o bin/
build-prod:
	go build -ldflags "-s -w" -o bin/
run:
	go run .
tests:
	go test ./...
clean:
	rm -rf bin/ muscurdi_db/

build-mac:
	fyne package -os darwin -icon assets/logo.png --release
	mv muscurdig.app bin/