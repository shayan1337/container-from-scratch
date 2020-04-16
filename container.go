package main

import (
  "fmt"
  "os"
  "os/exec"
  "syscall"
)


func main() {
  switch os.Args[1] {
  case "run":
    run()
  case "child":
    child()
  default:
    panic("invalid command!!!")
  }
}

func run() {
  fmt.Printf("Running %v as PID %d \n", os.Args[2:], os.Getpid())

  args := append([]string{"child"}, os.Args[2:]...)

  cmd := exec.Command("/proc/self/exe", args...)

  cmd.Stdin = os.Stdin
  cmd.Stdout= os.Stdout
  cmd.Stderr = os.Stderr

  cmd.SysProcAttr = &syscall.SysProcAttr{
    Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
  }

  must(cmd.Run())
}

func child() {
  fmt.Printf("Running %v as PID %d \n", os.Args[2:], os.Getpid())

  cmd := exec.Command(os.Args[2], os.Args[3:]...)

  cmd.Stdin = os.Stdin
  cmd.Stdout= os.Stdout
  cmd.Stderr = os.Stderr

  syscall.Chroot("/root/containerfs")
  os.Chdir("/")
  syscall.Mount("proc", "proc", "proc", 0, "")
  
  must(cmd.Run())
}

func must(err error) {
  if err != nil {
    panic(err)
  }
}
