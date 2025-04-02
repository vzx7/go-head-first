package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("hanldleSignal() Caught:", sig)
}

func main() {
	fmt.Printf("Process ID: %d\n", os.Getpid())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	start := time.Now()

	go func() {
		for sig := range sigs {
			switch sig {
			case syscall.SIGINT:
				fmt.Println("\nExecution time:", time.Since(start))
			case syscall.SIGTERM:
				handleSignal(sig)
				os.Exit(0)
			default:
				fmt.Println("Caught:", sig)
			}
		}
	}()
	for {
		fmt.Print("+")
		time.Sleep(10 * time.Second)
	}
}
