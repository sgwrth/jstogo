main.exe: ./src/main.go
	go build -o ./build/ ./src/main.go

test:
	@go test ./src/**/
