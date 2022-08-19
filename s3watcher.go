//This app will check if an s3 bucket has had a new file added in a given period

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Bucket name required\nUsage: %s bucket_name\n",
			os.Args[0])
		os.Exit(1)
	}

	bucket := os.Args[1]

	// Initialize a session in eu-west-1 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-1"))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	svc := s3.NewFromConfig(cfg)

	resp, err := svc.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	for _, key := range resp.Contents {
		fmt.Println(*key.Key)
	}

	if err != nil {
		log.Fatalf("failed to list objects in %v", bucket)
	}
}
