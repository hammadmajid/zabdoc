# zabdoc

zabdoc is a web application for SZABIST students to generate assignment cover sheets and lab tasks in PDF format. It uses a hybrid architecture with a Go backend API and a SvelteKit frontend deployed on Cloudflare.

## Architecture

### Backend (Go API)

- **Language:** Go 1.24.6, using [chi router](https://github.com/go-chi/chi)
- **App Logic:** `internal/app/app.go` (dependency injection, service wiring)
- **API/HTTP Handlers:** `internal/api/http/handler.go` (PDF generation, health endpoints)
- **Controllers:** `internal/controllers/` (business logic layer)
- **Routing:** `internal/router/router.go` (sets up API endpoints with CORS)
- **Templates:** `internal/templates/` (PDF template generation)
- **DTOs:** `internal/api/dto/dto.go` (form data structure)
- **Middleware:** `internal/middleware/` (CORS, logging)

### Frontend (SvelteKit)

- **Framework:** SvelteKit with TypeScript
- **UI Components:** Custom Svelte components in `worker/src/lib/components/`
- **Styling:** TailwindCSS with custom component library
- **Deployment:** Cloudflare Workers/Pages
- **Routes:** `worker/src/routes/` (pages and API proxy)
- **Form Handling:** Auto-form component with date picker and validation

## Setup & Usage

### Prerequisites

- Go 1.24.6+
- Node.js 18+ and pnpm
- [chromedp](https://github.com/chromedp/chromedp) dependencies (installed via `go mod tidy`)

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

3. Run the backend API:

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
pnpm deploy  # deploys to Cloudflare
```

## Usage

- Open the SvelteKit frontend in your browser (`http://zabdoc.localhost:5173` in development)
- Fill out the assignment form with student details
- Submit to generate and download a PDF cover sheet via the Go backend API

## API Endpoints

### Backend API (Go)

- `/health` - Health check (returns status)
- `/assignment` - POST, generates PDF from form data

### Frontend Routes (SvelteKit)

- `/` - Main assignment form page
- `/about` - About page
- `/assignment` - Assignment form (alternative route)
- `/lab-task` - Lab task form
- `/api/assignment` - Proxy endpoint that forwards to Go backend

## Contributing

- Follow Go idioms: dependency injection, error handling, structured logging
- Add features via `internal/services`, update templates as needed
- PRs welcome! See code comments for guidance

## License

Project is licensed under [GNU General Public License v3](https://www.gnu.org/licenses/gpl-3.0.html).

---

For questions or support, open an issue.
