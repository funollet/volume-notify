package shell

import (
	"log"
	"os/exec"
	"strconv"
)

func Run(cmd string) (result string, err error) {
	var out []byte
	out, err = exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Println(string(out))
	}
	return string(out), err
}

type Shell struct {
	Command string
	Process func(string) (string, error)
}

func (sh Shell) Do() (result string, err error) {
	out, err := Run(sh.Command)
	if err != nil {
		return "", err
	}
	return sh.Process(out)
}

func (sh Shell) DoInt() (result int, err error) {
	cmdOut, err := Run(sh.Command)
	if err != nil {
		return 0, err
	}
	out, err := sh.Process(cmdOut)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(out)
}
