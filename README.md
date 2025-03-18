golang template

cautious-memory/
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
