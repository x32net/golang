// http://stackoverflow.com/questions/8875038/redirect-stdout-pipe-of-child-process-in-golang
package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "os/exec"
)

func checkError(err error) {
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
}

func main() {
    // Replace `ls` (and its arguments) with something more interesting
    cmd := exec.Command("ls", "-l")

    // Create stdout, stderr streams of type io.Reader
    stdout, err := cmd.StdoutPipe()
    checkError(err)
    stderr, err := cmd.StderrPipe()
    checkError(err)

    // Start command
    err = cmd.Start()
    checkError(err)

    // Don't let main() exit before our command has finished running
    defer cmd.Wait()  // Doesn't block

    // Non-blockingly echo command output to terminal
    go io.Copy(os.Stdout, stdout)
    go io.Copy(os.Stderr, stderr)

    // I love Go's trivial concurrency :-D
    fmt.Printf("Do other stuff here! No need to wait.\n\n")
}
