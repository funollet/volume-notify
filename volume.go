package main

import (
	"errors"
	"fmt"
	"log"
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

func getVolume() (volume int, err error) {
	out, err := run("pactl get-sink-volume @DEFAULT_SINK@")
	if err != nil {
		log.Println("unable to determine volume")
		return 0, err
	}
	return processGetVolume(out)
}

func processGetMute(s string) (bool, error) {
	return s == "Mute: yes\n", nil
}

func getMute() (muted bool, err error) {
	out, err := run("pactl get-sink-mute @DEFAULT_SINK@")
	if err != nil {
		return false, errors.New("unable to determine muted status")
	}
	return processGetMute(out)
}

func setVolume(change string) error {
	cmd := fmt.Sprintf("pactl set-sink-volume @DEFAULT_SINK@ %s", change)
	_, err := run(cmd)
	if err != nil {
		log.Println("unable to change volume")
		return err
	}
	return nil
}

func toggleMute() error {
	_, err := run("pactl set-sink-mute @DEFAULT_SINK@ toggle")
	if err != nil {
		log.Println("unable to change mute status")
		return err
	}
	return nil
}
