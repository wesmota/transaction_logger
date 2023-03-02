# Basic Makefile for Golang project

test: 
	go test -v ./... -short

## Run test coverage and generate html report
test-cover:    
	rm -fr coverage
	mkdir coverage
	go list -f '{{if gt (len .TestGoFiles) 0}}"go test -covermode count -coverprofile {{.Name}}.coverprofile -coverpkg ./... {{.ImportPath}}"{{end}}' ./... | xargs -I {} bash -c {}
	echo "mode: count" > coverage/cover.out
	grep -h -v "^mode:" *.coverprofile >> "coverage/cover.out"
	rm *.coverprofile
	go tool cover -html=coverage/cover.out -o=coverage/cover.html

test-all: test test-cover

run: 
	go run main.go

compile:
	protoc api/v1/*.proto \
		--go_out=. \
		--proto_path=. \
		--go_opt=paths=source_relative \