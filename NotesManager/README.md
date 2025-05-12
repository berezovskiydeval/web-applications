# 📓 Notes Manager

A full‑stack web app for creating lists of notes, powered by **PostgreSQL + Go (Gin) API + React/Vite frontend**. The whole stack runs in Docker, so you only need Docker & docker‑compose to get started.

---

## ✨ Features

* **Lists & Notes** – create any number of lists, each with unlimited notes
* **Pin / Unpin** notes for quick access
* Full **CRUD** via REST API (`/api`) 
* **Search & sort** on the fly (both lists and notes)
* JWT‑based auth (signup / login)
* Hot‑reload dev setup (Vite + bind‑mounts) – save → browser updates instantly

---

## 🗂️ Project structure

```
├─ backend/                     # Go (Gin) API
│  ├─ cmd/                      # entrypoints (main.go)
│  ├─ configs/                  # YAML / TOML configs
│  ├─ docs/                     # Swagger docs
│  ├─ internal/                 # business logic (clean‑arch)
│  ├─ pkg/                      # reusable helpers
│  ├─ schema/                   # SQL migrations & seed data
│  ├─ .env                      # local env vars (dev only)
│  ├─ Dockerfile
│  ├─ Makefile                  # shortcuts (lint, test, run)
│  ├─ go.mod / go.sum
│  └─ ...
│
├─ frontend/
│  ├─ frontend/                 # Vite workspace
│  │   ├─ src/                  # React components & pages
│  │   ├─ public/               # static assets
│  │   ├─ .env                  # Vite env vars
│  │   ├─ Dockerfile
│  │   ├─ vite.config.js
│  │   └─ package.json / lockfile
│  └─ docker-compose.yml        # (frontend‑only hot‑reload, optional)
│
├─ docker-compose.yml           # stack: postgres, backend, frontend
├─ .env                         # root‑level overrides (optional)
└─ README.md
```

---

## 🚀 Quick start

### 1. Prerequisites

* **Docker** 20+
* **docker‑compose v2** (`docker compose` CLI)

### 2. Development run

```bash
# start stack (uses docker-compose.override.yml for fast dev images)
$ docker compose up
```

* **Frontend** → [http://localhost:5173](http://localhost:5173)
* **Backend API** → [http://localhost:8000](http://localhost:8000)
* **Postgres** (host) → `localhost:5436` (`postgres / n0t3sMaNaG3R`)

Hot‑reload ✅
Any change in `frontend/frontend/src` or `backend` is reflected immediately.

### 3. Production build (one‑off)

```bash
$ docker compose -f docker-compose.yml -f docker-compose.prod.yml up --build
```

Creates slim images with static assets served by Go.

---

## ⚙️ Configuration

Environment variables live in **docker‑compose.yml** – adjust if needed.

| Variable            | Default               | Description             |
| ------------------- | --------------------- | ----------------------- |
| `DB_HOST`           | `db`                  | service name in compose |
| `DB_PORT`           | `5432`                | internal PG port        |
| `DB_USERNAME`       | `postgres`            |                         |
| `DB_PASSWORD`       | `n0t3sMaNaG3R`        |                         |
| `JWT_SECRET`        | `supersecret`         | token signing key       |
| `VITE_API_BASE_URL` | `http://backend:8000` | used by React app       |

---

## 🐘 Database

PostgreSQL 15 runs in a named volume `notes-manager-db-data` so data persists between container restarts.

### Useful Makefile targets *(optional)*

```bash
make start-db    # standalone PG for debugging
make psql        # psql shell
make backup      # tar.gz dump into current dir
make restore     # restore from tar.gz
```

---

## 📑 API overview (quick)

| Method   | Endpoint               | Body                         |
| -------- | ---------------------- | ---------------------------- |
| `POST`   | `/api/auth/sign-up`    | `{name, username, password}` |
| `POST`   | `/api/auth/sign-in`    | `{username, password}`       |
| `GET`    | `/api/lists`           | params: `q, sort`            |
| `POST`   | `/api/lists`           | `{title, description}`       |
| `PUT`    | `/api/lists/:id`       | `{title, description}`       |
| `DELETE` | `/api/lists/:id`       |                              |
| `GET`    | `/api/lists/:id/items` | params: `q, sort`            |
| `POST`   | `/api/lists/:id/items` | `{title, content, pinned}`   |
| `PUT`    | `/api/items/:id`       | full object                  |
| `DELETE` | `/api/items/:id`       |                              |

> **Note:** `PUT /api/items/:id` expects full payload (title, content, pinned).

Full Swagger docs available at `/api/docs` (served by backend).

---

## 🖥️ Frontend scripts (inside container)

```bash
npm run dev        # vite hot‑reload (already in docker)
npm run build      # prod build
npm run preview    # preview prod build locally
```

---

## 🧪 Running backend tests

```bash
cd backend
go test ./...
```

---

## ✏️ Contributing

1. Fork → create feature branch → commit → PR
2. Follow Conventional Commits (`feat:`, `fix:` …)
3. Run `go vet` & `npm run lint` before pushing.

---

## 📄 License

MIT © 2025 Notes Manager Team
