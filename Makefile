.PHONY: compile
compile: compile_darwin compile_linux compile_windows compile_bsd

compile_darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o ./bin/envinject_darwin_amd64 ./main.go

compile_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./bin/envinject_linux_amd64 ./main.go

compile_windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o ./bin/envinject_windows_amd64 ./main.go

compile_bsd:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -a -o ./bin/envinject_bsd_amd64 ./main.go