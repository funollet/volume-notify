package main

import (
	"flag"
	"fmt"
	"funollet/volume-notify/internal/notify"
	"funollet/volume-notify/internal/volume"
	"log"
	"os"
)

func main() {
	var step int
	flag.IntVar(&step, "step", 1, "volume change step")
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Println("please provide a value")
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case "up":
		change := fmt.Sprintf("+%d%%", step)
		log.Println(change)
		errSetVolume := volume.SetVolume(change)
		if errSetVolume != nil {
			os.Exit(1)
		}
	case "down":
		change := fmt.Sprintf("-%d%%", step)
		errSetVolume := volume.SetVolume(change)
		if errSetVolume != nil {
			os.Exit(1)
		}
	case "mute":
		volume.ToggleMute()
	}

	volumePercent, err := volume.GetVolume()
	if err != nil {
		os.Exit(1)
	}

	muted, errMute := volume.GetMute()
	if errMute != nil {
		os.Exit(1)
	}

	out, errNotify := notify.NotifyVolumeOsd(volumePercent, muted, "")
	if errNotify != nil {
		log.Println(out)
		os.Exit(1)
	}
}
