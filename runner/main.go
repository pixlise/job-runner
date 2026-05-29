package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pixlise/core/v4/api/job/jobrunner"
	"github.com/pixlise/core/v4/core/awsutil"
	"github.com/pixlise/core/v4/core/fileaccess"
)

func main() {
	bucket := os.Getenv(jobrunner.EnvBucketName)
	jobPath := os.Getenv(jobrunner.EnvPathName)
	nodeIdxStr := os.Getenv(jobrunner.EnvNodeIndexName)

	nodeIdx, err := strconv.Atoi(nodeIdxStr)
	if err != nil {
		log.Fatalf("Failed to read %v: %v. Error %v", jobrunner.EnvNodeIndexName, nodeIdxStr, err)
	}

	fmt.Println("PIXLISE Job Runner Starting, job path: s3://%v/%v, node index: %v...", bucket, jobPath, nodeIdx)

	// Get a session for the bucket region
	sess, err := awsutil.GetSession()
	if err != nil {
		log.Fatalf("Failed to create AWS session. Error: %v", err)
	}

	s3svc, err := awsutil.GetS3(sess)
	if err != nil {
		log.Fatalf("Failed to create AWS S3 service. Error: %v", err)
	}

	remoteFS := fileaccess.MakeS3Access(s3svc)

	err = jobrunner.RunJob(bucket, jobPath, uint(nodeIdx), remoteFS)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("PIXLISE Job Runner Completed")
}
