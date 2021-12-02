all: testapp1
	docker build -t sokoide/testapp1 .

testapp1: go.mod main.go
	CGO_ENABLED=0 go build -o testapp1

push: testapp1
	docker push sokoide/testapp1:latest

clean:
	rm testapp1
