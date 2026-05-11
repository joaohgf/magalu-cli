## Tarefeiro - Task management CLI for the Magalu Cloud challenge
Tarefeiro is a command-line application written in Go to manage tasks locally.

### Prerequisites

Before installing the CLI, make sure Go and Make are installed on your machine.

- Go version: `1.26` (as defined in `go.mod`)
- Official download page: https://go.dev/dl/
- Official installation guide (macOS, Linux, Windows): https://go.dev/doc/install

Install Make (if needed):

```bash
# macOS (Homebrew)
brew install make

# Ubuntu/Debian
sudo apt-get update && sudo apt-get install -y make

# Fedora
sudo dnf install -y make
```

Windows options:

- Use Git Bash (includes `make` in many setups), or
- Install via Chocolatey: `choco install make`, or
- Use WSL and follow the Linux commands above.

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
go build -o ./ ./cmd/tarefeiro
./tarefeiro --help
```

### Docker (optional)

Build the image:

```bash
docker build -t tarefeiro:local .
```

Run the CLI help:

```bash
docker run --rm tarefeiro:local --help
```

Persist tasks locally by mounting the project `data` directory:

```bash
docker run --rm -v "$(pwd)/data:/app/data" tarefeiro:local add "Study Go" --priority high --tags dev,estudos
docker run --rm -v "$(pwd)/data:/app/data" tarefeiro:local list
```

### Makefile commands

List all available automation commands:

```bash
make help
```

Common targets:

- `make build`: Build the CLI binary (`./cli`)
- `make test`: Run unit tests
- `make test-output`: Run unit tests and generate `coverage.out` + `coverage.html`

### Usage

Show all available commands:

```bash
tarefeiro --help
```

### Autocomplete (zsh)

Load autocomplete for the current session:

```bash
source <(tarefeiro completion zsh)
```

Persist for new terminals:

```bash
mkdir -p "$(brew --prefix)/share/zsh/site-functions"
tarefeiro completion zsh > "$(brew --prefix)/share/zsh/site-functions/_tarefeiro"
autoload -U compinit && compinit
```

Validate that `add --priority` suggestions are available:

```bash
tarefeiro __complete add --priority ""
```

Expected suggestions include: `low`, `medium`, `high`.

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
-- Filter by status and priority
tarefeiro list -s done -p high
-- Filter by title
tarefeiro list -t "study" -T work
-- Filter by description
tarefeiro list -d "cli"
-- Paginate results (page 2, 5 items per page)
tarefeiro list -P 2 -S 5
-- Interactive selection mode
tarefeiro list --interactive
```

Show task details by ID:

```bash
tarefeiro show "01KR8021T5FZE79CSANG5FACK0"
```

Mark a task as completed:

```bash
tarefeiro complete "01KR8021T5FZE79CSANG5FACK0"
```

Edit a task:

```bash
tarefeiro edit "01KR8021T5FZE79CSANG5FACK0" --title "Study Go deeply" --priority high --tags dev,go
```

Delete a task:

```bash
tarefeiro delete "01KR8021T5FZE79CSANG5FACK0"
```

Use output formats (`table`, `json`, `yaml`):
#### Default is `table`

```bash
tarefeiro list -o table
tarefeiro list -o json
tarefeiro show "01KR8021T5FZE79CSANG5FACK0" -o yaml
```
