package main

import (
	"funollet/volume-notify/internal/notify"
	"funollet/volume-notify/internal/volume"
	"log"
	"os"
)

func main() {
	actualVolume, err := volume.GetVolume()
	if err != nil {
		os.Exit(1)
	}

	errSetVolume := volume.SetVolume("+10%")
	if errSetVolume != nil {
		os.Exit(1)
	}

	volume.ToggleMute()

	muted, errMute := volume.GetMute()
	if errMute != nil {
		os.Exit(1)
	}

	out, errNotify := notify.NotifyVolumeOsd(actualVolume, muted, "")
	// out, errNotify := notify.NotifyVolumeOsd(-1, true, "")
	if errNotify != nil {
		log.Println(out)
		os.Exit(1)
	}
}
