package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func notifyVolumeOsd(percentage int, muted bool, icon string) ([]byte, error) {
  // The hint is only shown if percentage is in the 0..100 range.

	defaultArgs := make([]string, 0)
	defaultArgs = append(defaultArgs,
		"--app-name", "volume-osd",
		"--replace-id", "2593",
		"--transient",
		"--expire-time", fmt.Sprint(1000),
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

func notifySend(args []string, summary string, body string) ([]byte, error) {
  switch summary {
  case "":
    // notify-send raises an error if it receives an empty body
	  args = append(args, " ")
  default:
	  args = append(args, summary)
  }

  args = append(args, body)

	out, err := exec.Command("notify-send", args...).CombinedOutput()
	return out, err
}

func main() {
	// out, err := notifyVolumeOsd(50, false, "")
	out, err := notifyVolumeOsd(-1, true, "")
	if err != nil {
		log.Println(string(out))
		os.Exit(1)
	}

	fmt.Println()
}
