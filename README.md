# Krishak 🌾

![CI](https://github.com/ravishgithub/krishak/actions/workflows/go-ci.yml/badge.svg?branch=main)

Krishak is a lightweight Go-based microservice application to help manage agricultural land contracts. It features JWT-based authentication and a simple in-memory backend suitable for local and prototyping use.

---

## 🌱 Features

- `/login`: Authenticates a user and returns a signed JWT token
- `/check_auth`: Validates JWT token from the Authorization header
- `/contractors`, `/lands`, `/contracts`: APIs to manage agri-land records

---

## 🔐 Secure Configuration

You must define a secure `JWT_SECRET` before starting the app.

### ✅ Generate a Secure JWT Secret

```bash
openssl rand -hex 32
```

Set it in your environment:

**Linux/macOS:**
```bash
export JWT_SECRET="your_generated_secret"
```

**Windows CMD:**
```cmd
set JWT_SECRET=your_generated_secret
```

---

## ⚙️ Configuration File (`configs/config.json`)

```json
{
  "login": {
    "username": "admin",
    "password": "<bcrypt_hash_of_your_password>"
  },
  "server": {
    "port": 8080,
    "hostname": "localhost"
  },
  "database": {
    "username": "admin",
    "password": "admin123",
    "name": "mydatabase"
  }
}
```

### 🔑 Generate a bcrypt password hash

To generate a secure bcrypt hash of your password:

```bash
go run scripts/genhash.go
```

Copy the generated hash and replace it in your `configs/config.json` under the `"password"` field.

---

## 🚀 Running the App

From the project root:

```bash
go run main.go
```

---

## 📬 API Testing with Curl

### Login Request

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"your_password_here"}'
```

### Auth Check Request

```bash
curl -X POST http://localhost:8080/check_auth \
  -H "Authorization: <JWT_TOKEN>"
```

---

## 🧪 Unit Testing with Coverage

This project includes unit tests for:
- Authentication
- Handlers (contractor, land, contracts)

Run tests locally:

```bash
export JWT_SECRET="ebee1a4380a9ab9a0a84b091c1f7abcf30c3428608f122dbd91e13db134b16bc"
go test -cover ./...
```

---

## 🧪 API Testing Script

A helper script `scripts/test_krishak_api.sh` is available to simulate a full login and business flow.

### Purpose

The script automates end-to-end testing of your APIs. It:

- Authenticates via `/login` and retrieves a JWT token
- Sends authenticated POST requests to:
  - `/contractors` – Adds a contractor
  - `/lands` – Adds land information
  - `/contracts` – Links a contractor and land

### Prerequisite

Install `jq` if not already installed:

**macOS:**
```bash
brew install jq
```

**Ubuntu/Debian:**
```bash
sudo apt install jq
```

### Run Script

```bash
chmod +x ./scripts/test_krishak_api.sh
./scripts/test_krishak_api.sh
```

---

## 🐳 Docker (Optional)

```bash
docker build -t krishak-login-app .
docker run -p 8080:8080 -e JWT_SECRET=your_generated_secret krishak-login-app
```

---

## 📁 Project Structure

```
krishak/
├── authentication/
│   └── loginapi.go
├── handlers/
│   ├── contractors.go
│   ├── lands.go
│   └── contracts.go
├── configs/
│   └── config.json
├── scripts/
│   ├── test_krishak_api.sh
│   └── genhash.go
├── main.go
├── .github/
│   └── workflows/
│       └── go-ci.yml
├── go.mod
├── go.sum
└── README.md
```

---

## 🔐 Security Notes

- Do **not** commit real JWT secrets or passwords
- Always use strong secrets for production environments
- This README uses **placeholders** for guidance

---

## 🧩 License

MIT License