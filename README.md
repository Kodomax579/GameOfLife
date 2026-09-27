# Conway's Game of Life in Go

A terminal-based **zero-player game** built in Go (Golang). This is my first Go project, developed to get experience in basic go features such as slices, loops and package architecture.

---

## 📌 About the Project

*Conway's Game of Life* is a game that plays itself. It runs with just 4 rules

### The Rules:

1. **Underpopulation:** Any live cell with fewer than 2 live neighbors dies.
2. **Survival:** Any live cell with 2 or 3 live neighbors lives on to the next generation.
3. **Overpopulation:** Any live cell with more than 3 live neighbors dies.
4. **Reproduction:** Any dead cell with exactly 3 live neighbors becomes a live cell.

---

## 🛠️ Project Structure

```text
├── internal/
│   └── handlers
│       └── board_handler.go    # Terminal rendering and output formatting
│   └── services
│       └── rules.go            # Game logic (Rule implementation)
├── cmd/
│   └── main.go                 # Entrypoint, grid initialization, and game loop
└── go.mod

```

---

## 🚀 Getting Started

### Prerequisites

* [Go](https://go.dev/dl/?utm_source=gemini) (1.18 or higher recommended)

### Running Locally

1. Clone the repository:

2. Run the application:
```bash
go run cmd/main.go

```

---

## ⚙️ Configuration

Board dimensions and refresh rate can be adjusted directly in `main.go`:

```go
width  := 30                    
height := 40                     
time.Sleep(100 * time.Millisecond)

```

---

## 📄 License

MIT License - see the [LICENSE](LICENSE) file for details.