#grpc/client.go
curl "http://localhost:8080/user?id=123"
curl -X POST http://localhost:8081/graphql \
     -H "Content-Type: application/json" \
     -d '{"query":"{ user(id:\"123\"){ id name } }"}'
