
clean:
	go clean -modcache
	go mod tidy

build-win-debug:
	cls
	go build -x -v -gcflags="all=-N -l" -o app.exe
	
build-win:
	go build -o app.exe

s:
	git push -u origin main