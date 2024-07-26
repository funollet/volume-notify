package volume

import (
	"errors"
	"funollet/volume-notify/internal/shell"
	"strings"
)

func processGetVolume(text string) (volume string, err error) {
	lines := strings.Split(strings.TrimSpace(text), "\n")
	firstLine := lines[0]
	fields := strings.Split(firstLine, "/")
	if len(fields) < 2 {
		return "", errors.New("invalid commmand output")
	}
	value := strings.ReplaceAll(fields[1], "%", "")
	return strings.TrimSpace(value), nil
}

func GetVolume() (volume int, err error) {
	runner := shell.Shell{
		Command: "pactl get-sink-volume @DEFAULT_SINK@",
		Process: processGetVolume,
	}
	volume, err = runner.DoInt()
	return volume, err
}
