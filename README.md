# go-modular-sample

A sample Go project, with source code organized using modular architecture and Domain-Driven Design.

## Source code structure

```
.
├── cmd/                        # Entry points for applications (main packages)
│   └── api/                    # (Do not change) API
│   └── cli/                    # (Do not change) CLI
│   └── web/                    # (Do not change) Web
├── config/
│   └── config.go               # List of deps and module configs
│   └── setup.go                # (Do not change) Setup DI container & fx life cycle
├── internal/
│   ├── module/                 # Modules
│       ├── healthcheck/        # A sample module
│           ├── app/            # Application
│               ├── usecase/    # Usecase
│           ├── domain/         # Domain layer
│           ├── infra/          # Infrastructure layer (implement of domain/service/usecase)
│           ├── present/        # Presentation layer (API/CLI/Web)
│           ├── config.go       # (Do not change) Config file
│           ├── deps.go         # List of usecases / dependencies / routes
│   ├── shared/                 # Shared
│       ├── api/                # API / HTTP
│       ├── constant/           # Constant
│       ├── db/                 # Database
│       ├── module/             # Module config
├── test/                       # Integration test
├── air.toml                    # Air, used for hot-reloading
├── .dockerignore
├── .env.template
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```
