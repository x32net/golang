package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	subProcess := exec.Command("sub") // your subProcess
	/*subProcessIn, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer subProcessIn.Close()*/ // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line

	var subProcessStdout bytes.Buffer
	subProcess.Stdout = &subProcessStdout //os.Stdout
	var subProcessStderr bytes.Buffer
	subProcess.Stderr = &subProcessStderr     //os.Stderr
	if err = subProcess.Start(); err != nil { // Use start, not run
		fmt.Println("An error occured: ", err)
	}
	/*io.WriteString(subProcessIn, "4\n")*/
	subProcess.Wait()
	fmt.Println("subProcessStderr", subProcessStderr.String())
	fmt.Println("subProcessStdout", subProcessStdout.String())
}

/* http://stackoverflow.com/questions/8875038/redirect-stdout-pipe-of-child-process-in-golang
Now I want to have the stdout of the child program in my terminal window where I started the parent program.
No need to mess with pipes or goroutines, this one is easy.

func main() {
    cmd := exec.Command("ls", "-l") // Replace `ls` (and its arguments) with something more interesting
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}
*/
