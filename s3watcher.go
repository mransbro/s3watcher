//This app will check if an s3 bucket has had a new file added in a given period

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Bucket name required\nUsage: %s bucket_name\n",
			os.Args[0])
		os.Exit(1)
	}
	bucket := os.Args[1]
	hoursold := os.Args[2]
	i, _ := strconv.Atoi(hoursold)

	datenow := time.Now()
	fileage := datenow.Add(-time.Hour * time.Duration(i))

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-1"))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	svc := s3.NewFromConfig(cfg)

	resp, err := svc.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	for _, key := range resp.Contents {
		if key.LastModified.After(fileage) {
			fmt.Println(*key.Key, *key.LastModified)
		}
	}

	if err != nil {
		log.Fatalf("failed to list objects in %v", bucket)
	}
}
