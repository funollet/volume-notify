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
	// The hint is only shown if percentage is in the 0..100 range.

	defaultArgs := make([]string, 0)
	defaultArgs = append(defaultArgs,
		"--app-name", "volume-osd",
		"--replace-id", "2593",
		"--transient",
		"--expire-time", fmt.Sprint(1000), // ms
		"--urgency", "low",
	)

	args := make([]string, 0)
	args = append(defaultArgs, args...)

	var argIcon string
	switch {
	case icon != "":
		argIcon = icon
	case muted:
		argIcon = "audio-volume-muted"
	default:
		argIcon = "audio-volume-medium"
	}
	args = append(args, "--icon", argIcon)

	switch {
	case percentage > 0 && percentage <= 100:
		args = append(args, "--hint", fmt.Sprintf("int:value:%d", percentage))
	}

	return notifySend(args, "", "")
}

func notifySend(args []string, summary string, body string) (out []byte, err error) {
	switch summary {
	case "":
		// notify-send raises an error if it receives an empty body
		args = append(args, " ")
	default:
		args = append(args, summary)
	}

	args = append(args, body)

	return exec.Command("notify-send", args...).CombinedOutput()
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
