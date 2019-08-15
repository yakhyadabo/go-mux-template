default: build

build: 
	go build go-mux-template.go

run: 
	go run go-mux-template.go

install:	
	go install go-mux-template.go