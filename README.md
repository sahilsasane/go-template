golang template
```
go/
├── cmd/
│   └── api/                   # Application entry point & HTTP handlers
│       ├── main.go            # Main application entry point
│       ├── server.go          # HTTP server setup and graceful shutdown
│       ├── routes.go          # API routing definitions
│       ├── users.go           # User-specific handlers
│       ├── tokens.go          # Authentication token handlers
│       └── middleware.go      # HTTP middleware components
└── internal/                  # Internal packages
    ├── data/                  # Data models and database interactions
    │   └── tokens.go          # Token model and DB operations
    ├── jsonlog/               # JSON logging package
    ├── validator/             # Request validation
    └── utils/                 # utils
```

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/sahilsasane/cautious-memory.git
    cd cautious-memory
    ```

2. Set up environment variables:
    ```sh
    cp .env.example .env
    # Edit .env file with your configuration
    ```

3. Initialize go module:
    ```sh
    go mod init <project_name>.<username>.<net>
    ```

## Makefile

### Commands

- **help**: Print this help message.
- **run/api**: Run the `cmd/api` application.
- **db/psql**: Connect to the database using `psql`.
- **db/migrations/new name=$1**: Create a new database migration.
- **db/migrations/up**: Apply all up database migrations.
- **db/check**: Check the database connection and print the current DSN.
- **audit**: Tidy dependencies, format, vet, and test all code.
- **vendor**: Tidy, verify, and vendor dependencies.
- **build/api**: Build the `cmd/api` application.