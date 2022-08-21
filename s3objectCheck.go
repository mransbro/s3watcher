package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type envs struct {
	Bucket         string `env:"BUCKET"`
	FileAgeInHours int    `env:"FILE_AGE_IN_HOURS" envDefault:"24"`
	Region         string `env:"AWS_REGION"`
	FilePath       string `env:"FILE_PATH" envDefault:"/"`
	ObjectSizeMB   int    `env:"OBJECT_SIZE_MB" envDefault:"1"`
}

func main() {

	envconf := envs{}
	FileAgeInHours := time.Now().Add(-time.Hour * time.Duration(envconf.FileAgeInHours))

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(envconf.Region))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	svc := s3.NewFromConfig(cfg)

	resp, err := svc.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{Bucket: aws.String(envconf.Bucket)})
	for _, object := range resp.Contents {
		if object.LastModified.After(FileAgeInHours) && object.Size > int64(envconf.ObjectSizeMB) {
			fmt.Println(*object.Key, *object.LastModified, object.Size)
		}
	}

	if err != nil {
		log.Fatalf("failed to list objects in %v", envconf.Bucket)
	}
}
