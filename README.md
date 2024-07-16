Sample Requests: 
GET: curl -v localhost:9090
POST: curl -i -X POST -H "Content-Type: application/json" -d "{\"ID\":3, \"name\": \"Alex\", \"location\" : \"Toronto\"}" localhost:9090
PUT: curl -i -X PUT -H "Content-Type: application/json" -d "{\"ID\":1, \"name\": \"Alex\", \"location\" : \"Toronto\"}" localhost:9090/1
