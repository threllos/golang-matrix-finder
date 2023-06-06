run: build
	./build/app

build:
	mkdir -p build
	go build -o ./build/app cmd/app/app.go

bench:
	go test -benchmem -bench . ./internal/task/ 

clean:
	rm -rf ./build

.PHONY: clean run build bench clean