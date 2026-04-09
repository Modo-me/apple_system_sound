package main

import (
	"time"

	"github.com/Modo-me/apple_system_sound/system_tools"
)

func main() {
	system_tools.PlaySound()
	time.Sleep(1 * time.Second)
}
