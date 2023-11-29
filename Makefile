build:
	go build

install:
	go install

run: build
	./sknowR

bi: 
	go build
	go install

clean:
	rm -f sknowR

test:
	go test -v ./...