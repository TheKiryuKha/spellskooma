package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatalf("expected 2 args")
	}

	delay, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("failed to cast %s to int", args[1])
	}
	key := args[2]

	fmt.Printf("Typing key %s with interval %d miliseconds \n", key, delay)

	cast(key, delay)
}

func cast(key string, delay int) {
	fmt.Printf("==== Type %s to start casting ==== \n", key)

	hook.Register(hook.KeyDown, []string{key}, func(e hook.Event) {
		for {
			robotgo.KeyDown(key)
			time.Sleep(50 * time.Millisecond)
			robotgo.KeyUp(key)

			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}
