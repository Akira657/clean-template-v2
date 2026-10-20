# clean-template

A Go web service scaffold built around Uncle Bob's Clean Architecture, using
[Gin](https://github.com/gin-gonic/gin) for HTTP and [GORM](https://gorm.io/)
over PostgreSQL for persistence.

## Layers

Dependencies only ever point inward, toward `domain`:

```
cmd  ->  bootstrap  ->  api/route  ->  api/controller  ->  usecase  ->  repository  ->  domain
                                                        \_____________________________/
                                                     (both usecase and repository depend
                                                      only on the interfaces in domain)
```

- **`domain/`** — the innermost layer. Plain structs (`User`, `Task`, `Profile`)
  plus the interfaces every other layer is written against
  (`TaskRepository`, `TaskUsecase`, `UserRepository`, `ProfileUsecase`).
  It has no dependency on any other package in the project.
- **`usecase/`** — business logic. Each usecase (`taskUsecase`,
  `profileUsecase`) implements a `domain` usecase interface and depends only
  on a `domain` repository interface, never on a concrete repository
  implementation.
- **`repository/`** — data access. Implements the `domain` repository
  interfaces on top of `*gorm.DB` (`taskRepository`, `userRepository`).
- **`api/controller/`** — HTTP handlers. Depend only on `domain` usecase
  interfaces, decode/encode JSON, and translate results into
  `domain.SuccessResponse` / `domain.ErrorResponse`.
- **`api/route/`** — the composition root per feature. Each `NewXRouter`
  wires a concrete repository into a concrete usecase into a controller,
  then registers its Gin routes (see `task_route.go`, `profile_route.go`,
  `health_route.go`). `route.Setup` groups these under public/protected
  routers.
- **`bootstrap/`** — application wiring, independent of the feature layers:
  environment loading (`env.go`, via Viper), logging setup (`logging.go`,
  via `log/slog`), and Postgres connection + schema bootstrap
  (`database.go`).
- **`cmd/`** — the entrypoint. `main.go` builds the `bootstrap.Application`,
  calls `route.Setup`, and starts the Gin server.
- **`db/postgres/`** — `postgres.go` defines the `PostgresDatabase` wrapper
  type; `migrations/` holds the raw SQL executed at startup.
- **`internal/fakeutil/`** — test-only helpers, not importable outside this
  module.
- **`domain/mocks/`** — generated mocks of the `domain` interfaces, used by
  the `usecase`/`repository` test suites.

## Request flow example

```
cmd/main.go
  -> bootstrap.App()            // load Env, init slog, connect Postgres
  -> route.Setup(...)
       -> route.NewTaskRouter(...)
            repository.NewTaskRepository(db)      // domain.TaskRepository
            usecase.NewTaskUsecase(taskRepo, ...)  // domain.TaskUsecase
            controller.TaskController{TaskUsecase: ...}
            group.POST("/task", taskController.Create)
```

A request to `POST /task` only ever flows through the `domain` interfaces —
the controller knows `domain.TaskUsecase`, the usecase knows
`domain.TaskRepository` — so any layer can be swapped or mocked
independently (see `domain/mocks/`).

> Note: `api/route/route.go` currently only registers the public health
> route; the profile/task routers are wired but commented out pending an
> auth middleware.

## Configuration

Environment variables are loaded via Viper in `bootstrap/env.go`
(`APP_ENV`, `LOG_LEVEL`, `SERVER_ADDRESS`, `CONTEXT_TIMEOUT`, admin/JWT
settings, and `PG_*` Postgres settings). Sensible defaults are set for local
development.

## Database bootstrap

On startup, `bootstrap.NewPostgresDatabase` connects as the admin user,
creates the target database if missing, then reconnects as the app user and
executes `db/postgres/migrations/file_management.sql` directly via GORM
(`bootstrap.CreateDatabase`) to create the schema. `golang-migrate` is a
dependency but the version-tracked migration path (`runMigrations`) is
currently unused in favor of this simpler one-shot script.
