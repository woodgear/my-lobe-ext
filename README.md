```bash
curl -X POST http://localhost:3400/api/gateway \
  -H "Content-Type: application/json" \
  -d '{
    "id": "xx",
    "type": "api",
    "apiName": "currentTime",
    "arguments": "{\"ok\": true}",
    "identifier": "xx"
  }' 

curl -X POST http://localhost:3400/api/gateway \
  -H "Content-Type: application/json" \
  -d '{
    "id": "xx",
    "type": "api",
    "apiName": "listZZ",
    "arguments": "{}",
    "identifier": "xx"
  }' 
curl -X POST http://localhost:3400/api/gateway \
  -H "Content-Type: application/json" \
  -d '{
    "id": "xx",
    "type": "api",
    "apiName": "readZZ",
    "arguments": "{\"paths\": [\"250618/250618.json\"]}",
    "identifier": "xx"
  }' 
```
