# Tarefeiro - Architecture Documentation

## Overview

Tarefeiro is a CLI task management application built with Go, following **Clean Architecture principles** with strong separation of concerns through **ports and adapters pattern**.

## Project Structure

```
magalu-cli/
├── cmd/tarefeiro/              # Application entrypoint
│   └── main.go                 # CLI bootstrap
├── internal/
│   ├── cli/                    # CLI command definitions (Cobra)
│   │   ├── root.go             # Root command handler
│   │   ├── start/              # Initialize CLI command
│   │   └── task/               # Task commands (add, list, update, delete, done, show)
│   ├── core/                   # Business logic layer
│   │   ├── domain/             # Domain entities (Task, Config)
│   │   ├── enum/               # Type-safe enumerations (Status, Priority)
│   │   └── usecase/            # Use cases (business operations)
│   ├── persistence/            # Data layer
│   │   └── service.go          # Generic persistence services
│   ├── port/                   # Interfaces (contracts)
│   │   ├── domain.go           # Domain entity contracts
│   │   ├── persistence.go      # Persistence contracts
│   │   ├── render.go           # Render contracts
│   │   ├── runner.go           # Handler contracts
│   │   └── usecase.go          # Use case contracts
│   ├── render/                 # Presentation/output formatters
│   │   └── task/               # Task renderers (table, detail, list)
│   └── runner/                 # Command handlers (cobra RunE implementations)
│       ├── root/               # Root handler
│       ├── start/              # Start/init handler
│       └── task/               # Task command handlers
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Architecture Layers

### 1. **Presentation Layer** (`cli/` + `runner/`)
Handles user interaction via CLI commands (Cobra framework).

- **CLI Definition** (`internal/cli/`): Defines commands, flags, and command structure
- **Runners** (`internal/runner/`): Implements the actual business logic that Cobra calls via `RunE`

**Key Components:**
- `root.go`: Bootstrap, initialization checks, global pre-runs
- `task/`: Commands for CRUD operations on tasks
- `start/`: Setup/initialization command

**Example Flow:**
```
User Input → Cobra Command → Runner.Run() → UseCase → Domain → Persistence
```

### 2. **Business Logic Layer** (`core/`)
Encapsulates all business rules and domain concepts.

#### **Domain** (`core/domain/`)
Core entities representing real-world objects:
- `Task`: Represents a task with title, description, priority, status, timestamps, tags
- `ConfigCommand`: Configuration data (database path)

**Key Methods:**
- `GetCollection()`: Returns storage collection name (e.g., "tasks")
- `GetID()`: Returns unique identifier
- `IsEqual(other)`: Filtering logic (case-insensitive contains match)
- `IsOverdue()`: Business rule (task past due date and not done)
- `MarkAsDone()`: State transition
- `IsEmpty()`: Checks if entity is uninitialized

#### **Enumerations** (`core/enum/`)
Type-safe enums replacing magic strings:
- `Status`: todo, in_progress, done, unknown
- `Priority`: low, medium, high, unknown
- Color mappings for CLI output

#### **Use Cases** (`core/usecase/`)
Orchestrates domain entities and persistence:
- `Create`: Validates and saves new tasks
- `List`: Retrieves tasks with filtering
- `Update`: Modifies existing tasks
- `Delete`: Removes tasks
- `Done`: Marks task as completed

### 3. **Ports & Adapters** (`port/`)
**Contracts** that define interfaces between layers (dependency inversion).

**Key Interfaces:**
- `Domain`: Any entity that stores/retrieves (GetCollection, GetID)
- `FilterDomain[T]`: Entities that support filtering + pagination
- `Persistence[T]`: Save, Find, FindAll, Delete operations
- `SaveUseCase[T]`: Save contract
- `FindAllUseCase[T]`: Query contract  
- `Renderer`: Render contract for outputting data

**Example:**
```go
type Persistence[T Domain] interface {
    Save(target T) (T, error)
    Find(target T) (T, error)
    FindAll(filter T) ([]T, error)
    Delete(target T) error
}
```

### 4. **Data Layer** (`persistence/`)
Implements port contracts for actual storage.

**Service Types:**
- `ServiceSaver[T]`: Saves entities
- `ServiceFinder[T FilterDomain]`: Queries entities (with filtering)
- `ServiceDeleter[T]`: Deletes entities

**Storage Backend:** Scribble (file-based JSON store)
```
./tarefeiro/
├── tasks/              # Collection: task entities
├── config/             # Collection: configuration
└── start/              # Collection: bootstrap state
```

**Filtering Logic:**
1. `Unmarshal` each JSON item → `T`
2. Call `T.IsEqual(filter)` 
3. Append matching items
4. Return (max 10 items)

### 5. **Presentation Layer** (`render/`)
Formats domain entities for output.

**Renderers:**
- `List`: Compact table view (ID, Title, Priority, Status)
- `Detail`: Full table with all fields and proper formatting
- `YAML` (extensible): Could output YAML format

**Implementation:**
- Uses `tablewriter` for structured output
- Normalizes dates with padding (e.g., `"2006-01-02 15:04"`)
- Handles nil values gracefully

## Data Flow Examples

### Creating a Task
```
CLI Input:
tarefeiro add "Study Go" --priority high --tags dev,estudos

