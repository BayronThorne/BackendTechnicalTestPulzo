# Test /health endpoint
curl -i -X GET http://localhost:8080/health

# Get token
curl -X POST http://localhost:8080/token -H "Content-Type: application/json"

# Use token
curl -X GET http://localhost:8080/characters -H "Authorization: Bearer your_token_here"
