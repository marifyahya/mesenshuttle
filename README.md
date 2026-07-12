# MesenShuttle

A web-based shuttle bus booking system built with modern technologies using a simplified microservices architecture.

## Tech Stack

- **Frontend:** Nuxt.js 3 (Vue 3)
- **Backend:** Golang (Gin Framework + GORM)
- **Database:** MySQL 8
- **Cache / Message Broker:** Redis
- **Infrastructure:** Docker & Docker Compose

## Prerequisites

Ensure you have the following software installed on your machine:
- [Docker](https://www.docker.com/products/docker-desktop) and Docker Compose
- [Node.js](https://nodejs.org/) (minimum v20.x, recommended if you want to develop the frontend locally without Docker)
- [Golang](https://go.dev/) (minimum v1.22, recommended if you want to develop the backend locally without Docker)

---

## 🛠 Running the Project (Development Mode)

In development mode, the source code is mounted directly into the Docker containers. You can make code changes locally and the system will auto-reload (using `air` for the backend and Vite/Nuxt hot-module replacement for the frontend).

**Steps:**

1. **Setup Environment Variables**
   Copy the example Docker override file:
   ```bash
   cp docker-compose.override.yml.example docker-compose.override.yml
   ```
   *Optional: Edit `docker-compose.override.yml` to change default mapped ports or local database credentials.*

2. **Start Docker Compose**
   Run the following command in the project root terminal:
   ```bash
   docker compose up -d --build
   ```

3. **Access the Application**
   - Frontend: [http://localhost:3000](http://localhost:3000)
   - Backend API: [http://localhost:8081](http://localhost:8081)
   - MySQL: `localhost:3306` (Use DBeaver / TablePlus to access your local database)
   - Redis: `localhost:6380`

**Development Note:**
Whenever you install new packages via `npm install` or add Go modules via `go get`, you do not need to rebuild the containers since the volumes are mounted.

---

## 🚀 Running the Project (Production Mode)

In production mode, the application (both Frontend and Backend) is compiled into a lightweight binary/static build and runs without relying on local source code. This ensures a minimal footprint and maximum security.

**Steps:**

1. Ensure you are in the project root directory.
2. Run Docker Compose without the development override file:
   ```bash
   docker compose -f docker-compose.yml up -d --build
   ```

---

## Directory Structure

- `/frontend` - Nuxt.js UI source code
- `/backend` - Golang REST API source code
- `docker-compose.yml` - Main Docker Compose configuration (uses production images)
- `docker-compose.override.yml` - Additional configuration for development mode (hot-reloading, exposed local ports, etc.)
