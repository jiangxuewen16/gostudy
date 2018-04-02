package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
)

func main() {
	t := exec.Command("ps","aux")
	t2 := exec.Command("grep","docker")

	out,err := t.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

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

	t2.Stdout = &t.Stdout
	out2, err := t2.StdoutPipe()

	if err:= t2.Start(); err != nil {
		fmt.Println(err)
	}

	c, err := ioutil.ReadAll(out2)
	fmt.Println(c)

	if err := t2.Wait(); err != nil {
		fmt.Println(err)
	}

	fmt.Println()
}
