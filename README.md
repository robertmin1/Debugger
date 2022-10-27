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

## TODO.

## Information About The Program

Since we want to print `A breakpoint has been hit` before our Word Generator program prints a random string we need to pause the Word Generator before it prints. 

Line 21 is where the printing occurs, thus we need to establish a breakpoint there. The instruction address on line 21 will be used to set the breakpoint.

The command : `objdump --dwarf=decodedline ./WordGenerator` will be used to get the address of line 21. (The command: `go install WordGenerator.go` must before executed beforehand.)

We use `PtracePeekData` to read the memory first since it will be used later. Then, we use `PtracePokeData` to replace the data at the address with 0xCC, causing the CPU to halt the application anytime it detects data int 3.

We need the program to continue running normally once the breakpoint is reached. The application doesn't function normally since we previously changed the data at the address. As a result, original data must be used to replace the int 3. Since we captured the original data using `PtracePeekData`, we just revert to the original data.

