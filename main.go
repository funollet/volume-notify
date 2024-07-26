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
		log.Println("Error: can't get volume")
		os.Exit(1)
	}

	muted, _ := volume.GetMute()

	var out string
	out, err = notify.NotifyVolumeOsd(actualVolume, muted, "")
	// out, err := notify.NotifyVolumeOsd(-1, true, "")
	if err != nil {
		log.Println(out)
		os.Exit(1)
	}
}
