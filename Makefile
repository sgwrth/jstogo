main.exe: ./src/main.go ./src/structs/ast.go
	go build -o ./build/ ./src/main.go ./src/structs/ast.go
