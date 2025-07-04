# go-modular-sample

A sample Go project, with source code organized using modular architecture and Domain-Driven Design.

## How to use

Use docker to build and run this project.

## Package used

- [Uber Go FX](https://github.com/uber-go/fx): dependency injection system
- [pq](https://github.com/lib/pq): PostgreSQL driver
- [go-redis](https://github.com/redis/go-redis): Redis driver

## How this project works

### Setup phase

```mermaid
flowchart LR
  subgraph B[Setup DI Container]
    direction LR
    C[Setup general dependencies]
    D[Setup modules & its dependencies]
    E[Setup lifecycle]
  end

  A[cmd/api/main.go] --> B
  B --> F[Invoke and run app]
```

### Execute phase (handle a HTTP request)

```mermaid
flowchart LR
  subgraph C[Module]
    direction TB
    D[Present layer] --> E[Application layer]
    E --> D
    E --> F[Domain layer]
    F --> E
    F --> |Use DI container| G[Infrastructure layer]
  end

  A[cmd/api/main.go] --> B[HTTP server]
  B --> C
```

## Source code structure

```
.
├── cmd/                        # Entry points for applications
│   └── api/                    # (Do not change) API
│   └── cli/                    # (Do not change) CLI
│   └── web/                    # (Do not change) Web
├── config/
│   └── config.go               # List of module configs & deps (between interface & implement in internal/shared)
│   └── setup.go                # (Do not change) Setup DI container & fx life cycle
├── internal/
│   ├── module/                 # Modules
│       ├── healthcheck/        # A sample module
│           ├── app/            # Application layer
│               ├── usecase/    # Usecase
│           ├── domain/         # Domain layer
│           ├── infra/          # Infrastructure layer (implement of domain/usecase)
│           ├── present/        # Presentation layer (API/CLI/Web)
│           ├── config.go       # (Do not change) Config file
│           ├── deps.go         # List of usecases / dependencies / routes
│   ├── shared/                 # Shared, used among modules
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
