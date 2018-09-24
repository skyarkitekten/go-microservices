curl -XPOST -H 'Content-Type: application/json' \
    -d '{
      "service": "shippy.consignment",
      "method": "ShippingService.CreateConsignment",
      "request": {
        "description": "This is a test",
        "weight": 500,
        "containers": []
      }
    }' --url http://localhost:8080/rpc