package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func run(cmd string) (result string, err error) {
  var out []byte
	out, err = exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Println(string(out))
	}
	return string(out), err
}

// type runShell struct {
//   command string
//   process func()
// }

// getVolume := runShell{
//   "command": "pactl get-sink-mute @DEFAULT_SINK@",
//   "process": isMutedProcess
// }

func processGetVolume(text string) (volume int, err error) {
	parsed := strings.Split(text, "\n")[0]
	parsedd := strings.Split(parsed, "/")[1]
	parseddd := strings.ReplaceAll(parsedd, "%", "")
	return strconv.Atoi(strings.TrimSpace(parseddd))
}

func getVolume() (volume int, err error) {
	// Get the current volume
	var cmdOut string
	cmdOut, err = run("pactl get-sink-volume @DEFAULT_SINK@")
	if err != nil {
		return 0, err
	}
  return processGetVolume(cmdOut)
}

func notifyVolumeOsd(percentage int, muted bool, icon string) (out string, err error) {
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

	return run(cmd)
}

func main() {
	// out, err := notifyVolumeOsd(50, false, "")
	// // out, err := notifyVolumeOsd(-1, true, "")
	// if err != nil {
	// 	log.Println(string(out))
	// 	os.Exit(1)
	// }

	volume, err := getVolume()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(volume)

	fmt.Println()
}
