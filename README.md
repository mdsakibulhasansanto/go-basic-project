
 📝 Project Overview

A Go (Golang) web app using the Gin framework, designed with a modular and clean architecture. It includes user authentication,product management, and Redis caching.

 🚀 Key Features
 🔐 JWT Authentication** for secure endpoints
 📦 Product CRUD** operations
 🗃️ MySQL database with **GORM** ORM
 ⚡ Redis caching for performance
 🧱 Modular code structure: handler, service, repo, etc.

 📁 Project Structure

`main.go` — App entry point
`models/` — Data models
 handler/` — HTTP handlers
`services/` — Business logic
`repositories/` — DB & Redis access
`db/` — DB connection & migration
`middlewares/` — JWT auth
`routes/` — API routes
`cache/` — Redis logic
`utils/` — Helper functions
 go.mod` / `go.sum` — Dependencies

- 🛠️ Tech Stack

* Go (Golang)
* Gin
* GORM
* MySQL
* Redis
