docker run -d -p 8080:8080 \
-e MICRO_REGISTRY=mdns \
--handler=rpc \
--address=:8080 \
--namespace=shippy \
microhq/micro api