# cli-password-manager

A command-line password manager written in Go. Stores credentials per service in a local vault file (`.vault`).

## Requirements

- Go 1.25+

## Installation

```bash
git clone https://github.com/Leocanela279/cli-password-manager.git
cd cli-password-manager
go build -o vault ./cmd/main.go
```

## Usage

```
vault <command> [arguments]
```

## Commands

| Command | Arguments | Description |
|---------|-----------|-------------|
| `init` | — | Initialize a new empty vault |
| `add` | `<service> <username> <password>` | Add credentials for a service |
| `get` | `<service>` | Retrieve the password for a service |
| `list` / `ls` | — | List all saved services |
| `remove` / `rm` | `<service>` | Remove a service from the vault |
| `help` | — | Show available commands |

## Examples

```bash
# Initialize the vault
vault init

# Add credentials for a service
vault add github myuser mysecretpassword

# Retrieve the password for a service
vault get github

# List all saved services
vault list
vault ls

# Remove a service
vault remove github
vault rm github
```

## Project Structure

```
cli-password-manager/
├── cmd/
│   └── main.go          # Entry point, command routing
├── internal/
│   ├── vault/
│   │   └── vault.go     # Core vault logic (add, get, list, remove)
│   └── crypto/
│       └── crypto.go    # Encryption utilities (in progress)
├── vault-test/
│   └── .vault           # Local vault storage file (JSON)
└── go.mod
```

## Storage

Credentials are stored in a `.vault` file as a JSON map keyed by service name:

```json
{
  "github": {
    "username": "myuser",
    "password": "mysecretpassword"
  }
}
```

> **Note:** Encryption is not yet implemented. Passwords are currently stored in plain text.
