package volume

import (
	"funollet/volume-notify/internal/shell"
	"strconv"
	"strings"
)

func processGet(text string) (volume string) {
  // Example input to parse:
  //
  // volume: Volume: front-left: 30146 /  46% / -20.24 dB,   front-right: 30146 /  46% / -20.24 dB
  //         balance 0.00

	lines := strings.Split(text, "\n")[0]
	fields := strings.Split(lines, "/")
  if len(fields) < 2 {
    return ""
  }
	value := strings.ReplaceAll(fields[1], "%", "")
	return strings.TrimSpace(value)
}

func Get() (volume int, err error){
  runner := shell.Shell{
    Command: "pactl get-sink-volume @DEFAULT_SINK@",
    Process: processGet,
  }
  out, err := runner.Do()
	if err != nil {
		return 0, err
	}
  volume, err = strconv.Atoi(out)
  return volume, err
}
