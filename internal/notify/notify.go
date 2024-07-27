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
    --app-name volume-notify \
    --replace-id 2593 \
    --transient \
    --expire-time %d \
    --urgency low \
    --icon %s \
  `, timeout, iconWithDefaults)

	if !muted {
		cmd += fmt.Sprintf("  --hint int:value:%d \\\n", percentage)
	}

	// notify-send raises an error if it receives an empty body
	cmd += "    ' '"

	return shell.Run(cmd)
}
