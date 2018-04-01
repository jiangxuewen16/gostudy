package main

import (
	"os/exec"
	"fmt"
	"bytes"
	"io"
)

func main() {
	t := exec.Command("echo","-n","hello world!")
	if err := t.Start(); err != nil {
		fmt.Println(err)
	}

	out,err := t.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	var outBuff bytes.Buffer

	for {
		tmpout := make([]byte, 5)
		n,err := out.Read(tmpout)
		if err != nil {
			if err == io.EOF {
				return
			} else {
				fmt.Println(err)
			}
		}
		if n > 0 {
			outBuff.Write(tmpout[:n])
		}

	}

	fmt.Println(outBuff.String())
}
