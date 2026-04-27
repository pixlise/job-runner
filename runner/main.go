package main

import (
	"fmt"
	"log"

	jobrunner "github.com/pixlise/core/v4/api/job/runner"
)

func main() {
	fmt.Println("PIXLISE Job Runner Starting...")
	err := jobrunner.RunJob(true)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("PIXLISE Job Runner Completed")
}
