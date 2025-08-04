LDFLAGS := "-O2" "-g"

clean:
	go clean -modcache
	go clean -i -r
	go mod tidy

build-win-debug:
	cls
	set GOOS=windows
	set GOARCH=amd64
	go build  -x -v -o app.exe
	go build -ldflags="-linkmode 'external' -extldflags '-static'" -o app.exe
	
build-win:
	cls
	go build -ldflags="-linkmode 'external'" -o app.exe

build: build-win


s:
	git push -u origin main