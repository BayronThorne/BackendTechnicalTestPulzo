# 🧪 Backend Technical Test

This project consists of developing two microservices written in Go:

1. **Authentication Service**: generates and manages temporary tokens (maximum of 5 uses).
2. **Data Fetcher Service**: retrieves character data from the public Rick and Morty API.

---

## 📁 Project Structure

```
backend-test/
├── .env.example
├── docker-compose.yml
├── Makefile
├── go.mod
├── go.sum
├── README.md
│
├── auth/
│   ├── Dockerfile
│   ├── main.go
│   ├── handler.go
│   └── token_manager.go
│
├── datafetcher/
    ├── Dockerfile
    ├── main.go
    └── handler.go
```

---

## ▶️ Running the Project

### 1. Clone the repository

```bash
git clone https://github.com/BayronThorne/BackendTechnicalTestPulzo.git
cd BackendTechnicalTestPulzo
```

### 2. Create the `.env` file (optional)

Duplicate the `.env.example` file and configure the required variables:

```bash
cp .env.example .env
```

---

## 🛠️ Using the Makefile

This project includes a `Makefile` to simplify Docker-related tasks.

### Available Commands

- `make build` – Builds the Docker images.
- `make up` – Starts the services in detached mode.
- `make down` – Stops and removes containers, networks, and volumes.
- `make test` – Runs all tests in both services.

### Example

```bash
make build
make up
make test
make down
```

---

## 📌 Endpoints

### Authentication Service

- `POST /token`  
  Generates a secure 64-character token valid for 5 uses.

- `GET /characters`  
  Requires a valid token in the `Authorization` header. If valid, fetches character data from the Data Fetcher Service.

- `GET /health`  
  Returns `OK` to indicate the service is running.

### Data Fetcher Service

- `GET /characters`  
  Proxies a request to the Rick and Morty public API and returns the character data.

- `GET /health`  
  Returns `OK` to indicate the service is running.

---

## 🧪 Running Tests

This project includes unit and HTTP tests for both services.

To run all tests:

```bash
make test
```

Or manually:

```bash
go test ./auth
go test ./datafetcher
```

---

## 🧠 Notes

- Tokens are generated securely using `crypto/rand` and have a fixed length of 64 characters.
- Each token can be used **up to 5 times**. On the 6th use, it will be considered expired.
- Token format and length are validated before any processing.

---

## 🌐 API Reference

[Rick and Morty API](https://rickandmortyapi.com/)