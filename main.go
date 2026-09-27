package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatalf("expected 2 args")
	}

	sleep, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("failed to cast %s to int", args[1])
	}
	key := args[2]

	fmt.Printf("Typing key %s with interval %d seconds \n", key, sleep)

	for {
		robotgo.KeyTap(key)
		time.Sleep(time.Duration(sleep) * time.Second)
	}
}
