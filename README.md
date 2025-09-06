# Zabcover

Zabcover is a web application for SZABIST students to generate assignment cover sheets in PDF format. Written in Go, it provides a simple, responsive web form and produces institution-branded, printable PDFs for assignment submissions.

## Architecture

### Backend
- **Language:** Go 1.24, using [chi router](https://github.com/go-chi/chi)
- **Entry Point:** `cmd/zabcover/main.go` (sets up server, routes, logging)
- **App Logic:** `internal/app/` (dependency injection, service wiring)
- **API/HTTP Handlers:** `internal/api/http/handler.go` (form, PDF, health)
- **Routing:** `internal/router/router.go` (sets up endpoints, static file serving)
- **Services:** `internal/services/assignment.go` (PDF generation via chromedp)
- **DTOs:** `internal/api/dto/dto.go` (form data structure)
- **Middleware:** `internal/middleware/logger.go` (request logging)

### Frontend
- **Templates:** `web/templates/index.html` (main form), `assignment.html` (PDF layout)
- **CSS:** `web/css/terminal.css` forked from [Terminal CSS](https://terminalcss.xyz/)
- **JavaScript:** `web/js/form.js`, `storage.js` (form logic, local storage)
- **Static Assets:** `web/static/` (logo, favicon)

## Setup & Usage

### Prerequisites
- Go 1.24+
- [chromedp](https://github.com/chromedp/chromedp) dependencies (installed via `go mod tidy`)

### Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/hammadmajid/zabcover.git
   cd zabcover
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

### Running the Server
1. Set the `PORT` environment variable (e.g. `export PORT=8080`)
2. Start the development server:
   ```sh
   go run ./cmd/zabcover
   ```
3. Access the app at `http://localhost:<PORT>`

### Building for Production
```sh
go build ./cmd/zabcover
```

## Usage
- Open the web UI in your browser
- Fill out the assignment form
- Submit to generate and download a PDF cover sheet

## Endpoints
- `/` - Main form page
- `/assignment` - POST, generates and downloads PDF
- `/health` - Health check (returns status)
- `/static/`, `/css/`, `/js/` - Static assets

## Contributing
- Follow Go idioms: dependency injection, error handling, structured logging
- Add features via `internal/services`, update templates as needed
- PRs welcome! See code comments for guidance

## License
Project is licensed under [GNU General Public License v3](https://www.gnu.org/licenses/gpl-3.0.html).

---

For questions or support, open an issue.
