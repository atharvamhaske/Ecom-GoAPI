#GO E-com API backend
A RESTful API for e-commerce operations built with Go, featuring product management and order processing capabilities.
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
├── docker-compose.yaml    # Docker services configuration
├── openapi.yaml           # OpenAPI specification
└── sqlc.yaml              # SQLC configuration
```

## API Endpoints List

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
## Middlewares

- Request ID (for tracing and rate limiting)
- Real IP extraction
- Request logging
- Panic recovery
- Request timeout (60 seconds)

## License

This project is licensed under the MIT License.

