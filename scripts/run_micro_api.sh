docker run -d -p 8080:8080 \
-e MICRO_REGISTRY=mdns \
microhq/micro api \
--handler=rpc \
--address=:8080 \
--namespace=shippy 