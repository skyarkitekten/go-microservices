build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/skyarkitekten/go-microservices/vessel-service \
    proto/vessel/vessel.proto
	docker build -t vessel-service .

run:
	docker run -d --net="host" \
		-p 50053 \
		-e MICRO_SERVER_ADDRESS=:50053 \
		-e MICRO_REGISTRY=mdns \
		vessel-service

push:
	docker build -t gcr.io/chad-thomas-neudesic/vessel-service:latest .
	docker push gcr.io/chad-thomas-neudesic/vessel-service:latest
