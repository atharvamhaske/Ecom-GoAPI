# E-commerce API

A RESTful API for e-commerce operations built with Go, featuring product management and order processing capabilities.

## Overview

This project provides a backend API for managing an e-commerce platform with support for products, orders, and order items. Built with modern Go practices, it uses PostgreSQL for data persistence and includes comprehensive API documentation.

## Features

- Product management (create, list, retrieve)
- Order processing with line items
- PostgreSQL database with migrations
- Type-safe SQL queries using SQLC
- RESTful API design
- Health check endpoints
- Graceful server shutdown
- OpenAPI 3.0 specification
- Interactive documentation with Mintlify

## Tech Stack

- **Language**: Go 1.25.4
- **Router**: Chi v5
- **Database**: PostgreSQL 16
- **Database Driver**: pgx/v5
- **Query Builder**: SQLC
- **Migration Tool**: Goose
- **Documentation**: Mintlify, OpenAPI 3.0
- **Containerization**: Docker Compose

## Project Structure

```
.
├── cmd/                    # Application entry points
│   ├── main.go            # Main application
│   └── api.go             # API server configuration
├── internal/              # Private application code
│   ├── adapters/
│   │   ├── migrations/    # Database migrations
│   │   └── sqlc/          # Generated SQLC code
│   ├── orders/            # Order domain logic
│   ├── products/          # Product domain logic
│   ├── env/               # Environment configuration
│   └── json/              # JSON utilities
├── docs/                  # Mintlify documentation
├── scripts/               # Utility scripts
├── docker-compose.yaml    # Docker services configuration
├── openapi.yaml           # OpenAPI specification
└── sqlc.yaml              # SQLC configuration
```

## Prerequisites

- Go 1.25.4 or higher
- Docker and Docker Compose
- PostgreSQL 16 (or use Docker Compose)
- Node.js (for documentation)

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/atharvamhaske/Ecom-GoAPI.git
cd ecom-api
```

### 2. Set up environment variables

Create a `.env` file based on the example:

```bash
cp .env.example .env
```

Configure your environment variables as needed.

### 3. Start the PostgreSQL database

```bash
docker-compose up -d
```

This will start a PostgreSQL instance on port 5432 with:
- User: `postgres`
- Password: `postgres`
- Database: `ecom`

### 4. Run database migrations

```bash
goose -dir internal/adapters/migrations postgres "user=postgres password=postgres dbname=ecom sslmode=disable" up
```

### 5. Build and run the application

```bash
go build -o ecom-api ./cmd
./ecom-api
```

The server will start on `http://localhost:8080`.

## API Endpoints

### Health Check

- `GET /` - Root endpoint
- `GET /health` - Health check endpoint

### Products

- `GET /api/v1/products` - List all products
- `POST /api/v1/products` - Create a new product
- `GET /api/v1/products/{id}` - Get product by ID

### Orders

- `POST /api/v1/orders` - Place a new order
- `GET /api/v1/orders/{id}` - Get order by ID

### Documentation

- `GET /openapi.yaml` - OpenAPI specification
- `GET /docs` - Redirects to Mintlify documentation (local)

## API Documentation

### View Interactive Documentation

Start the Mintlify documentation server:

```bash
npm install
npm run docs:dev
```

The documentation will be available at `http://localhost:3000`.

### OpenAPI Specification

The complete OpenAPI specification is available at `/openapi.yaml` when the server is running, or view the `openapi.yaml` file directly.

## Database Schema

### Products Table

```sql
CREATE TABLE products (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  price_in_centers INTEGER NOT NULL CHECK (price_in_centers >= 0),
  quantity INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

### Orders Table

```sql
CREATE TABLE orders (
  id BIGSERIAL PRIMARY KEY,
  customer_id BIGINT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

### Order Items Table

```sql
CREATE TABLE order_items (
  id BIGSERIAL PRIMARY KEY,
  order_id BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
  product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE RESTRICT,
  quantity INTEGER NOT NULL CHECK (quantity > 0),
  price_cents BIGINT NOT NULL CHECK (price_cents >= 0),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

## Development

### Running Tests

```bash
go test ./...
```

### Generating SQLC Code

After modifying SQL queries in `internal/adapters/sqlc/queries.sql`:

```bash
sqlc generate
```

### Creating New Migrations

```bash
goose -dir internal/adapters/migrations create migration_name sql
```

### Building for Production

```bash
go build -ldflags="-s -w" -o ecom-api ./cmd
```

## Configuration

The application uses the following configuration:

- **Server Address**: `:8080` (configurable via code)
- **Write Timeout**: 30 seconds
- **Read Timeout**: 10 seconds
- **Idle Timeout**: 1 minute
- **Request Timeout**: 60 seconds

## Middleware

The API includes the following middleware:

- Request ID (for tracing and rate limiting)
- Real IP extraction
- Request logging
- Panic recovery
- Request timeout (60 seconds)

## Graceful Shutdown

The server supports graceful shutdown with a 5-second timeout for completing in-flight requests. Send `SIGINT` or `SIGTERM` to trigger shutdown.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Contact

For API support or questions, please open an issue in the GitHub repository.

## Acknowledgments

- Built with [Chi](https://github.com/go-chi/chi) router
- Database queries generated with [SQLC](https://sqlc.dev/)
- Migrations managed with [Goose](https://github.com/pressly/goose)
- Documentation powered by [Mintlify](https://mintlify.com/)
