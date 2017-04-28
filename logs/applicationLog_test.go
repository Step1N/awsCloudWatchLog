package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func TestFindLogStreamByLogEvent(t *testing.T) {
	apiKey := "--AWS Key---"
	secreteKey := "--AWS Secrete---"

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(apiKey, secreteKey, ""),
	}))
	svc := cloudwatchlogs.New(sess)

	cc := NewCloudWatchClient(svc)
	param := &LogEvent{
		LogGroupName:  aws.String("API-Gateway-Execution-Logs_01dz2qdtre/cAuth"),
		LogStreamName: aws.String("98f13708210194c475687be6106a3b84"),
		StartFromHead: aws.Bool(true),
		Limit:         aws.Int64(10),
	}

	res, err := cc.FindLogStreamByLogEvent(param)
	require.Nil(t, err)
	require.NotNil(t, res)

	assert.NotEmpty(t, res, "No event found")
}

func TestFindLogStream(t *testing.T) {
	apiKey := "--AWS Key---"
	secreteKey := "--AWS Secrete---"

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(apiKey, secreteKey, ""),
	}))
	svc := cloudwatchlogs.New(sess)

	cc := NewCloudWatchClient(svc)
	param := &LogStream{
		LogGroupName: aws.String("API-Gateway-Execution-Logs_01dz2qdtre/cAuth"),
		Limit:        aws.Int64(10),
	}

	res, err := cc.FindLogStream(param)
	require.Nil(t, err)
	require.NotNil(t, res)

	assert.NotEmpty(t, res, "No event found")
}

func TestFindLogEvent(t *testing.T) {
	apiKey := "--AWS Key---"
	secreteKey := "--AWS Secrete---"

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(apiKey, secreteKey, ""),
	}))
	svc := cloudwatchlogs.New(sess)

	cc := NewCloudWatchClient(svc)
	param := &LogEvent{
		LogGroupName:  aws.String("API-Gateway-Execution-Logs_01dz2qdtre/cAuth"),
		LogStreamName: aws.String("98f13708210194c475687be6106a3b84"),
		StartFromHead: aws.Bool(true),
		Limit:         aws.Int64(10),
	}

	res, err := cc.FindLogEvent(param)
	require.Nil(t, err)
	require.NotNil(t, res)

	assert.NotEmpty(t, res, "No event found")
}
