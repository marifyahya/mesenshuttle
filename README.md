# MesenShuttle 🚌

> A full-stack shuttle bus booking platform built to solve the real-world problem of concurrent seat reservations — without requiring user registration.

---

## The Problem

When multiple passengers attempt to book the last available seat simultaneously, a naive implementation allows both requests to pass the availability check before either write completes — resulting in double-booking.

MesenShuttle solves this using **Redis Lua Scripts** for atomic seat locking: the check-and-lock operation executes as a single indivisible unit, making it physically impossible for two users to acquire the same seat.

---

## Project Status

| Feature | Status |
|---|---|
| Admin Authentication (JWT) | ✅ Done |
| Route Management (GET, POST) | ✅ Done |
| Fleet & Schedule Management | 📋 Planned |
| Seat Map + Redis Atomic Locking | 📋 Planned |
| Xendit Payment Integration | 📋 Planned |
| Asynq Background Email Queue | 📋 Planned |

---

## Tech Stack

| Layer | Technology |
|---|---|
| **Frontend** | Nuxt.js 3 (Vue 3) |
| **Backend** | Golang — Gin Framework + GORM |
| **Database** | MySQL 8 |
| **Cache & Seat Locking** | Redis (Lua Scripts) |
| **Task Queue** | Asynq (Redis-backed) |
| **Payment Gateway** | Xendit |
| **API Spec** | OpenAPI 3.0 |
| **Infrastructure** | Docker & Docker Compose |

---

## System Architecture

```mermaid
graph TD
    Client[Frontend Nuxt.js] -->|HTTP REST API| Gin[Golang Gin Router]

    subgraph Backend System
        Gin -->|Lock Seats TTL 30m| Redis[(Redis In-Memory)]
        Gin -->|CRUD & Master Data| DB[(MySQL Database)]
        Gin -->|Enqueue Email Task| AsynqQueue[(Redis Task Queue)]
        AsynqWorker[Background Worker] -->|Fetch Tasks| AsynqQueue
    end

    Gin -->|Create Invoice| Xendit[Xendit Payment Gateway]
    Xendit -->|Webhook Notification| Gin
    AsynqWorker -->|Send ETicket with Auto Retry| SMTP[Email SMTP]
```

---

## Database Design

```mermaid
erDiagram
    ROUTES ||--o{ SCHEDULES : "has"
    FLEETS ||--o{ SCHEDULES : "has"
    SCHEDULES ||--o{ BOOKINGS : "has"
    BOOKINGS ||--|{ BOOKING_SEATS : "contains"

    ADMINS {
        uuid id PK
        varchar name
        varchar email "UNIQUE, INDEX"
        varchar password_hash
        datetime created_at
        datetime updated_at
    }
    ROUTES {
        uuid id PK
        varchar origin_city "INDEX"
        varchar origin_pool
        varchar destination_city "INDEX"
        varchar destination_pool
        datetime created_at
        datetime updated_at
    }
    FLEETS {
        uuid id PK
        varchar plate_number "UNIQUE"
        varchar type "Premium/Reguler"
        int total_seats
        bool active_status
        datetime created_at
        datetime updated_at
    }
    SCHEDULES {
        uuid id PK
        uuid route_id FK "INDEX"
        uuid fleet_id FK "INDEX"
        datetime departure_time
        decimal price
        varchar status
        datetime created_at
        datetime updated_at
    }
    BOOKINGS {
        uuid id PK
        uuid schedule_id FK "INDEX"
        varchar booking_code "UNIQUE"
        varchar email "INDEX"
        varchar name
        varchar phone
        tinyint status "0:PENDING 1:PAID 2:EXPIRED"
        int total_price
        varchar payment_url
        datetime created_at
        datetime updated_at
    }
    BOOKING_SEATS {
        uuid id PK
        uuid booking_id FK "INDEX"
        varchar seat_number
    }
```

---

## Checkout Flow — Concurrency Design

The most critical design in the system. The sequence below shows how a race condition between two simultaneous requests is handled atomically using a Redis Lua Script.

```mermaid
sequenceDiagram
    participant User as Frontend (Vue)
    participant API as Backend (Gin)
    participant Redis as Redis (Cache)
    participant DB as MySQL
    participant Xendit as Xendit Gateway

    User->>API: POST /api/checkout (schedule_id, seats, user_data)

    Note over API,Redis: Lua Script executes Check & Lock atomically
    API->>Redis: EVAL Lua Script (Check availability & Lock seats)

    alt Seats already booked or locked
        Redis-->>API: Error: Seat Unavailable
        API-->>User: 409 Conflict
    else Seats are available
        Redis-->>API: Success (Locked for 30 mins)
        API->>DB: Insert Booking (Status: PENDING)
        API->>DB: Insert Booking Seats
        API->>Xendit: Create Invoice (Amount, external_id=booking_code)
        Xendit-->>API: Return invoice_url
        API->>DB: Update Booking (payment_url = invoice_url)
        API-->>User: 200 OK (invoice_url, booking_code)
    end
```

