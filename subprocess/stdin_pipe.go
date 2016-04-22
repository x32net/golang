// http://stackoverflow.com/questions/23166468/how-can-i-get-stdin-to-exec-cmd-in-golang
package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	subProcess := exec.Command("abc") // your subProcess
	subProcessIn, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer subProcessIn.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line

	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr

	fmt.Println("START")                      //for debug
	if err = subProcess.Start(); err != nil { // Use start, not run
		fmt.Println("An error occured: ", err)
	}

	io.WriteString(subProcessIn, "4\n")
	subProcess.Wait()
	fmt.Println("END") //for debug
}
