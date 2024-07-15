package notify

import (
  "fmt"
  "funollet/volume-notify/internal/shell"
)

func NotifyVolumeOsd(percentage int, muted bool, icon string) (out string, err error) {
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

	return shell.Run(cmd)
}

