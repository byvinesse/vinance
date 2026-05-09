# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

Vinance is a personal finance application with three sub-apps in a monorepo:

- `apps/api` — Go REST API (Echo framework, PostgreSQL)
- `apps/web` — Next.js 15 web frontend (React 19, Tailwind CSS v4, TypeScript)
- `apps/ios` — SwiftUI iOS app
- `database/migrations` — standalone Go migration tool (separate module)

Each app is an independent module with its own go.mod / package.json — there is no root-level package manager.

## API (`apps/api`)

### Commands
```bash
# Run
make start          # go run main.go (requires .env)

# Test
make test           # go test ./...
make test-vc        # verbose + HTML coverage report
go test ./pkg/authenticator   # single package

# Lint
make lint           # golangci-lint run ./...
```

### Required `.env` variables
```
PORT=
DATABASE_URI=
AUTH_JWT_KEY=
AUTH_ACCESS_TOKEN_DURATION=   # e.g. 24h
```

### Architecture

Entry: `main.go` → `cmd/application` (wires dependencies) → `cmd/server` (registers Echo routes).

Layered architecture with strict dependency direction:

```
handler/server  →  pkg/service  →  repository/db
                                    ↓
                                  entity (plain structs, DB table constants)
```

- **`entity/`** — DB structs and domain constants (currency types, account types, record types, etc.). No business logic.
- **`model/`** — Request/response DTOs (`requests.go`, `responses.go`).
- **`repository/`** — Interfaces in `repository.go`; PostgreSQL implementations in `repository/db/` using `sqlx` + `squirrel`.
- **`pkg/service/`** — Business logic; interfaces in `interfaces.go`.
- **`handler/server/`** — Echo handlers; one file per endpoint.
- **`pkg/errors/`** — Typed error types (`BadRequestError`, `UnauthorizedError`, etc.) with a central Echo error middleware.
- **`pkg/response/`** — `Ok[T]` and `OkCreated[T]` generic helpers that wrap responses in `{code, status, data}`.
- **`pkg/authenticator/`** — JWT generation and parsing (HS256).
- **`pkg/middleware/`** — `Authentication` middleware that validates JWT and sets `user_email` on the context.
- **`config/`** — Loads env vars via `godotenv`; panics on missing values at startup.

Routes follow the pattern `/resource/v1/_action` (versioned, verbs for non-CRUD actions). All non-auth routes use the `withAuth` middleware.

Records use cursor-based pagination (keyset on `recorded_at DESC, id DESC`).

### Testing conventions
Uses `testify/assert` and `testify/mock`. Mocks are created per test file using interface-based dependency injection. Repository tests use `DATA-DOG/go-sqlmock`.

## Web (`apps/web`)

### Commands
```bash
cd apps/web
npm run dev          # Next.js on port 4000 (Turbopack)
npm run build
npm run lint
npm test             # Jest
npm run test:watch
npm run test:coverage
```

### Architecture

Next.js App Router. Sidebar layout in `src/app/layout.tsx` wraps all pages with `SidebarProvider` + `AppSidebar`.

- **`src/app/`** — App Router pages and global layout.
- **`src/components/ui/`** — Radix UI primitives (button, sidebar, tooltip, etc.) — shadcn-style components.
- **`src/components/app-sidebar.tsx`** — Main nav sidebar (Home, Accounts, Records, Assets, Statistics, Budgets).
- **`src/hooks/`** — Custom React hooks (e.g. `use-mobile.ts`).
- **`src/lib/utils.ts`** — `cn()` helper combining `clsx` + `tailwind-merge`.
- **`src/__tests__/`** — Jest tests mirroring `src/` structure; `test-utils.tsx` provides a custom `render` with providers.

### Testing conventions
Coverage thresholds: 70% for branches, functions, lines, and statements. Mock Next.js modules (`next/image`, `next/navigation`) at the test file level. Use semantic queries (`getByRole`, `getByText`) over test IDs.

## iOS (`apps/ios`)

SwiftUI app using MVVM. Features are organized as `Core/<FeatureName>/{Model,View,ViewModel}`.

- `Core/Network/APIConfig.swift` — environment-based base URL (`localhost:8080` in DEBUG, `api.vinance.app` in production).
- `Core/Root/ContentView.swift` — root view; auth wall is currently bypassed (shows `AppTabView` directly); backend integration pending.
- `App/VinanceApp.swift` — `@main` entry; injects `AuthViewModel` as an `@EnvironmentObject`.
- `Components/` — Shared reusable views (charts, transaction items, budget views).
- `Model/` — Domain models (`Transaction`, `User`, `BudgetCategory`, etc.).

## Database Migrations (`database/migrations`)

Standalone Go module — separate from the API.

```bash
cd database/migrations
go run main.go generate <migration_name>   # creates <timestamp>_<name>_up.sql + _down.sql
go run main.go migrate                    # runs the next un-applied migration
```

Migration names must use underscores (no spaces). The tool tracks applied versions in the database.
