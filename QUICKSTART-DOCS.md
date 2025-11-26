# Quick Start: Mintlify Documentation

## 🚀 Start Documentation Server

### Method 1: Using npm scripts (Recommended)

```bash
# Install dependencies (first time only)
npm install

# Start docs server
npm run docs:dev
```

### Method 2: Using npx (No installation)

```bash
npx mintlify dev
```

### Method 3: Using scripts

**Windows:**
```bash
scripts\dev-docs.bat
```

**Linux/Mac:**
```bash
chmod +x scripts/dev-docs.sh
./scripts/dev-docs.sh
```

## 📍 Access Points

- **Documentation**: http://localhost:3000
- **API Server**: http://localhost:8080
- **OpenAPI Spec**: http://localhost:8080/openapi.yaml
- **Docs Redirect**: http://localhost:8080/docs → redirects to Mintlify

## 🎯 Features

✅ **Interactive API Playground** - Test endpoints directly in the browser  
✅ **Auto-generated from OpenAPI** - Always in sync with your API  
✅ **Beautiful UI** - Modern, responsive design  
✅ **Search** - Find endpoints quickly  
✅ **Code Examples** - Copy-paste ready snippets  

## 📝 Development Workflow

1. **Start API Server:**
   ```bash
   go run cmd/main.go
   ```

2. **Start Docs Server** (in another terminal):
   ```bash
   npm run docs:dev
   ```

3. **Edit Documentation:**
   - Edit files in `docs/` folder
   - Changes auto-reload in browser
   - OpenAPI spec updates automatically

## 🔧 Configuration

- **Mintlify Config**: `docs/mint.json`
- **OpenAPI Spec**: `openapi.yaml`
- **Documentation Files**: `docs/**/*.mdx`

## 📚 Documentation Structure

```
docs/
├── introduction.mdx          # Welcome page
├── quickstart.mdx            # Getting started guide
├── mint.json                 # Mintlify configuration
└── api-reference/
    ├── introduction.mdx      # API overview
    ├── health/
    │   ├── health-check.mdx
    │   └── root.mdx
    ├── products/
    │   ├── create-product.mdx
    │   └── get-product.mdx
    └── orders/
        └── get-order.mdx
```

## 🎨 Customization

Edit `docs/mint.json` to customize:
- Colors and branding
- Navigation structure
- Social links
- API playground settings

## 🚢 Production Build

```bash
npm run docs:build
```

This generates static files in `.mint/` directory that can be deployed to:
- Vercel
- Netlify
- GitHub Pages
- Any static hosting

## 💡 Tips

- The OpenAPI spec is automatically served at `/openapi.yaml`
- Mintlify reads the OpenAPI spec and generates interactive docs
- All API endpoints are testable directly from the docs
- Use `<RequestExample>` and `<ResponseExample>` components in MDX files

