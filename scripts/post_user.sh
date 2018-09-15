curl -XPOST -H 'Content-Type: application/json' \
-d '{ "service": "shippy.auth", "method": "Auth.Create", "request": { "user": { "email": "ewan.valentine89@gmail.com", "password": "testing123", "name": "Ewan Valentine", "company": "BBC" } } }' \
http://localhost:8080/rpc
