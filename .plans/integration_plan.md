# Frontend-Backend Integration Plan (Epic)

This document tracks all tasks (User Stories) for integrating the Frontend UI (Nuxt.js) with the Backend API (Golang). Ensure the respective backend endpoints are completed in `backend_plan.md` before working on their integration here.

## Phase 1: Integration Environment Setup
- `[ ]` As a FE, I want to configure Axios / Nuxt `$fetch` base URL to point to Golang Backend API.
- `[ ]` As a FE, I want to setup global error handling (e.g., auto redirect to login if 401 Unauthorized).
- `[ ]` As a FE, I want to setup Pinia Store to manage Authentication state (Token JWT & User Profile).

## Phase 2: Admin Authentication
- `[ ]` As a FE, I want to integrate the Login Page (`/admin/login`) with `POST /api/admin/login`.
- `[ ]` As a FE, I want to save the received JWT token into Cookies/LocalStorage securely.
- `[ ]` As a FE, I want to protect all `/admin/*` pages using Nuxt Route Middleware (redirect to login if no token).

## Phase 3: Admin Master Data (CRUD)
- `[ ]` As a FE, I want to integrate Route Management page with `GET`, `POST`, `PUT`, `DELETE` `/api/admin/routes`.
- `[ ]` As a FE, I want to integrate Fleet Management page with `GET`, `POST`, `PUT`, `DELETE` `/api/admin/fleets`.
- `[ ]` As a FE, I want to integrate Schedule Management page with `GET`, `POST`, `PUT`, `DELETE` `/api/admin/schedules`.
- `[ ]` As a FE, I want to attach JWT Token in Authorization Header (Bearer) for all Master Data requests.

## Phase 4: Public Booking Flow (Customer)
- `[ ]` As a FE, I want to integrate the Home Search Form with `GET /api/schedules` to fetch available shuttles based on date, origin, and destination.
- `[ ]` As a FE, I want to integrate the Seat Selection Map with `GET /api/schedules/:id/seats` to display real-time locked/booked seats.
- `[ ]` As a FE, I want to integrate the Checkout Form with `POST /api/checkout` to lock the seat and get the Payment URL (Invoice).

## Phase 5: Payment & E-Ticket
- `[ ]` As a FE, I want to handle redirection to the Payment Gateway (Xendit URL).
- `[ ]` As a FE, I want to display the Payment Success/E-Ticket page by verifying the booking status from the backend.
