package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func notifySend(args []string) ([]byte, error) {
	replaceId := "2593"
	timeout := 1000

	defaultArgs := make([]string, 0)
	defaultArgs = append(defaultArgs,
		"--app-name", "volume-osd",
		"--replace-id", replaceId,
		"--transient",
		"--expire-time", fmt.Sprint(timeout),
		"--urgency", "low",
	)

	allArgs := make([]string, 0)
	allArgs = append(defaultArgs, args...)
	allArgs = append(allArgs, " ")

	out, err := exec.Command("notify-send", allArgs...).CombinedOutput()
	return out, err
}

func notifySendVolume(percentage int, muted bool) ([]byte, error) {
	icon := "audio-volume-medium"
	if muted {
		icon = "audio-volume-muted"
	}

	args := make([]string, 0)
	args = append(args,
		"--hint", fmt.Sprintf("int:value:%d", percentage),
		"--icon", icon,
	)
	out, err := notifySend(args)
	return out, err
}

func main() {
	muted := true
	out, err := notifySendVolume(50, muted)
	if err != nil {
		log.Println(string(out))
		os.Exit(1)
	}

	fmt.Println()
}
