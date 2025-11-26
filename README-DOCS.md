# API Documentation Setup

This project uses [Mintlify](https://mintlify.com) for beautiful, interactive API documentation.

## Quick Start

### Option 1: Using npm (Recommended)

1. Install dependencies:
```bash
npm install
```

2. Start the documentation server:
```bash
npm run docs:dev
```

The docs will be available at `http://localhost:3000`

### Option 2: Using npx (No installation needed)

```bash
npx mintlify dev
```

## Documentation Structure

- `docs/` - All documentation markdown files
- `docs/mint.json` - Mintlify configuration
- `openapi.yaml` - OpenAPI specification (served at `/openapi.yaml`)

## API Routes

- `/openapi.yaml` - OpenAPI specification endpoint
- `/docs` - Redirects to Mintlify docs (when running locally)

## Building for Production

```bash
npm run docs:build
```

This generates static files that can be deployed to any static hosting service.

## Development Workflow

1. Start the API server: `go run cmd/main.go`
2. Start the docs server: `npm run docs:dev`
3. Access docs at: `http://localhost:3000`
4. Access API at: `http://localhost:8080`

## Mintlify Features

- ✅ Interactive API playground
- ✅ Auto-generated from OpenAPI spec
- ✅ Beautiful, modern UI
- ✅ Search functionality
- ✅ Code examples
- ✅ Request/Response examples

