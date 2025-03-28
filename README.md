# LeetX

LeetX is a command-line utility designed to streamline the process of fetching LeetCode problems and setting them up in your local development environment. It automates the creation of a directory for each problem, including the problem description and a code template in your preferred programming language.

## Installation

To install LeetX, follow these steps:

1. **Ensure Go is installed**: Download and install Go from [golang.org](https://golang.org/dl/) if you don’t already have it.

2. **Clone the repository**:
   ```bash
   git clone https://github.com/vitalygi/leetx.git
   ```

3. **Navigate to the project directory**:
   ```bash
   cd leetx
   ```

4. **Build and install the tool**:
   ```bash
   make install
   ```
   This will compile the Go code and install the `leetx` binary to your Go bin directory (e.g., `$GOPATH/bin` or `$HOME/go/bin`).

## Usage

To fetch a LeetCode problem and set it up locally, use the following command:

```bash
leetx -get <url/id/title>
```

### Example
```bash
leetx -get 1 -l go 
```
This will create a directory named after the problem (e.g., `1.Two_Sum`) containing:
- `problem.md`: A file with the problem description.
- `main.go`: A template in the specified language (in example in golang).

### Optional Flags
- **Specify a programming language** with the `-l` flag:
  ```bash
  leetx -get two-sum -l go
  ```
  This generates a `main.go` file with a Go code template.

- **Customize the code filename** with the `-f` flag:
  ```bash
  leetx -get 1 -l go -f solution.go
  ```
  This creates a `solution.go` file instead of the default `main.go`.

If no language is specified, only the `problem.md` file will be created.

## Supported Languages

LeetX supports code templates for the following programming languages:

- **C++** (`cpp`, `c++`) → `main.cpp`
- **Java** (`java`) → `Main.java`
- **Python** (`python`, `python3`) → `main.py`
- **C** (`c`) → `main.c`
- **C#** (`c#`) → `Program.cs`
- **JavaScript** (`javascript`, `js`) → `index.js`
- **TypeScript** (`typescript`, `ts`) → `index.ts`
- **PHP** (`php`) → `index.php`
- **Swift** (`swift`) → `main.swift`
- **Kotlin** (`kotlin`) → `Main.kt`
- **Dart** (`dart`) → `main.dart`
- **Go** (`go`, `golang`) → `main.go`
- **Ruby** (`ruby`) → `main.rb`
- **Scala** (`scala`) → `Main.scala`
- **Rust** (`rust`) → `main.rs`
- **Racket** (`racket`) → `main.rkt`
- **Erlang** (`erlang`) → `main.erl`
- **Elixir** (`elixir`) → `main.ex`

## Example Output

After running:
```bash
leetx -get https://leetcode.com/problems/two-sum/ -l go
```

Your directory will look like this:
```
1.Two_Sum/
├── main.go      # Go code template for "Two Sum"
└── problem.md   # Problem description
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.