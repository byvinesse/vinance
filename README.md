# Vinance

Vinance is a personal finance management platform that helps users track spending, manage budgets, and monitor their financial health. The monorepo consolidates all client, server, and infrastructure code in one place.

---

## Apps

### `apps/api` — Backend API (Go)
REST API built with [Echo](https://echo.labstack.com/) and PostgreSQL. Handles authentication, transaction records, budgets, and all core business logic.

- **Language:** Go 1.19+
- **Key deps:** Echo, sqlx, squirrel, zerolog, golang-jwt
- **Entry point:** `main.go` · **Docs:** `apps/api/docs/`

### `apps/web` — Web App (Next.js)
Customer-facing web frontend built with Next.js 15 (App Router) and Tailwind CSS.

- **Language:** TypeScript
- **Key deps:** Next.js, Radix UI, shadcn/ui, Recharts
- **Dev:** `cd apps/web && npm run dev` (runs on port 4000)

### `apps/ios` — iOS App (Swift / SwiftUI)
Native iOS application built with SwiftUI. Mirrors the feature set of the web app for mobile users.

- **Language:** Swift
- **Project file:** `apps/ios/Vinance.xcodeproj`
- **Min target:** see Xcode project settings

---

## Database

### `database/migrations` — DB Migrations (Go)
Schema migration runner that applies versioned SQL scripts to the PostgreSQL database.

- **Language:** Go
- **Entry point:** `database/migrations/main.go`
- **Scripts:** `database/migrations/scripts/`

---

## Getting Started

Each app has its own setup instructions — refer to the README (where available) inside each subdirectory:

| App | README |
|-----|--------|
| API | `apps/api/README_TESTS.md` |
| Web | `apps/web/README.md` |
| DB Migrations | `database/migrations/README.md` |
