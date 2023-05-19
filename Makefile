build:
	go build -o bin/
build-prod:
	go build -ldflags "-s -w" -o bin/ -tags prod
run:
	go run .
tests:
	go test ./...
clean:
	rm -rf bin/ muscurdi_db/

build-mac: clean
	mkdir bin
	sh scripts/update_version.sh
	fyne package -os darwin -icon assets/logo.png --release --tags prod
	mv muscurdig.app bin/
	sh scripts/restore_version.sh

build-linux: clean
	mkdir bin
	sh scripts/update_version.sh
	fyne package -os linux -icon assets/logo.png --release --tags prod
	mv muscurdig.tar.xz bin/
	sh scripts/restore_version.sh

build-win: clean
	mkdir bin
	fyne package -os windows -icon assets/logo.png --release --tags prod --appID me.vikkio.muscurdigwin
	mv muscurdig.exe bin/