package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	fmt.Println("Unlocker starting...")

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// time for alt + tab
	time.Sleep(5 * time.Second)

	for i := 0; i <= 9999; i++ {
		inputCode(kb, i)
		clearInput(kb, 4)
	}
}

func inputCode(kb keybd_event.KeyBonding, num int) {
	var digitKeycodes = [10]int{
		keybd_event.VK_0,
		keybd_event.VK_1,
		keybd_event.VK_2,
		keybd_event.VK_3,
		keybd_event.VK_4,
		keybd_event.VK_5,
		keybd_event.VK_6,
		keybd_event.VK_7,
		keybd_event.VK_8,
		keybd_event.VK_9,
	}

	numStr := fmt.Sprintf("%04d", num)

	fmt.Println(numStr)

	for _, ch := range numStr {
		// Convert the character to a key code
		key := digitKeycodes[ch-'0']

		// Press the key
		kb.SetKeys(key)
		err := kb.Launching()
		if err != nil {
			panic(err)
		}

		sleepRandom(50, 200)
	}

	// Press the Enter key
	kb.SetKeys(keybd_event.VK_ENTER)
	err := kb.Launching()
	if err != nil {
		panic(err)
	}

	sleepRandom(50, 200)
}

func clearInput(kb keybd_event.KeyBonding, inputSize int) {
	key := keybd_event.VK_BACKSPACE
	for i := 1; i <= inputSize; i++ {
		kb.SetKeys(key)
		err := kb.Launching()
		if err != nil {
			panic(err)
		}
		sleepRandom(40, 150)
	}
}

func sleepRandom(min, max int) {
	milisec := rand.Intn(max-min) + min
	time.Sleep(time.Duration(milisec) * time.Millisecond)
}