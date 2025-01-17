# Snippetbox

Snippetbox is a web application built using Go, designed for sharing and managing text snippets. This project is structured to progressively cover the development of a professional, secure, and maintainable Go-based web application.

---

## Features

- **Dynamic Routing**: Customizable routes for various application functionalities.
- **Secure Database Interaction**: Uses MySQL with connection pooling and secure queries.
- **Template Rendering**: Supports dynamic HTML rendering with Go's `html/template` package.
- **Middleware Integration**: Includes security headers, logging, and panic recovery.
- **User Authentication**: Secure signup, login, and logout flows with session management.
- **HTTPS Support**: Configures self-signed TLS certificates and connection timeouts.
- **Testing Suite**: Comprehensive testing including unit, integration, and end-to-end tests.

---

## Project Structure

```
Snippetbox/
├── cmd/
│   └── web/
│       ├── main.go            # Entry point for the application
│       ├── handlers.go        # HTTP handlers
│       ├── middleware.go      # Middleware functions
│       ├── routes.go          # Route definitions
│       ├── config.go          # Application configuration
│       └── logger.go          # Logging setup
├── internal/
│   ├── models/
│   │   ├── models.go          # Database model definitions
│   │   ├── users.go           # User model and queries
│   │   └── snippets.go        # Snippet model and queries
│   ├── validator/
│   │   └── validator.go       # Input validation helpers
│   └── context/
│       └── context.go         # Context utilities for request handling
├── migrations/
│   └── 001_create_tables.sql  # SQL migration scripts
├── ui/
│   ├── html/
│   │   ├── base.tmpl          # Base HTML template
│   │   ├── partials/
│   │   │   ├── nav.tmpl       # Navigation partial
│   │   └── pages/
│   │       ├── home.tmpl      # Home page template
│   │       ├── snippet.tmpl   # Snippet view template
│   │       └── user.tmpl      # User-related templates
│   ├── static/
│   │   ├── css/
│   │   │   └── main.css       # CSS styles
│   │   ├── js/
│   │   │   └── main.js        # JavaScript scripts
│   │   └── img/               # Static images
├── go.mod                     # Go module file
├── go.sum                     # Module checksum file
└── README.md                  # Project README
```

---

## Chapters and Features

### 1. **Introduction**
- Overview of the project.
- Tools and prerequisites.

### 2. **Foundations**
- **Project Setup**: Initializing a Go module and setting up directories.
- **Web Server**: Creating a simple HTTP server.
- **Routing**: Handling requests using custom routes.
- **Templates**: Building reusable HTML templates with inheritance.
- **Static Files**: Serving CSS, images, and JavaScript.

### 3. **Configuration and Error Handling**
- Centralized management of configuration settings.
- Implementing leveled logging and structured error handling.
- Dependency injection for flexibility.

### 4. **Database-Driven Responses**
- Setting up MySQL and database drivers.
- Designing and interacting with database models.
- Handling transactions and multiple queries.

### 5. **Dynamic HTML Templates**
- Rendering dynamic data.
- Implementing template caching for performance.
- Custom template functions and runtime error handling.

### 6. **Middleware**
- Security headers: CSP, HSTS, and more.
- Request logging and panic recovery mechanisms.
- Building composable middleware chains.

### 7. **Advanced Routing**
- Selecting and integrating third-party routers.
- Creating clean URLs and method-based routing.

### 8. **Processing Forms**
- HTML form creation and data validation.
- Handling errors and repopulating fields dynamically.
- Writing reusable form validation helpers.

### 9. **Stateful HTTP**
- Implementing session management using cookies.
- Storing and retrieving session data securely.

### 10. **Security Improvements**
- Generating TLS certificates and configuring HTTPS.
- Setting secure timeouts for connections.

### 11. **User Authentication**
- Secure user registration and password encryption.
- Login and logout functionality.
- Role-based authorization and CSRF protection.

### 12. **Using Request Context**
- Leveraging context for authentication and authorization.

### 13. **Optional Go Features**
- Embedding files directly into the binary.
- Utilizing generics for cleaner and reusable code.

### 14. **Testing**
- Unit tests for isolated functionality.
- Testing HTTP handlers and middleware.
- End-to-end and integration tests for workflows.
- Mocking dependencies for robust testing.

### 15. **Conclusion**
- Summary of learnings and final thoughts.

### 16. **Guided Exercises**
- Additional tasks to extend the application:
  - Adding new pages like "About" or "Account".
  - Debugging and enhancing features.
  - Implementing user-specific functionalities like password changes.

---

## Prerequisites

- **Go**: Version 1.18 or higher.
- **MySQL**: For database operations.
- **curl**: For testing HTTP requests.
- **Browser**: For accessing the application.

---

## Getting Started

### 1. Clone the Repository
```bash
$ git clone https://github.com/rakesh-garlapati2908/snippetbox.git
$ cd Snippetbox
```

### 2. Install Dependencies
```bash
$ go mod tidy
```

### 3. Set Up MySQL Database
- Create a database and update the configuration in the project.

### 4. Run the Application
```bash
$ go run ./cmd/web
```

### 5. Access the Application
Open your browser and navigate to:
```
http://localhost:4000
```

---

## Testing

Run the test suite:
```bash
$ go test ./...
```

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.


