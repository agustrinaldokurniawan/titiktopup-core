# TitikTopup Core

Backend service for **TitikTopup** – a simple top-up platform with REST + gRPC APIs built in Go, using PostgreSQL and gRPC-Gateway.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)

---

## Features

- ✅ **Dual Protocol**: REST API (via gRPC-Gateway) and gRPC server.
- ✅ **Auto-Generated Docs**: Interactive Swagger/OpenAPI UI.
- ✅ **UUID Support**: Secure transaction IDs using PostgreSQL `gen_random_uuid`.
- ✅ **Persistence**: GORM ORM with PostgreSQL.
- ✅ **Easy Dev**: One-command database seeding & containerization.
---

## Tech Stack

| Technology   | Version | Purpose                          |
| ------------ | ------- | -------------------------------- |
| Go           | 1.24+   | Backend language                 |
| PostgreSQL   | Latest  | Primary database                 |
| gRPC         | Latest  | Service-to-service communication |
| gRPC-Gateway | Latest  | REST API gateway                 |
| GORM         | Latest  | ORM for database operations      |
| Docker       | Latest  | Containerization                 |

---

## Prerequisites

- **Go** 1.24 or higher
- **Docker** & **Docker Compose**
- **Git** and Bash shell (Git Bash recommended on Windows)

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/agustrinaldokurniawan/titiktopup-core.git
cd titiktopup-core
```

### 2. Set Up Environment Variables

Create a `.env` file in the project root. This file is **not committed** to version control:

```env
# PostgreSQL Container Configuration
POSTGRES_USER=topupuser
POSTGRES_PASSWORD=topuppass
POSTGRES_DB=titiktopup

# Application Database Connection
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=topupuser
DB_PASSWORD=topuppass
DB_NAME=titiktopup
```

**⚠️ Note:** Update credentials with secure values for production environments.

### 3. Start PostgreSQL

```bash
# Start services
./docker.sh up

# View database logs
./docker.sh logs

# Stop services
./docker.sh down

# Reset everything (containers + volumes)
./docker.sh reset
```

### 4. Run Migrations & Seed Data

```bash
./run.sh seed
```

This will:

- Run database migrations
- Load seed data (categories, products, etc.)

### 5. Start the API Server

```bash
./run.sh
```

The server will be available at:

| Service  | URL                   |
| -------- | --------------------- |
| REST API | http://localhost:8080 |
| gRPC     | localhost:50051       |

---

## Project Structure

```
titiktopup-core/
├── cmd/
│   ├── api/                 # API server entrypoint
│   └── seed/                # Database seeder entrypoint
├── internal/
│   ├── config/              # Database & app configuration
│   ├── server/              # HTTP + gRPC server setup
│   ├── handler/             # Request handlers (HTTP/gRPC)
│   ├── repository/          # Data access layer (PostgreSQL)
│   ├── domain/              # Domain models & business logic
│   └── clients/             # Internal gRPC clients
├── gen/
│   └── openapiv2/           # Generated OpenAPI/Swagger definitions
├── migrations/              # SQL or migration files
├── proto/                   # Protocol Buffer definitions (.proto)
├── pb/                      # Generated Go code from protobuf
├── docker-compose.yml       # Docker Compose configuration
├── docker.sh                # Docker management helper script
├── run.sh                   # Application runner (loads .env)
├── buf.yaml                 # Buf module configuration
├── buf.gen.yaml             # Buf code generation configuration
├── proto.sh                 # Helper script to regenerate pb/ code
├── .env                     # Local environment variables (Git-ignored)
├── .gitignore               # Git ignore rules
├── go.mod                   # Go module dependencies
├── go.sum                   # Go module checksums
└── README.md                # This file
```

### Key Directories

- **`cmd/`** – Executable entry points for API server and seeder
- **`internal/`** – Application business logic (not exported)
- **`proto/`** – Service definitions in Protocol Buffer format
- **`pb/`** – Auto-generated gRPC and HTTP gateway code
- **`migrations/`** – Database migration files (if you choose to use them explicitly)

---

## API Documentation

This project uses **Swagger UI** to provide real-time, interactive documentation. Since the API is under active development, please refer to the live docs for the most up-to-date endpoint list, request schemas, and response formats.

### 📖 Interactive Explorer
- **Live Docs:** [http://localhost:8080/docs](http://localhost:8080/docs)
- **OpenAPI Spec (JSON):** [http://localhost:8080/swagger.json](http://localhost:8080/swagger.json)

> **Pro Tip:** You can import the `swagger.json` URL directly into **Postman** or **Insomnia** to automatically generate a full request collection.

## Development Workflow

To add a new endpoint/feature:

1. Define the new service or message in `proto/titiktopup.proto`.
2. Run `./proto.sh` to regenerate the Go code and Swagger spec.
3. Implement the new handler in `internal/handler/`.
4. Register the logic in `internal/repository/`.
5. The new endpoint will automatically appear in the `/docs` UI.

### Regenerate Code
```bash
./proto.sh
---

## Development

### Running Tests

```bash
go test ./...
```

### Code Formatting

```bash
go fmt ./...
```

### Linting

```bash
golangci-lint run
```

---

## Quick Test

After starting the server, you can test the flow using `curl`:

### Check Menu
```bash
curl http://localhost:8080/api/v1/menu
```
   
---


## Troubleshooting

| Issue                       | Solution                                                   |
| --------------------------- | ---------------------------------------------------------- |
| Port 5432 already in use    | Change `DB_PORT` in `.env` or kill the conflicting process |
| Database connection refused | Ensure `docker.sh up` completed successfully               |
| Migrations failed           | Run `docker.sh reset` to start fresh                       |
| `.env` file not found       | Create it manually in the project root                     |

---

## Contributing

1. Create a feature branch: `git checkout -b feature/your-feature`
2. Make your changes and commit: `git commit -am 'Add feature'`
3. Push to branch: `git push origin feature/your-feature`
4. Submit a Pull Request

---

## License

[Add your license here - e.g., MIT, Apache 2.0, etc.]

---

## Support

For issues or questions, please open a GitHub issue or contact the maintainers.

---
