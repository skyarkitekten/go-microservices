build:
	docker build -t consignment-cli .

run:
	docker run --net="host" \
		-e MICRO_REGISTRY=mdns \
		consignment-cli \
		consignment.json \
		31cc7877-4f43-4e03-9147-52d66215ab31