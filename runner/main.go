package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	jobrunner "github.com/pixlise/core/v4/api/job/runner"
)

func main() {
	fmt.Println("PIXLISE Job Runner Starting...")

	// Check if we need to install any libraries
	cfgStr := os.Getenv(jobrunner.JobConfigEnvVar)

	var err error
	fmt.Println("Installing required libraries...")
	if len(cfgStr) > 0 {
		cfg := &jobrunner.JobConfig{}
		err = json.Unmarshal([]byte(cfgStr), cfg)
		if err == nil {
			if strings.Contains(cfg.Command, "python") {
				err = installPythonLibs()
			} else if strings.Contains(cfg.Command, "lua") {
				err = installLuaLibs()
			}
		}
	}
	fmt.Println("Running job...")

	err = jobrunner.RunJob(true)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("PIXLISE Job Runner Completed")
}

func installPythonLibs() error {
	f, err := os.Stat("requirements.txt")
	if err == nil && !f.IsDir() && f.Name() == "requirements.txt" {
		// Run pip
		cmd := exec.Command("pip", "install", "r", "requirements.txt")

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to install python libraries:\n%v\n", string(out))
			return err
		}
	}

	// No requirements.txt found or it worked... no errors!
	return nil
}

func installLuaLibs() error {
	// If we're dealing with a rockspec file, treat it as such
	allargs := [][]string{}

	f, err := os.Stat("requirements.rockspec")
	if err == nil && !f.IsDir() && f.Name() == "requirements.rockspec" {
		allargs = append(allargs, []string{"luarocks-5.3", "install", "requirements.rockspec"})
	} else {
		// See if there's a lua-requirements.txt, we'll read it line-by-line in that case and install each
		b, err := os.ReadFile("lua-requirements.txt")
		if err == nil {
			lines := strings.Split(string(b), "\n")
			for _, line := range lines {
				allargs = append(allargs, []string{"luarocks-5.3", "install", line})
			}
		}
	}

	// Run all commands, return if an error happens
	for _, args := range allargs {
		fmt.Printf("Executing: %v\n", strings.Join(args, " "))
		cmd := exec.Command(args[0], args[1:]...)

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error while installing lua library [%v]: %v\n", strings.Join(args, ","), string(out))
			return err
		}
	}

	return nil
}
