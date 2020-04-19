# container-from-scratch

The file uses some Go code to create a container from scratch. It used some Linux terminologies like namespaces, cgroup and chroot.

To run it:

`go build container.go`

`./container run <command-you-want-to-run>`

This will create an executable called container and run it with the provided command line arguments.

The command here could be any shell command. For example: echo, /bin/bash
