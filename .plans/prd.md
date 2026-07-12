# Product Requirements Document (PRD): Shuttle Booking Application

## 1. Executive Summary
A web-based application for shuttle ticket booking that allows users to book tickets quickly without needing to register an account (guest checkout). The system will use the email address as the unique user identifier. This application is equipped with an interactive seat selection feature, automated payment integration, and a robust concurrency system to prevent double-booking. There is also an Admin portal to manage operational data such as routes and schedules.

## 2. Tech Stack
*   **Backend:** Golang
*   **Frontend:** Nuxt.js 3 (Vue framework)
*   **Persistent Database:** MySQL
*   **In-Memory Database:** Redis
*   **Payment Gateway:** Xendit
*   **Notifications:** Email (via SMTP)

## 3. User Personas
1.  **Customer (Passenger):** Wants to book a shuttle ticket quickly, select a seat, and pay without the hassle of registration.
2.  **Admin:** Operational staff who manage routes, shuttle fleets, and departure schedules.

## 4. User Flows

### 4.1. Customer Flow (No Registration)
1.  **Search:** User searches for departure schedules based on Origin Pool and Destination Pool (from a list of routes defined by the Admin), along with the departure date.
2.  **Schedule Selection:** System displays a list of schedules along with the shuttle type (**Premium** or **Reguler**). User selects a suitable schedule.
3.  **Seat Selection (Seat Map):** User views the available seat layout/map and selects the desired seat.
4.  **Passenger Data Input:** User inputs their Email address (as ID) and basic personal data (Name, Phone Number).
5.  **Seat Locking:** System places a temporary lock on the selected seat using **Redis** so it cannot be selected by others.
6.  **Payment:** System generates an invoice via **Xendit** and redirects the user to the payment page.
7.  **Confirmation:** After Xendit sends a successful payment webhook, the backend will update the ticket status to `PAID`.
8.  **Notification:** E-Ticket and payment receipt are sent to the user's email.

### 4.2. Admin Flow
1.  **Route Management:** Admin can create, update, and delete routes (Example: Jakarta - Bandung).
2.  **Fleet & Shuttle Type Management:** Admin defines fleet types along with their capacities. Seat layouts are static: **Reguler (1-2 configuration)** and **Premium (1-1 configuration)**.
3.  **Schedule Management:** Admin creates departure schedules by pairing a Route, Fleet, departure time, and ticket price.

## 5. Technical Requirements & Concurrency (Race Condition)
*   **Redis Pool / Lua Scripting:** When multiple users try to select the same seat at the same time (race condition), the Golang backend will use Redis Lua Scripts. The operations of checking seat availability and locking (set with expiry) will be executed **atomically**. This guarantees absolutely that no two users will get the same seat.
*   **Lock Expiration:** If the user does not complete the Xendit payment within the **30-minute** time limit, the Redis key will automatically expire, and the seat will return to `Available` status.
*   **Webhook Handling:** The Golang endpoint for Xendit will be made as secure as possible (verifying token/signature) and can quickly process `PAID` or `EXPIRED` statuses.

## 6. Out of Scope (Phase 1)
To prevent feature creep and ensure a successful initial launch, the following features are explicitly out of scope for Phase 1:
1.  **Admin User Management (CRUD):** There will be no UI for creating, updating, or deleting Admin staff accounts. The primary admin account will be provisioned directly via a Database Seeder.
2.  **Track Booking Status Page:** There will be no public page for guests to track their booking status using a Booking Code. Users will rely entirely on email notifications (E-Ticket and Invoice) to know the status of their order.
