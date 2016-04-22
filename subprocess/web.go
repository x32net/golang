package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	subProcess := exec.Command("sub") // your subProcess
	subProcessIn, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer subProcessIn.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line

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
