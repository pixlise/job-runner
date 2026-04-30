package jobrunner

import (
	"fmt"
	"os/exec"
	"strings"
)

var testMode bool

func runCommand(command string, args []string) (string, error) {
	if testMode {
		fmt.Printf("Run Command: %v [%v]\n", command, strings.Join(args, ","))
		return "", nil
	}

	cmd := exec.Command(command, args...)
	cmdStdOut, err := cmd.CombinedOutput()
	return string(cmdStdOut), err
}
