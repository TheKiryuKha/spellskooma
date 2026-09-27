package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	hook "github.com/robotn/gohook"

	_ "embed"
)

// Amount of casts to reach out 10 levels in magic school
//
// 100 exp for new level
// 100 * 10 levels = 1000
// 1000 / 4 exp per cast = 250
const casts = 250

//go:embed assets/level_up.mp3
var levelUpSound []byte

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
		playLevelUpSound()

		hook.End()
	})

	s := hook.Start()
	<-hook.Process(s)
}

func playLevelUpSound() {
	reader := io.NopCloser(bytes.NewReader(levelUpSound))

	streamer, format, err := mp3.Decode(reader)
	if err != nil {
		log.Fatalf("failed to decode melody file: %v", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
