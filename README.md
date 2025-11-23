# CSAT Survey Service

A Go-based microservice implementing Hexagonal Architecture (Ports & Adapters) for Customer Satisfaction (CSAT) surveys.

## 🏗️ Architecture

This project follows **Hexagonal Architecture** principles, separating business logic from external dependencies:

```
┌─────────────────────────────────────┐
│         Core Domain Layer           │
│  (Business Logic & Ports)           │
└─────────────────────────────────────┘
              ↕
┌─────────────────────────────────────┐
│         Adapter Layer               │
│  (Fiber, MongoDB, External APIs)    │
└─────────────────────────────────────┘
```

### Project Structure

```
.
├── cli/                    # CLI commands (Cobra)
│   ├── cmd/               # Command definitions
│   └── main.go            # Application entry point
├── configs/               # Configuration management
│   ├── const/             # Constants
│   └── env/               # Environment configuration
├── internal/
│   ├── adapter/           # External adapters
│   │   ├── calls/         # External API clients
│   │   ├── fiber/         # HTTP adapter (Fiber framework)
│   │   ├── mailer/        # Email adapter
│   │   └── mongo/         # MongoDB adapter
│   └── core/              # Core business logic
│       ├── domain/        # Domain models
│       ├── port/          # Port interfaces
│       └── service/       # Business services
├── pkg/                   # Shared packages
│   ├── errs/              # Error handling
│   ├── json/              # JSON utilities
│   ├── jwt/               # JWT token handling
│   ├── logs/              # Logging (Zap)
│   └── tracing/           # OpenTelemetry tracing
├── docker-compose.yml     # Docker services
└── Taskfile.yaml          # Task automation
```

## 🚀 Features

- **Hexagonal Architecture**: Clean separation of concerns
- **RESTful API**: Built with Fiber framework
- **Database**: MongoDB with OpenTelemetry instrumentation
- **Authentication**: JWT-based authentication
- **Distributed Tracing**: Jaeger integration via OpenTelemetry
- **Structured Logging**: Zap logger
- **External API Integration**: JSONPlaceholder example
- **Docker Support**: Docker Compose for local development

## 📋 Prerequisites

- Go 1.24.1 or higher
- Docker and Docker Compose
- MongoDB (via Docker)
- Redis (via Docker)
- Jaeger (via Docker)

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd hexagonal
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   
   Create `.env.dev` file in the root directory:
   ```env
   # Application
   APP_NAME=csat-servay
   APP_VERSION=1.0.0
   APP_ENV=development

   # Fiber Server
   FIBER_HOST=0.0.0.0
   FIBER_PORT=5000
   FIBER_ALLOW_ORIGINS=*
   FIBER_ALLOW_HEADERS=*
   FIBER_ALLOW_METHODS=GET,POST,PUT,DELETE,PATCH,OPTIONS
   FIBER_ALLOW_CREDENTIALS=true

   # MongoDB
   MONGO_USER=user_mongo
   MONGO_PASSWORD=JIX6dneXEvv
   MONGO_HOST=127.0.0.1
   MONGO_PORT=27018
   MONGO_BASE=csat

   # JWT
   JWT_HMAC_SECRET=your-secret-key-here
   JWT_ACCESS_EXP=3600
   JWT_REFRESH_EXP=86400

   # Tracing (Jaeger)
   TRACING_HOST=127.0.0.1
   TRACING_PORT=4317
   ```

4. **Start Docker services**
   ```bash
   docker-compose up -d
   ```

   This will start:
   - MongoDB on port `27018`
   - Redis on port `6378`
   - Jaeger UI on port `16686`

## 🎯 Usage

### Development

Start the development server using Taskfile:

```bash
task dev
```

Or manually:

```bash
go run cli/main.go start --zone dev
```

### Available Tasks

```bash
# Start development server
task dev

# Run tests
task test

# Run linter
task lint

# Generate Swagger documentation
task swag

# Build Docker image
task build

# Run Docker container
task run
```

### API Endpoints

Once the server is running, you can access:

- **Health Check**: `GET http://localhost:5000/v1/ping/`
- **JSONPlaceholder Example**: `GET http://localhost:5000/v1/ping/jsonplaceholder`

### Jaeger Tracing

View distributed traces in Jaeger UI:
- URL: http://localhost:16686

## 🧪 Testing

Run tests:

```bash
task test
```

Or manually:

```bash
go test -v ./...
```

## 📦 Dependencies

### Core Dependencies

- **Fiber v2**: Fast HTTP web framework
- **MongoDB Driver**: Official MongoDB Go driver
- **Cobra**: CLI framework
- **Zap**: Structured logging
- **OpenTelemetry**: Distributed tracing
- **JWT**: JWT token handling
- **Resty**: HTTP client for external APIs

See `go.mod` for complete dependency list.

## 🏛️ Architecture Details

### Hexagonal Architecture Layers

1. **Domain Layer** (`internal/core/domain/`)
   - Pure business logic
   - Domain models and entities

2. **Port Layer** (`internal/core/port/`)
   - Interfaces defining contracts
   - Inbound ports (use cases)
   - Outbound ports (repositories, external services)

3. **Service Layer** (`internal/core/service/`)
   - Business logic implementation
   - Orchestrates domain operations

4. **Adapter Layer** (`internal/adapter/`)
   - **Primary Adapters** (Inbound):
     - Fiber HTTP handlers
   - **Secondary Adapters** (Outbound):
     - MongoDB repositories
     - External API clients
     - Email service

### Dependency Flow

```
HTTP Request → Fiber Adapter → Service → Port Interface → Repository Adapter → MongoDB
```

## 🔧 Configuration

The application uses environment-based configuration. Configuration files are located in `configs/env/`:

- Environment files: `.env.{zone}` (e.g., `.env.dev`, `.env.prod`)
- Configuration model: `env.model.go`
- Configuration loader: `env.config.go`

## 📝 Code Style

- Follow Go standard formatting (`gofmt`)
- Use `golangci-lint` for linting
- Follow hexagonal architecture principles
- Keep business logic in the core layer
- Adapters should be thin and focused

## 🐳 Docker

### Build Image

```bash
task build
```

### Run Container

```bash
task run
```

### Docker Compose Services

- **MongoDB**: Database service
- **Redis**: Caching service
- **Jaeger**: Distributed tracing backend

## 📚 Additional Resources

- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)
- [Fiber Documentation](https://docs.gofiber.io/)
- [MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/current/)
- [OpenTelemetry Go](https://opentelemetry.io/docs/instrumentation/go/)

## 🤝 Contributing

1. Create a feature branch
2. Make your changes following the architecture principles
3. Write tests for new features
4. Run linter and tests
5. Submit a pull request

## 📄 License

[Add your license here]

## 👥 Authors

[Add author information here]

---

**Note**: This is a template project demonstrating hexagonal architecture in Go. Customize it according to your specific requirements.

