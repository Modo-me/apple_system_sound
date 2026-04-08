package system_tools

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
}
