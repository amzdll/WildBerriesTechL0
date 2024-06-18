# WildBerriesTech L0


## Run
...

### Prerequisites
- Docker
- Docker Compose
- Go (version 1.22.3)

## Go dependencies
- chi
- chi/render
- uber/fx
- uber/config
- zerolog
- go-redis
- pgx
- nats-streaming-server

## Project Structure

```
.
├── bin                         # Migration tool and binaries
├── cmd
│   └── publisher               # Publisher service for message streaming
├── config                      # Configuration files
├── internal
│   ├── api
│   │   └── order               # Order-related API handlers and CRUD operations
│   │       └── response        # API responses
│   ├── app
│   │   ├── api                 # API initialization and routes
│   │   ├── broker
│   │   │   └── nats-streaming  # NATS streaming initialization
│   │   ├── core                # Core functionality
│   │   ├── db                  # Database initialization
│   │   └── logger              # Logger initialization
│   ├── broker
│   │   └── nats-streaming
│   │       └── order           # Order-related message handling
│   │           └── message     # Message definitions
│   ├── config
│   │   ├── broker              # Broker configuration
│   │   └── db                  # Database configuration
│   ├── model                   # Model definitions
│   ├── repository
│   │   └── order               # Order repository for CRUD operations
│   └── service
│       └── order               # Order service for business logic
└── migrations                  # Database migration scripts
```