Flow:
1. Cobra parses flags → CreateRunner
2. CreateRunner reads args + flags → domain.NewTask()
3. SaveUseCase validates → calls persistence.Save()
4. ServiceSaver writes to ./tarefeiro/tasks/{ulid}.json
5. Response: "Task created with ID: ..."
```

### Listing Tasks with Filter
```
CLI Input:
tarefeiro list --title study --status done

Flow:
1. Cobra parses flags → ListRunner
2. ListRunner builds filter Task with Title="study", Status=StatusDone
3. FindAllUseCase calls persistence.FindAll(filter)
4. ServiceFinder reads all from ./tarefeiro/tasks/
5. For each file:
   - Unmarshal → Task
   - Call filter.IsEqual(task)
   - If match, append
6. List renderer formats results as table
7. Output: Table with matching tasks
```

### Initialization (Start Command)
```
CLI Input:
tarefeiro start

Pre-Check (in root.PersistentPreRunE):
1. If not "start"/"help"/"completion" → check ./tarefeiro/start/default.json exists
2. If missing → Error: "CLI not initialized. Run `tarefeiro start` first"

Start Command:
1. Prompts user for database path (default: env TAREFEIRO_PATH)
2. Creates ConfigCommand
3. SaveUseCase → persistence.Save()
4. Writes config
5. Creates bootstrap marker → ./tarefeiro/start/default.json
6. Future commands: bootstrap check passes
```

## Key Design Patterns

### 1. **Generic Services**
```go
type ServiceFinder[T FilterDomain[T]] struct { *scribble.Driver }
```
Allows same persistence logic for any entity implementing `FilterDomain`.

### 2. **Null Object (IsEqual for Filters)**
```go
func (t *Task) IsEqual(other *Task) bool {
    if t.IsEmpty() {  // Empty filter = match all
        return true
    }
    // Matching logic
}
```

### 3. **Dependency Injection**
Runners receive use cases in constructor:
```go
func NewListRunner(useCase port.FindAllUseCase[*Task]) *ListRunner
```

### 4. **Adapter Pattern**
- `persistence.ServiceFinder` adapts `scribble.Driver` to `Persistence` port
- `runner.ListRunner` adapts CLI to business logic

### 5. **Strategy Pattern (Filtering)**
- Each domain entity defines its own filter logic via `IsEqual()`
- Filtering is composable and extensible

## Configuration & Environment

- **Database Path**: Controlled by `TAREFEIRO_PATH` env or `--path` flag (if added)
- **Default**: `./tarefeiro/`
- **Bootstrap Check**: Looks for `./tarefeiro/start/default.json`
- **Data Format**: JSON files organized by collection

## Command Lifecycle

```
1. User runs: tarefeiro <command> [args] [flags]
2. Cobra parses input
3. root.PersistentPreRunE executes:
   - Checks if bootstrap marker exists
   - Blocks if not initialized (except start/help)
