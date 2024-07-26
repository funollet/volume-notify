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
