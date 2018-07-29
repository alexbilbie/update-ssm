package main

import (
	"os"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Expected two arguments, the first is the SSM key, the second is the new value")
		os.Exit(1)
	}

	key := os.Args[1]
	value := os.Args[2]

	s, _ := session.NewSession(&aws.Config{})
	client := ssm.New(s)

	_, err := client.PutParameter(&ssm.PutParameterInput{
		Name:      aws.String(key),
		Overwrite: aws.Bool(true),
		Type:      aws.String(ssm.ParameterTypeString),
		Value:     aws.String(value),
	})

	if err != nil {
		fmt.Printf("SSM PutParameter error: %s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
