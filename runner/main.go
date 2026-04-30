package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
