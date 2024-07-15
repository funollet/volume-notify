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

type shell struct {
  command string
  process func(string) (string)
}

func (sh shell) do() (result string, err error) {
  out, err := run(sh.command)
  if err != nil {
    return "", err
  }
  result = sh.process(out)
  return result, err
}

func processGetVolume(text string) (volume string) {
	lines := strings.Split(text, "\n")[0]
	fields := strings.Split(lines, "/")
  if len(fields) < 2 {
    return ""
  }
	value := strings.ReplaceAll(fields[1], "%", "")
	return strings.TrimSpace(value)
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
  getVolume := shell{
    command: "pactl get-sink-volume @DEFAULT_SINK@",
    process: processGetVolume,
  }
  volumeStr, err := getVolume.do()
	if err != nil {
		os.Exit(1)
	}

  var volume int
  volume, err = strconv.Atoi(volumeStr)
	out, err := notifyVolumeOsd(volume, false, "")
	// out, err := notifyVolumeOsd(-1, true, "")
	if err != nil {
		log.Println(out)
		os.Exit(1)
	}


	fmt.Println()
}
