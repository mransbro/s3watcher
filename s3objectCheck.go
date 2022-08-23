package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/caarlos0/env"
)

type envs struct {
	Bucket         string `env:"BUCKET"`
	FileAgeInHours int    `env:"FILE_AGE_IN_HOURS"`
	Region         string `env:"AWS_REGION"`
	FilePath       string `env:"FILE_PATH"`
}

func main() {

	envconf := envs{}
	if err := env.Parse(&envconf); err != nil {
		log.Fatalf("%+v\n", err)
	}
	FileAgeInHours := time.Now().Add(-time.Hour * time.Duration(envconf.FileAgeInHours))

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(envconf.Region))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	svc := s3.NewFromConfig(cfg)

	resp, err := svc.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{Bucket: aws.String(envconf.Bucket), Prefix: aws.String(envconf.FilePath)})
	if err != nil {
		exitErrorf("Unable to list objects in, %b %p", envconf.Bucket, envconf.FilePath)
	}
	if len(resp.Contents) == 0 {
		log.Fatalln("No objects to list.")
	}
	for _, object := range resp.Contents {
		if object.LastModified.After(FileAgeInHours) {

			s := object.Size
			smb := float64(s) / (1 << 20)

			fmt.Printf("%v, %v, %vMB\n", *object.Key, *object.LastModified, math.Round(smb*100)/100)
		}
	}

	if err != nil {
		log.Fatalf("failed to list objects in %v", envconf.Bucket)
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
