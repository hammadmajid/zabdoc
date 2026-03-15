# zabdoc

zabdoc is a full-stack application for SZABIST students to generate assignment cover sheets, lab tasks, and scrape ZabDesk course data. Built with Go (backend) and SvelteKit (frontend), deployed on Heroku and Cloudflare respectively.

## Project Structure

```sh
zabdoc/
├── cmd/                          # CLI entry points
├── internal/
│   ├── app/                      # Application initialization & DI container
│   ├── api/
│   │   ├── http/                 # HTTP handlers for all endpoints
│   │   └── dto/                  # Request/response DTOs
│   ├── services/                 # Business logic layer
│   ├── templates/                # PDF templates (HTML partials)
│   ├── middleware/               # HTTP middleware (CORS, logging, etc.)
│   └── router/                   # Route definitions
├── worker/                       # SvelteKit frontend (Cloudflare Workers)
│   ├── src/
│   │   ├── routes/              # SvelteKit file-based routing
│   │   ├── lib/
│   │   │   ├── components/      # Reusable Svelte components (shadcn-svelte)
│   │   │   ├── stores/          # Svelte reactive state (wizard, form stores)
│   │   │   └── utils/           # Helper functions
│   │   └── app.html             # Root HTML template
│   ├── wrangler.toml            # Cloudflare Workers config
│   └── package.json
├── go.mod / go.sum              # Go dependencies
└── README.md
```

## Technical Architecture

### Backend Stack (Go)

- **Runtime:** Go 1.24.6+
- **HTTP Router:** [chi](https://github.com/go-chi/chi) - lightweight, composable router with middleware support
- **PDF Generation:** [chromedp](https://github.com/chromedp/chromedp) - headless Chrome automation for rendering HTML to PDF
- **Deployment:** Heroku Container Stack (Docker-based)
- **Key Flow:** HTTP POST → Form DTO validation → PDF generation via headless browser → Binary response

**Core Files:**

- `internal/app/app.go` - Application setup, dependency injection container
- `internal/api/http/handler.go` - HTTP handler implementations (health check, PDF generation)
- `internal/router/router.go` - Route registration, middleware setup, CORS configuration
- `internal/api/dto/dto.go` - Form data structures (request validation)
- `internal/middleware/` - Cross-cutting concerns (CORS, request logging)
- `internal/templates/` - HTML template partials used for PDF rendering

### Frontend Stack (SvelteKit + Svelte 5)

- **Framework:** SvelteKit v2+ with TypeScript, Svelte 5 runes
- **Deployment:** Cloudflare Workers with `@sveltejs/adapter-cloudflare`
- **UI Library:** [shadcn-svelte](https://shadcn-svelte.com/) - unstyled, accessible component primitives
- **State Management:** Svelte stores (reactive `$state()` runes) in `lib/stores/`
- **Styling:** Tailwind CSS with custom neo-brutalism design tokens

**Key Features:**

- Multi-step form wizard for document generation (assignment/lab task/lab project)
- ZabDesk data scraper - authentication & fetching attendance/marks
- Local storage persistence for user preferences
- Responsive design with neo-brutalism aesthetic

## Prerequisites

### Required

- Go 1.24.6 or later
- Node.js 20+ (for pnpm compatibility)
- pnpm 9+ ([installation guide](https://pnpm.io/installation))

### Optional

- Docker (recommended for local headless Chrome testing)
- Heroku CLI (for backend deployment)
- Wrangler CLI (for Cloudflare Workers deployment)

## Development Setup

### Backend Setup

```sh
# Clone and navigate to project
git clone https://github.com/hammadmajid/zabdoc.git
cd zabdoc

# Install Go dependencies
go mod tidy

# Set port (defaults to 8080)
export PORT=8080

# Run development server
go run .
```

Backend will be available at `http://localhost:8080`

**Available endpoints:**

- `GET /health` - Health check endpoint
- `POST /generate` - Generate PDF from form data (expects multipart/form-data)
- `POST /scrap` - Scrape data from ZabDesk (expects username and password)

### Frontend Setup

```sh
# Navigate to frontend directory
cd worker

# Install dependencies using pnpm
pnpm install

# Start development server
pnpm dev
```

Frontend will be available at `http://zabdoc.localhost:5173`

**Key pnpm scripts:**

- `pnpm dev` - Development server with HMR
- `pnpm build` - Build for production (Cloudflare Workers)

## Building for Production

### Backend Build

```sh
# Build Go binary (outputs: zabdoc or zabdoc.exe)
go build -o zabdoc .

# Or use Docker for consistent environment
docker build -t zabdoc:latest .
```

Environment variables for production:

- `PORT` - HTTP port (default: 8080)
- `API_URL` - Frontend API endpoint (for CORS)

### Frontend Build

```sh
cd worker
pnpm build
```

Output is ready for Cloudflare Workers deployment. Configuration in `wrangler.toml`.

## Deployment

### Backend (Heroku)

Uses Heroku Container Stack with Docker. See [Heroku Container Registry docs](https://devcenter.heroku.com/articles/container-registry-and-runtime).

```sh
heroku login
heroku create zabdoc-api
heroku container:push web
heroku container:release web
```

### Frontend (Cloudflare Workers)

Uses `@sveltejs/adapter-cloudflare` with wrangler.

```sh
cd worker
pnpm build
wrangler deploy
```

## Contributing

### Code Style & Quality

- **Go:** Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
  - Use [staticcheck](https://staticcheck.io/) for linting: `staticcheck ./...`
  - Prefer explicit error handling over panic
  - Use dependency injection pattern (see `internal/app/app.go`)

- **TypeScript/Svelte:** Use strict TypeScript mode
  - Run `pnpm lint` before committing
  - Follow [Svelte best practices](https://svelte.dev/docs/svelte)

### Contribution Workflow

1. **Start with an issue**: Check [GitHub Issues](https://github.com/hammadmajid/zabdoc/issues)
   - Pick an existing issue or open a discussion
   - Comment on the issue before starting work to avoid duplicates

2. **Create a feature branch**:

   ```sh
   git checkout -b feature/your-feature-name
   ```

3. **Make changes following the code style above**

4. **Test locally**:
   - Backend: `go run .` + manual testing
   - Frontend: `pnpm dev` + browser testing
   - Both: Verify integration works

5. **Commit with clear messages**:

   ```sh
   git commit -m "add feature X"
   ```

6. **Push and open PR**:

   ```sh
   git push origin feature/your-feature-name
   ```

   - Reference the issue: "Closes #123"
   - Provide context on changes
   - Run `go fmt ./...` and `pnpm format` before submitting

### Adding New Features

- **Backend:** Add handler in `internal/api/http/handler.go`, register in `internal/router/router.go`
- **Frontend:** Use file-based routing in `worker/src/routes/`, add components in `worker/src/lib/components/`
- **State:** Use Svelte stores in `worker/src/lib/stores/` for shared state

## Troubleshooting

### Backend

- **Port already in use:** Change PORT env var: `export PORT=3000 && go run .`
- **Chromedp issues:** Ensure Chrome/Chromium is available or use Docker
- **CORS errors:** Check `internal/middleware/cors.go` and API_URL env var

### Frontend

- **Module not found:** Run `pnpm install` in `worker/` directory
- **Build fails:** Clear `.svelte-kit` and rebuild: `rm -rf .svelte-kit && pnpm build`
- **HMR not working:** Check `vite.config.ts` configuration

## License

Licensed under [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.html). See LICENSE file for details.

---

**Questions?** Open an [issue](https://github.com/hammadmajid/zabdoc/issues) or start a [discussion](https://github.com/hammadmajid/zabdoc/discussions).