**Expiry Strategy:**
- **User-facing**: 30-minute payment window communicated to the customer.
- **Xendit**: Invoice TTL set to exactly 30 minutes.
- **Asynq Delayed Task**: A background job is enqueued to expire the booking at the **35th minute** — a 5-minute buffer to ensure a late-arriving Xendit webhook isn't beaten by our own cleanup job.

---

## Key Technical Decisions

- **Redis Lua Script over distributed locks**: Lua scripts execute atomically inside Redis, avoiding the overhead of SETNX-based distributed locks while guaranteeing exactly-once seat acquisition.
- **Guest checkout (email as identifier)**: Removing registration friction was a deliberate product decision to maximize booking conversion. Email is used as the unique customer identifier.
- **Manual DI over Wire/Dig**: Kept dependency injection explicit and readable for a solo project. Container modules (`internal/container`) group related wiring per feature domain.
- **Asynq over RabbitMQ/Kafka**: Simpler operational profile — Asynq reuses the existing Redis instance and provides reliable retry semantics without a separate message broker.
- **`cmd/cli` for migrations**: AutoMigrate removed from the API boot path; a separate CLI binary (`go run ./cmd/cli -migrate`) runs migrations explicitly to prevent unsafe schema changes during rolling deployments.

---

## Backend Directory Structure

```
backend/
├── cmd/
│   ├── api/          # HTTP server entrypoint
│   └── cli/          # Migration & seeder CLI (-migrate, -seed)
├── internal/
│   ├── container/    # DI modules (AuthModule, RouteModule)
│   ├── controllers/  # HTTP handlers
│   ├── services/     # Business logic
│   ├── repositories/ # Data access layer
│   ├── dto/          # Request/Response DTOs (snake_case)
│   ├── middlewares/  # JWT auth, rate limiter, structured logger
│   └── models/       # GORM database models
└── pkg/
    ├── apperrors/    # Sentinel errors (ErrInvalidCredentials, etc.)
    ├── database/     # MySQL & Redis initializers
    └── utils/        # JWT, bcrypt, response helpers
```

---

## Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop) and Docker Compose
- [Node.js](https://nodejs.org/) v20+ (for local frontend development)
- [Golang](https://go.dev/) v1.22+ (for local backend development)

---

## 🛠 Running the Project (Development Mode)

Hot-reload enabled via `air` (backend) and Vite HMR (frontend).

```bash
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
cp docker-compose.override.yml.example docker-compose.override.yml

docker compose up -d --build
```

| Service | URL |
|---|---|
| Frontend | http://localhost:3000 |
| Backend API | http://localhost:8081 |
| MySQL | localhost:3306 |
| Redis | localhost:6380 |

**Running Locally Without Docker:**
```bash
# Backend
cd backend
go mod tidy
go run ./cmd/cli -migrate   # run migrations first
go run ./cmd/cli -seed      # seed admin account
go run ./cmd/api            # start the API server

# Frontend
cd frontend
npm install
npm run dev
```

---

## 🚀 Running the Project (Production Mode)

```bash
docker compose -f docker-compose.yml up -d --build
```

---

## 📚 API Documentation

Full OpenAPI 3.0 specification is available at [`backend/openapi.yaml`](backend/openapi.yaml).
Import into **Postman**, **Swagger UI**, or **Insomnia**.

---

## 🧪 Testing

Unit tests cover the business logic and HTTP handler layers using `testify` and mock interfaces.

```bash
cd backend
go test -v ./...
```

---

## Known Limitations & Roadmap

- **Migrations**: Using GORM `AutoMigrate` via CLI. Production-grade versioned migrations ([golang-migrate](https://github.com/golang-migrate/migrate)) are the next step.
- **Rate Limiting**: Currently IP + User-Agent based on the login endpoint. A JWT-scoped approach would be more robust for authenticated routes.
- **Tests**: Only unit tests with mocks exist today. Integration tests against a real test database are planned.
- **Phases 3–5**: Seat map, payment integration, and email queue are planned (see `.plans/backend_plan.md` for full execution roadmap).
