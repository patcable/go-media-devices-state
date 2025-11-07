package mediadevicesstate

import (
	"github.com/patcable/go-media-devices-state/pkg/camera"
	"github.com/patcable/go-media-devices-state/pkg/debug"
	"github.com/patcable/go-media-devices-state/pkg/microphone"
)

// IsCameraOn returns true is any camera in the system is ON
func IsCameraOn() (bool, error) {
	return camera.IsCameraOn(false)
}

// IsMicrophoneOn returns true is any camera in the system is ON
func IsMicrophoneOn() (bool, error) {
	return microphone.IsMicrophoneOn(false)
}

// Debug calls all available device functions and prints the results
func Debug() {
	debug.Debug()
}
