# container-from-scratch

The file uses some Go code to create a container from scratch. It used some Linux terminologies like namespaces, cgroup and chroot.

To run it:

go build container.go
./container run <command-you-want-to-run>
