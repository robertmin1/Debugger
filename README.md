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
~~The test was implemented but it doesn't work as expected~~
## TODO.
### Implement Cirrus Test
~~The test was implemented but it doesn't work as expected~~
- The bash script for the test does work as expected when executed manually.
        

## Information About The Program

Since we want to print `A breakpoint has been hit` before our Word Generator program prints a random string we need to pause the Word Generator before it prints. 

Line 21 is where the printing occurs, thus we need to establish a breakpoint there. The instruction address on line 21 will be used to set the breakpoint.

The command : `objdump --dwarf=decodedline ./WordGenerator` will be used to get the address of line 21. (The command: `go install WordGenerator.go` must before executed beforehand.)

We use `PtracePeekData` to read the memory first since it will be used later. Then, we use `PtracePokeData` to replace the data at the address with 0xCC, causing the CPU to halt the application anytime it detects data int 3.

We need the program to continue running normally once the breakpoint is reached. The application doesn't function normally since we previously changed the data at the address. As a result, original data must be used to replace the int 3. Since we captured the original data using `PtracePeekData`, we just revert to the original data.

We have to execute the instruction at the address once more because it has already been executed. Once the instruction pointer is set to the address, the CPU will once more execute it. `PtraceGetRegs` and `PtraceSetRegs` are commands used for manipulating CPU registers.

Since we changed the register, the program will now run according to its usual flow if we continue. However, we want to hit the breakpoint once again, so we'll ask `ptrace` to carry out only the subsequent instruction before setting the breakpoint once more. `PtraceSingleStep` allows us to execute only one instruction.

Before our WordGenerator program prints a random string during execution, the message `A breakpoint has been hit` is printed.
The software performs as intended.

![pro](https://user-images.githubusercontent.com/104002271/198543797-0ffffa48-bc9e-40a5-85b7-7288444417bd.PNG)

