PHONY: build clean

build: build/upmon

build/upmon: src/main.go
	go build -o $@ $<

clean:
	rm -rf build
