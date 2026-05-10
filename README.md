## Tarefeiro - Task management CLI for the Magalu Cloud challenge

[![Unit Tests](https://github.com/joaohgf/magalu-cli/actions/workflows/unit-tests.yml/badge.svg)](https://github.com/joaohgf/magalu-cli/actions/workflows/unit-tests.yml)
[![Coverage](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/joaohgf/magalu-cli/main/.github/badges/coverage.json)](https://github.com/joaohgf/magalu-cli/actions/workflows/unit-tests.yml)

Tarefeiro is a command-line application written in Go to manage tasks locally.

### Prerequisites

Before installing the CLI, make sure Go is installed on your machine.

- Official download page: https://go.dev/dl/
- Official installation guide (macOS, Linux, Windows): https://go.dev/doc/install

### Features
- Add tasks with title, description, status, priority, and optional estimated done date.
- List tasks with filters (`id`, `title`, `description`, `status`, `priority`).
- Persist tasks locally using a file-based database.

### Installation (recommended: `go install`)

Use Go to install the CLI directly from the module path:

```bash
go install github.com/joaohgf/magalu-cli/cmd/tarefeiro@latest
```

After installation, ensure your Go bin directory is in `PATH`:

```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

Run the CLI:

```bash
tarefeiro --help
```

### Local development build (optional)

```bash
git clone https://github.com/joaohgf/magalu-cli.git
cd magalu-cli
go build -o ./tarefeiro ./cmd/tarefeiro
./tarefeiro --help
```

### Usage

Show all available commands:

```bash
tarefeiro --help
```

Add a task:

```bash
tarefeiro add "Study Go" --priority high --tags dev,estudos
```

List tasks:

```bash
tarefeiro list
```

List tasks with filters:

```bash
tarefeiro list -s done -p high
tarefeiro list -t "study"
```

### Test coverage

Coverage is calculated by CI from unit tests and exposed as a badge in this README.
Click the badge to open the unit test workflow runs.
