build:
	docker build -t email-service .

run:
	docker run --net="host" \
		-d \
		-p 50054 \
		-e MICRO_SERVER_ADDRESS=:50054 \
		-e MICRO_REGISTRY=mdns \
		email-service email_service