package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
)

type Config struct {
	action string
	step   int
	help   bool
}

func parseFlags() Config {
	usage_header := `Changes volume up/down/muted and sends a desktop notification.

Usage: volume-notify [OPTION]... <up|down|mute>

Arguments:
  up      Turn volume up.
  down    Turn volume down.
  mute    Toggle mute status.

Options:
`
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage_header)
		pflag.PrintDefaults()
	}
	step := pflag.IntP("step", "s", 1, "volume change step, percentual")
	help := pflag.BoolP("help", "h", false, "show this message")

	pflag.Parse()

	if *help {
		pflag.Usage()
		os.Exit(0)
	}

	action := pflag.Arg(0)
	if action == "" {
		pflag.Usage()
		os.Exit(1)
	}

	return Config{
		action: action,
		step:   *step,
		help:   *help,
	}

}

func main() {
	config := parseFlags()

	switch config.action {
	case "up":
		change := fmt.Sprintf("+%d%%", config.step)
		err := setVolume(change)
		if err != nil {
			os.Exit(1)
		}
	case "down":
		change := fmt.Sprintf("-%d%%", config.step)
		err := setVolume(change)
		if err != nil {
			os.Exit(1)
		}
	case "mute":
		toggleMute()
	}

	volumePercent, err := getVolume()
	if err != nil {
		os.Exit(1)
	}

	muted, errMute := getMute()
	if errMute != nil {
		os.Exit(1)
	}

	out, errNotify := notifyVolumeOsd(volumePercent, muted, "")
	if errNotify != nil {
		log.Println(out)
		os.Exit(1)
	}
}
