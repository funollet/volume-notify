package main

import (
	"funollet/volume-notify/internal/notify"
	"funollet/volume-notify/internal/volume"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Println("please provide a value")
		os.Exit(1)
	}

	switch args[0] {
	case "up":
		errSetVolume := volume.SetVolume("+1%")
		if errSetVolume != nil {
			os.Exit(1)
		}
	case "down":
		errSetVolume := volume.SetVolume("-1%")
		if errSetVolume != nil {
			os.Exit(1)
		}
	case "mute":
		volume.ToggleMute()
	}

	actualVolume, err := volume.GetVolume()
	if err != nil {
		os.Exit(1)
	}

	muted, errMute := volume.GetMute()
	if errMute != nil {
		os.Exit(1)
	}

	out, errNotify := notify.NotifyVolumeOsd(actualVolume, muted, "")
	if errNotify != nil {
		log.Println(out)
		os.Exit(1)
	}
}
