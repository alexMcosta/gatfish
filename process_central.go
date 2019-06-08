package main

import (
	"github.com/alexmcosta/gatfish/pkg/cloud"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Process Central uses channels and go routines to run every resources processes concurrently
// It is the great Orchestrator of the cloud package
func processCentral(tags *Tags) string {

	// Initiate sessions
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	EC2Session := ec2.New(sess)

	cloud.EBSCompliance(EC2Session, tags.EBS)

	return "Process Central Complete"
}

//TODO
//Helper that processes the ALL configuration
//Helper that processes all the different product sessions
//Have Process Central return an error instead of a string
