package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	Backoff := time.Duration(3)
	Retries := 3

	if len(os.Args) < 2 {
		fmt.Println("Usage: BACKOFF=3 RETRIES=3 retry <command> [args]")
		fmt.Println("BACKOFF and RETRIES are optional")
		os.Exit(1)
	}
	if backoff := os.Getenv("BACKOFF"); backoff != "" {
		backoffI, err := strconv.Atoi(backoff)
		if err != nil {
			fmt.Println("Error: environment variable BACKOFF must be an integer")
			os.Exit(1)
		}
		Backoff = time.Duration(backoffI)
	}
	if retries := os.Getenv("RETRIES"); retries != "" {
		retriesI, err := strconv.Atoi(retries)
		if err != nil {
			fmt.Println("Error: environment variable RETRIES must be an integer")
			os.Exit(1)
		}
		Retries = retriesI
	}

	thecommand := os.Args[1]
	theargs := os.Args[2:]

	for i := 0; i < Retries; i++ {
		err := run(thecommand, theargs...)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Printf("that was attempt %d, will try again in %d seconds, max attempts %d\n", i, Backoff, Retries)
			time.Sleep(Backoff * time.Second)
		} else {
			os.Exit(0)
		}
	}
	fmt.Println("failed after ", Retries, " attempts")
}

func run(thecommand string, theargs ...string) error {
	cmd := exec.Command(thecommand, theargs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
