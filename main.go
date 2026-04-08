package main

import (
	"apple_system_sound/system_tools"
	"time"
)

func main() {
	system_tools.PlaySound()
	time.Sleep(1 * time.Second)
}
