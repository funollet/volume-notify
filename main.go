package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func run(cmd string) (out []byte, err error) {
	out, err = exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Println(string(out))
	}
	return out, err
}

func getVolume() (volume int, err error) {
	// Get the current volume
	var cmdOut []byte
	cmdOut, err = run("pactl get-sink-volume @DEFAULT_SINK@")
	if err != nil {
		return 0, err
	}
	parsed := strings.Split(string(cmdOut), "\n")[0]
	parsedd := strings.Split(parsed, "/")[1]
	parseddd := strings.ReplaceAll(parsedd, "%", "")
	return strconv.Atoi(strings.TrimSpace(parseddd))
}

func notifyVolumeOsd(percentage int, muted bool, icon string) (out []byte, err error) {
	timeout := 1000

	var iconWithDefaults string

	switch {
	case icon != "":
		iconWithDefaults = icon
	case muted:
		iconWithDefaults = "audio-volume-muted"
	default:
		iconWithDefaults = "audio-volume-medium"
	}

	cmd := fmt.Sprintf(`notify-send \
    --app-name volume-osd \
    --replace-id 2593 \
    --transient \
    --expire-time %d \
    --urgency low \
    --icon %s \
  `, timeout, iconWithDefaults)

	// The hint is only shown if percentage is in the 0..100 range.
	switch {
	case percentage > 0 && percentage <= 100:
		cmd += fmt.Sprintf("  --hint int:value:%d \\\n", percentage)
	}

	// notify-send raises an error if it receives an empty body
	cmd += "    ' '"

	fmt.Println(cmd)
	return run(cmd)
}

func main() {
	out, err := notifyVolumeOsd(50, false, "")
	// out, err := notifyVolumeOsd(-1, true, "")
	if err != nil {
		log.Println(string(out))
		os.Exit(1)
	}

	// volume, err := getVolume()
	// if err != nil {
	// 	os.Exit(1)
	// }
	// fmt.Println(volume)

	fmt.Println()
}
