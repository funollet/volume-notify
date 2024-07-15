package shell

import (
  "log"
  "os/exec"
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
  Process func(string) (string)
}

func (sh Shell) Do() (result string, err error) {
  out, err := Run(sh.Command)
  if err != nil {
    return "", err
  }
  result = sh.Process(out)
  return result, err
}
