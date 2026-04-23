package system_tools

import "time"

/*
#cgo LDFLAGS: -framework AudioToolbox
#include <AudioToolbox/AudioServices.h>

static void playUserPreferredAlert() {
    AudioServicesPlayAlertSound(kSystemSoundID_UserPreferredAlert);
}
*/
import "C"

func PlaySound() {
	C.playUserPreferredAlert()
	time.Sleep(1 * time.Second)
}