4. Runner.Run() executes:
   - Reads flags/args
   - Calls use case
   - Use case calls persistence
   - Persistence calls storage backend
   - Result returned through layers
5. Renderer formats output
6. Output displayed to terminal
```

## Testing Strategy

- **Unit Tests**: Domain logic (`IsEqual`, `IsOverdue`, `IsEmpty`)
- **Integration Tests**: UseCase + Persistence layer
- **CLI Tests**: Cobra command validation (present in `*_test.go` files)

**Example Test Pattern:**
```go
func TestTaskIsOverdue(t *testing.T) {
    // Domain logic testing
    past := time.Now().Add(-1 * time.Hour)
    task := &Task{EstimatedDoneAt: &past, Status: StatusTodo}
    assert.True(t, task.IsOverdue())
}

func TestListFilterByTitle(t *testing.T) {
    // Integration testing
    filter := &Task{Title: "study"}
    results, err := listUseCase.All(filter)
    assert.NoError(t, err)
    assert.Equal(t, 1, len(results))
}
```

## Extension Points

### Adding a New Command
1. Create domain entity if needed
2. Define use case in `core/usecase/`
3. Create runner in `internal/runner/`
4. Add CLI command builder in `internal/cli/`
5. Register in `root.go` Execute()

### Adding a New Output Format
1. Create renderer in `internal/render/task/` (e.g., `json.go`, `csv.go`)
2. Implement renderer interface
3. Add `--output` flag to CLI
4. Switch renderers in runner based on flag

### Changing Storage Backend
1. Implement `Persistence` port with new backend (e.g., PostgreSQL, SQLite)
2. Replace `scribble.Driver` in `Execute()` with new driver
3. No other changes needed (clean architecture benefit!)

**Example:**
```go
// Old
db, err := scribble.New(dbPath, nil)

// New - with PostgreSQL
db, err := postgres.New(connString)
// Everything else stays the same!
```

## Dependencies

- **Cobra**: CLI framework (command parsing)
- **Scribble**: File-based JSON storage
- **Tablewriter**: CLI table formatting
- **YAML**: Data marshaling (extensible)
- **JSON**: Default data marshaling
- **ULID**: Unique ID generation (sortable, distributed-friendly)

## Build & Run

```bash
# Build
make build

# Local install
go install ./cmd/tarefeiro

# Run
tarefeiro --help
tarefeiro start
tarefeiro add "Task" --priority high
tarefeiro list
tarefeiro list --title study
tarefeiro update <id> --status done
tarefeiro delete <id>
```

## Error Handling

- **Pre-Run Checks**: Initialization required before any operation
- **Domain Validation**: Entity methods check invariants (e.g., `IsEmpty()`)
- **Persistence Errors**: Wrapped with context for clarity
- **User Feedback**: Clear error messages guiding next steps

**Example Error Flow:**
```
CLI Input: tarefeiro list (without init)
↓
root.PersistentPreRunE → Checks for ./tarefeiro/start/default.json
↓
Not found → Returns error:
"CLI not initialized. Run `tarefeiro start` first"
↓
User sees message and knows what to do next
```

## Performance Considerations

- **Filtering**: Limited to 10 items max (configurable in persistence layer)
- **File I/O**: Suitable for small-to-medium task lists
- **Scalability**: Infrastructure-ready (swap Scribble for database without code changes)

## Future Improvements

1. **Pagination**: Extend `FilterDomain` with proper pagination support
2. **Database Backend**: PostgreSQL adapter for larger datasets
3. **Authentication**: User-based task isolation
4. **Sync**: Cloud sync capability
5. **Web UI**: API layer + frontend
6. **Webhooks**: Task event notifications
7. **Search**: Full-text search with tags/categories

