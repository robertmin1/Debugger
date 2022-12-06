package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

func setBreakpoint(pid int, addr uintptr) []byte {
	data := make([]byte, 1)
	// Read the memory since it will be used later
	if _, err := unix.PtracePeekData(pid, addr, data); err != nil {
		panic(err)
	}
	// Replacing the data at the address with 0xCC
	if _, err := unix.PtracePokeData(pid, addr, []byte{0xCC}); err != nil {
		panic(err)
	}
	// We return the original data
	return data
}

func resetBreakpoint(pid int, addr uintptr, data []byte) {
	// Replace the data in the address with the original data
	if _, err := unix.PtracePokeData(pid, addr, data); err != nil {
		panic(err.Error())
	}
	// set the execution flow to continue
	regs := &unix.PtraceRegs{}
	if err := unix.PtraceGetRegs(pid, regs); err != nil {
		panic(err)
	}
	
	regs.Rip = uint64(addr)
	if err := unix.PtraceSetRegs(pid, regs); err != nil {
		panic(err)
	}
	// execute only one instruction
	if err := unix.PtraceSingleStep(pid); err != nil {
		panic(err)
	}
	// wait for it's execution and set the breakpoint again
	var status unix.WaitStatus
	if _, err := unix.Wait4(pid, &status, 0, nil); err != nil {
		panic(err.Error())
	}
	// set the breakpoint again
	setBreakpoint(pid, addr)
}

func main() {
	runtime.LockOSThread()
	// start the process
	process := exec.Command("testdata/WordGenerator")
	process.SysProcAttr = &syscall.SysProcAttr{Ptrace: true, Setpgid: true, Foreground: false}
	process.Stdout = os.Stdout
	if err := process.Start(); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	// get the pid of the process
	pid := process.Process.Pid
	data := setBreakpoint(pid, 0x47f0d4)
	for {
		if err := unix.PtraceCont(pid, 0); err != nil {
			panic(err.Error())
		}
		// wait for the interrupt to come.
		var status unix.WaitStatus
		if _, err := unix.Wait4(pid, &status, 0, nil); err != nil {
			panic(err.Error())
		}
		fmt.Println("A breakpoint has been hit")
		// reset the breakpoint
		resetBreakpoint(pid, 0x47f0d4, data)
	}
}
