package jobrunner

import (
	"fmt"
)

func setupPythonVEnv() error {
	cmdStdOut, err := runCommand("python", []string{"-m", "venv", "."})
	if err != nil {
		outErr := fmt.Errorf("Failed to create python venv: %v. Output: %v", err, string(cmdStdOut))
		return outErr
	}

	// fs := fileaccess.FSAccess{}
	// d, _ := os.Getwd()
	// l, _ := fs.ListObjects("", d)
	// fmt.Printf("%v\n", strings.Join(l, "\n"))

	return nil
}
