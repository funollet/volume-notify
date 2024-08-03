package main

import (
	"testing"
)

func TestProcessGetVolumeHappyPath(t *testing.T) {
	input := `volume: Volume: front-left: 30146 /  46% / -20.24 dB,   front-right: 30146 /  46% / -20.24 dB
        balance 0.00
`
	if volume, _ := processGetVolume(input); volume != 46 {
		t.Error("wrong value")
	}
}

func TestProcessGetVolumeDealsWithWhitespace(t *testing.T) {
	input := `
    volume: Volume: front-left: 30146 /  46% / -20.24 dB,   front-right: 30146 /  46% / -20.24 dB
        balance 0.00

`
	if volume, _ := processGetVolume(input); volume != 46 {
		t.Error("fails to deal with whitespace")
	}
}

func TestProcessGetVolumeFailsWithBadInput(t *testing.T) {
	input := "zasca"
	if _, err := processGetVolume(input); err == nil {
		t.Error("fails to return an error when the text can't be parsed")
	}
}

func TestProcessGetMute(t *testing.T) {
	if r, _ := processGetMute("Mute: yes\n"); r != true {
		t.Error("fails to detect when it's muted")
	}
	if r, _ := processGetMute("Mute: no"); r != false {
		t.Error("fails to detect when it's unmuted")
	}
	if r, _ := processGetMute("Mute: maybe"); r != false {
		t.Error("should return false on unexpected results")
	}

}
