arm:
	mkdir -p ./build
	GOOS=linux GOARCH=arm go build -o build/goodixgt911go main.go

clean:
	rm -Rf ./build
