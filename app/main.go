package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"

	l "awsCloudWatchLog/logs"
)

func main() {
	apiKey := ""
	secreteKey := ""
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(apiKey, secreteKey, ""),
	}))
	svc := cloudwatchlogs.New(sess)

	cc := l.NewCloudWatchClient(svc)
	param := &l.LogEvent{
		LogGroupName:  aws.String("API-Gateway-Execution-Logs_01dz2qdtre/cAuth"),
		LogStreamName: aws.String("98f13708210194c475687be6106a3b84"),
		StartFromHead: aws.Bool(true),
		Limit:         aws.Int64(10),
	}

	res, err := cc.FindLogStreamByLogEvent(param)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
