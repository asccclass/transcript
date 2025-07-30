
clean:
	go clean -modcache
	go mod tidy
	
build-win: clean
	go build -o app.exe

s:
	git push -u origin main