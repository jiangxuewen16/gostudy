package main

import (
	"os"
	"syscall"
	"os/signal"
	"fmt"
)

func main() {

	sinrev := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sinrev,sigs...)
	for sig := range sinrev {
		fmt.Println(sig)
	}
}
