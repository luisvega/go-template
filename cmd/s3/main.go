package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

var bucket string = "test-bucket-002"
var filename string = "./some_report.csv"

func main() {

	development := os.Getenv("DEVELOP") == "true"

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	options := s3.Options{}

	if development {
		fmt.Printf("#####\n\tDEVELOPMENT: %v\n#####\n", development)
		options.BaseEndpoint = aws.String("http://localhost:4572")
	}

	s := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = options.BaseEndpoint
	})

	// read file
	file, err := os.Open(filename)
	if err != nil {
		log.Panicf("failed to read file: %v\n", err)
	}

	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)

	file.Read(buffer)
	// \read file

	input := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String("/fake/path"),
		Body:          file,
		ContentLength: &size,
	}

	output, err := s.PutObject(context.TODO(), input)

	if err != nil {
		log.Panicf("failed to put object: %v\n", err)
	}

	log.Printf("putobject output: %v\n", output)
}
