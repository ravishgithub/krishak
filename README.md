# Krishak 🌾

![CI](https://github.com/ravishgithub/krishak/actions/workflows/go-ci.yml/badge.svg?branch=main)

Krishak is a lightweight Go-based microservice application to help manage agricultural land contracts. It features JWT-based authentication and a simple in-memory backend suitable for local and prototyping use.

...

## 🧪 Unit Testing with Coverage

This project includes unit tests for:
- Authentication
- Handlers (contractor, land, contracts)

Run tests locally:
```bash
export JWT_SECRET="krishakdevsupersecret"
go test -cover ./...
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
├── main.go
├── .github/
│   └── workflows/
│       └── go-ci.yml
├── go.mod
├── go.sum
└── README.md
```
