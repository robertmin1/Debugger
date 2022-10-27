# DEBUGGER

The project's objective is to make our debugger print `A breakpoint has been hit` before our word generator program prints a random string.

Rewriting registers and writing data to a specific address was done using `Ptrace.`

# Building

#### Prerequisites:

1. Ensure you have the Go tools installed.

#### Option A : Using Go build commands

1. Clone `DEBUGGER` via Git

2. Open a terminal in the cloned repository and run `export GOBIN=[PATH OF CLONED REPOSITORY]`

3. Run `go install WordGenerator.go`

4. Run `go run main.go`
