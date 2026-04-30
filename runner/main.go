package main

import (
	"fmt"
	"log"

	"github.com/pixlise/job-runner/jobrunner"
)

func main() {
	fmt.Println("PIXLISE Job Runner Starting...")

	err := jobrunner.RunJob(false)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("PIXLISE Job Runner Completed")
}
