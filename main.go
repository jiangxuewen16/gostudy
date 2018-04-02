package main

import (
	"os/exec"
	"fmt"
	"bytes"
	"io"
	"io/ioutil"
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


	b,e := ioutil.ReadAll(out)
	if e != nil {
		fmt.Println(e)
	}


	fmt.Println(string(b))
}
