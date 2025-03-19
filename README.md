myapp/
│── cmd/            # Entry points for applications (main packages)
│   ├── myapp/      # Main application binary
│   │   ├── main.go # Application entry point
│   ├── another/    # Another binary (if applicable)
│       ├── main.go 
│
├── pkg/            # Reusable code (can be imported by other projects)
│   ├── util/       # Utility functions
│   ├── logger/     # Logging setup
│   ├── config/     # Configuration handling
│
├── internal/       # Private application-specific packages (not for public use)
│   ├── database/   # Database logic
│   ├── service/    # Business logic
│   ├── api/        # Handlers and API logic
│
├── api/            # API specifications (e.g., OpenAPI, GraphQL schemas)
│
├── configs/        # Configuration files (YAML, JSON, ENV files)
│
├── migrations/     # Database migration files
│
├── scripts/        # Helper scripts (e.g., for CI/CD or automation)
│
├── test/           # Additional test data or integration tests
│
├── web/            # Static files, templates, or frontend assets (if applicable)
│
├── go.mod          # Go module definition
├── go.sum          # Go module dependencies
├── Makefile        # Build and automation commands
├── README.md       # Project documentation



Breakdown:
cmd/

Contains entry points for different binaries.
cmd/myapp/main.go would be the main entry point for the primary application.
pkg/

Holds reusable Go packages that can be used by other projects.
It’s public but should not contain application-specific logic.
internal/

Contains application-specific code that should not be imported by external applications.
Best for business logic, services, database access, and APIs.
api/

Stores API-related definitions, such as OpenAPI (Swagger) specs or GraphQL schemas.
configs/

Stores configuration files like JSON, YAML, or ENV files.
migrations/

Database migration files if using a tool like migrate or golang-migrate.
scripts/

Contains scripts for automation, CI/CD, database setup, or deployment.
test/

Stores integration tests, fixtures, or test configurations.
web/

Holds frontend files if the project serves web pages.
go.mod & go.sum

Defines dependencies and versions.
Makefile

Contains helpful commands for building, testing, and running the application.
README.md

Project documentation.
Best Practices:

Use cmd/ for binaries instead of placing all logic in main.go.

Use internal/ for code that shouldn’t be accessed outside your project.

Keep business logic separate from HTTP handlers (use service/).

Follow idiomatic Go naming conventions.


# Command Line Tool for splitting files to manage large json files  created using Cobra 

