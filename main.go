package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

var (
	durationSeconds int
	serialNumber    string
)

func init() {
	flag.IntVar(&durationSeconds, "d", 900, "DurationSeconds: The duration, in seconds, that the credentials should remain valid. Acceptable durations for IAM user sessions range from 900 seconds (15 minutes) to 129600 seconds (36 hours), with 43200 seconds (12 hours) as the default. Sessions for AWS account owners are restricted to a maximum of 3600 seconds (one hour). If the duration is longer than one hour, the session for AWS account owners defaults to one hour.")
	flag.StringVar(&serialNumber, "s", "", "SerialNumber: The identification number of the MFA device that is associated with the IAM user who is making the GetSessionToken call. Specify this value if the IAM user has a policy that requires MFA authentication. The value is either the serial number for a hardware device (such as GAHT12345678) or an Amazon Resource Name (ARN) for a virtual device (such as arn:aws:iam::123456789012:mfa/user). You can find the device for an IAM user by going to the AWS Management Console and viewing the user's security credentials.")
}

func main() {
	flag.Parse()

	fmt.Printf("MFA Token Code: ")
	tokenCode, err := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("\n")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	tokenCode = strings.TrimSpace(tokenCode)

	service := sts.New(session.New())

	request := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(int64(durationSeconds)),
		SerialNumber:    aws.String(serialNumber),
		TokenCode:       aws.String(tokenCode),
	}
	fmt.Printf("Request: %v\n\n", request)

	response, err := service.GetSessionToken(request)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Response: %v\n\n", response)

	fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", *response.Credentials.AccessKeyId)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", *response.Credentials.SecretAccessKey)
	fmt.Printf("export AWS_SESSION_TOKEN=%s\n", *response.Credentials.SessionToken)
}
