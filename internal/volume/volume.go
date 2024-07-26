package volume

import (
	"errors"
	"funollet/volume-notify/internal/shell"
	"strconv"
	"strings"
)

func processGetVolume(text string) (volume int, err error) {
	lines := strings.Split(strings.TrimSpace(text), "\n")
	firstLine := lines[0]
	fields := strings.Split(firstLine, "/")
	if len(fields) < 2 {
		return 0, errors.New("invalid commmand output")
	}
	value := strings.ReplaceAll(fields[1], "%", "")
	return strconv.Atoi(strings.TrimSpace(value))
}

func GetVolume() (volume int, err error) {
	out, err := shell.Run("pactl get-sink-volume @DEFAULT_SINK@")
	if err != nil {
		return 0, errors.New("unable to determine volume")
	}
	return processGetVolume(out)
}

func processGetMute(s string) (bool, error) {
	return s == "Mute: yes\n", nil
}

func GetMute() (muted bool, err error) {
	out, err := shell.Run("pactl get-sink-mute @DEFAULT_SINK@")
	if err != nil {
		return false, errors.New("unable to determine muted status")
	}
	return processGetMute(out)
}
