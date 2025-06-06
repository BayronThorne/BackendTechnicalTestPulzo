# /health endpoint response
OK

# Example response for /auth
{
  "token": "your_token_here"
}

# Example response for /data (valid token)
{
  "info": {
    "count": 826,
    "pages": 42,
    "next": "https://rickandmortyapi.com/api/character?page=2",
    "prev": null
  },
  "results": [
    {
      "id": 1,
      "name": "Rick Sanchez",
      "status": "Alive",
      "species": "Human",
      ...
    }
  ]
}

# Example response for /data (invalid token)
Invalid token format

# Example response for /data (expired token)
Token expired

# Example response for /data (no token)
Token required