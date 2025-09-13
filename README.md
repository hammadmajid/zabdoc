# zabdoc

zabdoc is a web application for SZABIST students to generate assignment cover sheets, lab tasks and lab projects in PDF format. It uses a hybrid architecture with a Go backend API and a SvelteKit frontend deployed on Cloudflare.

## Architecture

### Backend (Go API)

- **Language:** Go 1.24.6+, using [chi router](https://github.com/go-chi/chi)
- **Deployment:** Heroku Container Stack (see [Heroku Container Registry docs](https://devcenter.heroku.com/articles/container-registry-and-runtime))
- **App Logic:** `internal/app/app.go` (dependency injection, service wiring)
- **API/HTTP Handlers:** `internal/api/http/handler.go` (PDF generation, health endpoints)
- **Routing:** `internal/router/router.go` (sets up API endpoints with CORS)
- **Templates:** `internal/templates/` (PDF template generation)
- **DTOs:** `internal/api/dto/dto.go` (form data structure)
- **Middleware:** `internal/middleware/` (CORS, logging)
- **PDF Generation:** Uses [chromedp](https://github.com/chromedp/chromedp) for headless Chrome rendering

### Frontend (SvelteKit)

- **Framework:** SvelteKit v2+ with TypeScript ([docs](https://svelte.dev))
- **Deployment:** Cloudflare Workers ([deployment guide](https://developers.cloudflare.com/workers/))
- **UI Components:** [Shadcn Svelte](https://shadcn-svelte.com/) components in `worker/src/lib/components/`
- **Routes:** `worker/src/routes/`

## Project Structure

- **Go Backend:** Follows [Go Project Layout](https://github.com/golang-standards/project-layout)
- **Frontend:** Follows [SvelteKit project structure](https://kit.svelte.dev/docs/project-structure)

## Setup & Usage

### Prerequisites

- Go 1.24.6+
- Node.js 20+ and pnpm ([install guide](https://pnpm.io/installation))
- Docker (optional, for headless Chrome environments)

### Backend Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/hammadmajid/zabdoc.git
   cd zabdoc
   ```

2. Install Go dependencies:

   ```sh
   go mod tidy
   ```

3. Set environment variable

   ```sh
   export PORT=8080
   ```

4. Run the backend API:

   ```sh
   go run .
   ```

   The API will be available at `http://localhost:8080`

### Frontend Setup

1. Navigate to the worker directory:

   ```sh
   cd worker
   ```

2. Install dependencies:

   ```sh
   pnpm install
   ```

3. Start the development server:

   ```sh
   pnpm dev
   ```

   The frontend will be available at `http://zabdoc.localhost:5173`

### Building for Production

#### Backend

```sh
go build .
```

#### Frontend

```sh
cd worker
pnpm build
```

## Usage

- Open the SvelteKit frontend in your browser (`http://zabdoc.localhost:5173` in development)
- Fill out the form with student details and optional content
- Submit to generate and download the PDF via the Go backend API

## API Endpoints

### Backend API (Go)

- `/health` - GET, Health check (returns status)
- `/generate` - POST, generates PDF from form data

### Frontend Routes (SvelteKit)

- `/` - Main form page
- `/about` - About page

## Contributing

- Follow Go idioms: dependency injection, error handling, structured logging
- Add features via `internal/services`, update templates as needed
- PRs welcome! See code comments for guidance
- Use [staticcheck](https://staticcheck.io/) for linting Go code
- Use [pnpm](https://pnpm.io/) for frontend package management

## License

Project is licensed under [GNU General Public License v3](https://www.gnu.org/licenses/gpl-3.0.html).

---
For questions or support, open an issue.
