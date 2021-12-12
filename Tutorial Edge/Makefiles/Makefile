hello:
	echo "Hello From makefile"
build:
	go build -o bin/main main.go
compile:
	echo "Compiling for every OS and Platform"
	set GOOS=freebsd set GOARCH=386
	go build -o bin/main-freebsd-386 main.go

	set GOOS=linux set GOARCH=386
	go build -o bin/main-linux-386 main.go

	set GOOS=windows set GOARCH=386
	go build -o bin/main-windows-386.exe main.go

run:
	go run main.go

all: hello run build

clean:
	rmdir /s bin