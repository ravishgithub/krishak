
# Krishak 🌾

Krishak is a lightweight Go-based microservice application to help manage agricultural land contracts. It features JWT-based authentication and a simple in-memory backend suitable for local and prototyping use.

---

## 🔐 Authentication

Use the `/login` endpoint to obtain a JWT token. The token must be passed via the `Authorization` header for all protected `POST` endpoints.

### Login (GET JWT token)
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"krishak2024"}'
```

### Check Token Validity
```bash
curl -X POST http://localhost:8080/check_auth \
  -H "Authorization: YOUR_TOKEN_HERE"
```

---

## 🔐 Protected Endpoints (require JWT)

### Add Contractor
```bash
curl -X POST http://localhost:8080/contractors \
  -H "Authorization: YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{"name":"Ravi Tiwari","contact":"9876543210","aadhar":"123456789012"}'
```

### Add Land
```bash
curl -X POST http://localhost:8080/lands \
  -H "Authorization: YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{"size":2.5,"location":"Dhanha Bandhiya","soil_type":"Clay"}'
```

### Add Contract
```bash
curl -X POST http://localhost:8080/contracts \
  -H "Authorization: YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "contractor_id": "REPLACE_CONTRACTOR_ID",
    "land_id": "REPLACE_LAND_ID",
    "start_date": "2024-07-01",
    "end_date": "2025-06-30",
    "expected_yield": 90
  }'
```

---

## 🌐 Public `GET` Endpoints

### List Contractors
```bash
curl http://localhost:8080/list_contractors
```

### List Lands
```bash
curl http://localhost:8080/list_lands
```

### List Contracts
```bash
curl http://localhost:8080/list_contracts
```

---

## 🧪 Running the App Locally

```bash
export JWT_SECRET="krishakdevsupersecret"
go run main.go
```

---

## 📁 Project Structure

```
krishak/
├── authentication/       # JWT login and auth logic
├── handlers/             # Contractors, lands, contracts
├── configs/config.json   # Login credentials
├── main.go               # App entrypoint
├── Dockerfile            # For containerization
├── go.mod, go.sum
└── .gitignore
```

---

## 🐳 Docker Build & Run (optional)

```bash
docker build -t krishak-app .
docker run -p 8080:8080 -e JWT_SECRET=krishakdevsupersecret krishak-app
```

---

## 🗃 Roadmap

- 🔲 Add persistent DB (SQLite / Autonomous JSON DB)
- 🔲 Add static frontend hosted on OCI Object Storage
- 🔲 Migrate handlers into OCI Functions for serverless hosting
