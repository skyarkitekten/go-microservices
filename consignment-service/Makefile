build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/skyarkitekten/go-microservices/consignment-service \
		proto/consignment/consignment.proto
	
	GOOS=linux GOARCH=amd64 go build
    
	docker build -t consignment-service .