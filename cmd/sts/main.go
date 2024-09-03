package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {

	development := os.Getenv("DEVELOP") == "true"

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	options := sts.Options{}

	if development {
		fmt.Printf("#####\n\tDEVELOPMENT: %v\n#####\n", development)
		options.BaseEndpoint = aws.String("http://localhost:4566")
	}

	s := sts.NewFromConfig(cfg, func(o *sts.Options) {
		o.BaseEndpoint = options.BaseEndpoint
	})

	identity, err := s.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})

	fmt.Printf("\n\tUser: %v,\n\tAccount: %v\n\tArn: %v\n", *identity.UserId, *identity.Account, *identity.Arn)

}
