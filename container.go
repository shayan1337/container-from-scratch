package main

import (
  "fmt"
  "os"
  "os/exec"
)


func main() {
  switch os.Args[1] {
  case "run":
    run()
  default:
    panic("invalid command!!!")
  }
}

func run() {
  fmt.Printf("Running %v s PID %d \n", os.Args[2:], os.Getpid())

  cmd := exec.Command(os.Args[2], os.Args[3:]...)

  cmd.Stdin = os.Stdin
  cmd.Stdout= os.Stdout
  cmd.Stderr = os.Stderr

  cmd.Run()
}
