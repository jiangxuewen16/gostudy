package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
)

func main() {
	cmd := exec.Command("ls")
	stdout, err := cmd.StdoutPipe() //指向cmd命令的stdout
	cmd.Start()
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))     //输出ls命令查看到的内容
}
