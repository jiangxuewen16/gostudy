package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
	"bytes"
)

func main() {
	t := exec.Command("ps","aux")
	t2 := exec.Command("grep","docker")

	out,err := t.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	var bf1 bytes.Buffer
	t.Stdout = &bf1

	if err := t.Start(); err != nil {
		fmt.Println(err)
	}

	b,e := ioutil.ReadAll(out)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(string(b))

	if err := t.Wait(); err != nil {
		fmt.Println(e)
	}


	out2, err := t2.StdoutPipe()

	if err:= t2.Start(); err != nil {
		fmt.Println(err)
	}

	t2.Stdout = &bf1
	var bf2 bytes.Buffer
	t2.Stdout = &bf2

	c, err := ioutil.ReadAll(out2)
	fmt.Println(c)

	if err := t2.Wait(); err != nil {
		fmt.Println(err)
	}

	fmt.Println()
}
