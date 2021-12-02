all: testapp2
	docker build -t sokoide/testapp2 .

testapp2: go.mod main.go
	CGO_ENABLED=0 go build -o testapp2

push: testapp2
	docker push sokoide/testapp2:latest

clean:
	rm testapp2
