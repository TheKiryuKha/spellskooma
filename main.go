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

// Amount of casts to reach out 10 levels in magic school
//
// 100 exp for new level
// 100 * 10 levels = 1000
// 1000 / 4 exp per cast = 250
const casts = 250

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
	fmt.Printf("==== Type ' to start casting ==== \n")

	hook.Register(hook.KeyDown, []string{"'"}, func(e hook.Event) {
		for i := 0; i < casts; i++ {
			robotgo.KeyDown(key)
			time.Sleep(50 * time.Millisecond)
			robotgo.KeyUp(key)

			time.Sleep(time.Duration(delay) * time.Millisecond)
		}

		fmt.Printf("U should already have +10 levels in ur magic school!")
		hook.End()
	})

	s := hook.Start()
	<-hook.Process(s)
}
