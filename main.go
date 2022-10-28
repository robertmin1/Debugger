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

type Breakpoint struct {
	Addr         uintptr
	OriginalData []byte
}

func setBreakpoint(pid int, addr uintptr) []byte {
	data := make([]byte, 1)
	if _, err := unix.PtracePeekData(pid, addr, data); err != nil {
		panic(err)
	}
	if _, err := unix.PtracePokeData(pid, addr, []byte{0xCC}); err != nil {
		panic(err)
	}
	return data
}

func resetBreakpoint(pid int, addr uintptr, data []byte) {
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
	if err := unix.PtraceSingleStep(pid); err != nil {
		panic(err)
	}
	// wait for it's execution and set the breakpoint again
	var status unix.WaitStatus
	if _, err := unix.Wait4(pid, &status, 0, nil); err != nil {
		panic(err.Error())
	}
	setBreakpoint(pid, addr)
}

func main() {
	runtime.LockOSThread()
	// start the process
	process := exec.Command("./WordGenerator")
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
		// wait for the interupt to come.
		var status unix.WaitStatus
		if _, err := unix.Wait4(pid, &status, 0, nil); err != nil {
			panic(err.Error())
		}
		fmt.Println("A breakpoint has been hit")
		// reset the breakpoint
		resetBreakpoint(pid, 0x47f0d4, data)
	}
}
