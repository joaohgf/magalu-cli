# Tarefeiro - Architecture Documentation

## Overview

Tarefeiro is a Go CLI for local task management. The project follows a ports-and-adapters style with clear separation between:

- command layer (`internal/cli`, `internal/runner`)
- business rules (`internal/core`)
- contracts (`internal/port`)
- infrastructure (`internal/persistence`)
- output rendering (`internal/render`)

## Current Project Structure

```text
magalu-cli/
├── cmd/tarefeiro/
│   └── main.go
├── internal/
│   ├── cli/
│   │   ├── root.go
│   │   ├── task/
│   │   │   ├── add.go
│   │   │   ├── list.go
│   │   │   ├── show.go
│   │   │   ├── done.go
│   │   │   ├── edit.go
│   │   │   └── delete.go
│   │   └── report/
│   │       └── show.go
│   ├── core/
│   │   ├── domain/
│   │   │   ├── task.go
│   │   │   └── task_filter.go
│   │   └── usecase/
│   │       └── task/
│   │           ├── add.go
│   │           ├── list.go
│   │           ├── find.go
│   │           ├── done.go
│   │           ├── update.go
│   │           └── delete.go
│   ├── enum/
│   ├── errors/
│   ├── port/
│   ├── persistence/
│   │   ├── saver.go
│   │   ├── finder.go
│   │   ├── lister.go
│   │   └── deleter.go
│   ├── render/
│   │   ├── render.go
│   │   └── task/
│   │       ├── list.go
│   │       └── detail.go
│   └── runner/
│       ├── root/
│       │   └── root.go
│       └── task/
│           ├── add.go
│           ├── list.go
│           ├── find.go
│           ├── done.go
│           ├── update.go
│           └── delete.go
├── Makefile
├── README.md
└── ARCHITECTURE.md
```

## Runtime Flow

```text
main.go
  -> cli.Execute(ctx)
  -> Cobra root command + global flags
  -> runner.PreRun stores output format in context
  -> command runner executes
  -> use case validates/orchestrates
  -> persistence service reads/writes JSON files (scribble)
  -> renderer prints Table / JSON / YAML
```

## CLI Layer

### Root command

`internal/cli/root.go`:

- creates root command (`tarefeiro`)
- registers persistent `--output` (`-o`) flag
- wires all task commands
- initializes scribble driver at `./tarefeiro`

`internal/runner/root/root.go`:

- `Run`: shows help for root
- `PreRun`: injects output type into command context

### Available commands

Current command set:

- `add "Title"`
- `list`
- `show <task_id>`
- `complete`
- `edit "<task_id>"`
- `delete "<task_id>"`

Notes:

- `list` supports filters (`title`, `description`, `status`, `priority`) and pagination flags (`size`, `page`).
- `add` and `edit` support tags and estimated done date.

## Domain Layer

### `Task`

`internal/core/domain/task.go` represents the aggregate used by all use cases.

Main behaviors:

- `MarkAsDone()`
- `IsOverdue()`
- `Matches(other *Task)` for filtering
- identity/storage metadata (`GetID`, `GetCollection`)

### `TaskFilter`

`internal/core/domain/task_filter.go` encapsulates:

- query criteria (`Filter *Task`)
- pagination (`Page`, `Size`)
- result envelope (`Data`, `Total`)

It also implements the filter contract expected by the lister service (`IsEqual`, paging accessors, content setters/getters).

## Use Case Layer

Task use cases live in `internal/core/usecase/task`:

- `Add` (`add.go`): validates and saves new tasks
- `List` (`list.go`): validates page/size defaults and fetches filtered results
- `Find` (`find.go`): fetches one task by ID
- `Done` (`done.go`): finds + marks complete + saves
- `Update` (`update.go`): patch-like update flow
- `Delete` (`delete.go`): removes task by ID

The use cases depend on interfaces from `internal/port`, not concrete storage/render code.

## Ports (Contracts)

### Persistence contracts (`internal/port/persistence.go`)

- `PersistenceSaver[T]`
- `PersistenceFinder[T]`
- `PersistenceLister[T, F]`
- `PersistenceDeleter[T]`

### Use case contracts (`internal/port/usecase.go`)

- `SaveUseCase[T]`
- `FindUseCase[T]`
- `FindAllUseCase[T]`
- `DeleteUseCase[T]`

### Domain/filter contracts (`internal/port/domain.go`)

- `Domain`
- `FilterDomain[T]`

These are the core abstractions that make generics-based services possible.

## Persistence Layer

The project uses `github.com/nanobox-io/scribble` as a file-based JSON store.

Services are split by responsibility:

- `ServiceSaver`
- `ServiceFinder`
- `ServiceLister`
- `ServiceDeleter`

`ServiceLister` behavior (`internal/persistence/lister.go`):

1. reads all documents from collection
2. unmarshals each item
3. applies domain filter (`target.IsEqual(item)`)
4. paginates filtered items
5. stores results in filter envelope (`SetContent`, `SetTotal`)

## Rendering Layer

`internal/render/render.go` acts as an output strategy selector using the context value set by root pre-run.

Supported formats:

- `table` (default)
- `json`
- `yaml`

Concrete task views:

- `internal/render/task/list.go` for list results
- `internal/render/task/detail.go` for single task detail

Both implement table/json/yaml outputs.

## Data Storage Model

Scribble persists collections under `./data`. For tasks, files are stored under the `tasks` collection (one JSON file per entity ID).

## Testing Strategy (Current)

Tests are concentrated in:

- `internal/core/usecase/task/*_test.go`
- `internal/runner/task/*_test.go`

There are unit-test build tags (`//go:build unit`) in the suite, and the Makefile includes targets for running tests and generating coverage outputs.

## Key Design Decisions

- **Generics across ports/persistence**: reduces duplication while preserving type safety.
- **Context propagation from root**: global output mode is passed through context.
- **Use-case centric validation/orchestration**: command runners stay focused on CLI input mapping.
- **Separate render layer**: output format changes do not affect use cases/persistence.

## How to Extend

### Add a new command

1. Create/adjust use case in `internal/core/usecase`
2. Add runner in `internal/runner`
3. Add cobra command builder in `internal/cli`
4. Register command in `internal/cli/root.go`

### Add a new output format

1. Extend `enum.Output`
2. Add method in `port.View`
3. Implement it in task renderers
4. Route in `internal/render/render.go`

### Change storage backend

Replace scribble-backed implementations in `internal/persistence` with another adapter implementing the same persistence interfaces.
