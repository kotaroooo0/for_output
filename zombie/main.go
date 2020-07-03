package main

import (
	"os/exec"
)

func main() {
	exec.Command("./15time").Run()

	// cpath, err := exec.LookPath(os.Args[1])
	// if err != nil {
	// 	log.Fatalf("%s not found in $PATH.", os.Args[1])
	// }
	// args := os.Args[1:]
	// attr := syscall.ProcAttr{Files: []uintptr{0, 1, 2}}
	// pid, err := syscall.ForkExec(cpath, args, &attr)
	// if err != nil {
	// 	panic(err)
	// }
	// proc, err := os.FindProcess(pid)
	// status, err := proc.Wait()
	// if err != nil {
	// 	panic(err)
	// }
	// if !status.Success() {
	// 	fmt.Println(status.String())
	// }
}
