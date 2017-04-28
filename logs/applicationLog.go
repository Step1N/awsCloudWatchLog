package logs

import (
	"fmt"
	cw "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type LogStream struct {
	LogGroupName        *string
	Descending          *bool
	Limit               *int64
	LogStreamNamePrefix *string
	NextToken           *string
	OrderBy             *string
}

type LogEvent struct {
	LogGroupName  *string // Required
	LogStreamName *string // Required
	NextToken     *string
	StartFromHead *bool
	Limit         *int64
	StartTime     *int64
	EndTime       *int64
}

type CloudWatchClient struct {
	cwLogs *cw.CloudWatchLogs
}

func NewCloudWatchClient(cLogs *cw.CloudWatchLogs) *CloudWatchClient {
	return &CloudWatchClient{cwLogs: cLogs}
}

func (cc *CloudWatchClient) FindLogStreamByLogEvent(lEvent *LogEvent) (*cw.GetLogEventsOutput, error) {
	des := true
	orderBy := "LastEventTime"
	in := &cw.DescribeLogStreamsInput{
		LogGroupName: lEvent.LogGroupName,
		Limit:        lEvent.Limit,
		NextToken:    lEvent.NextToken,
		Descending:   &des,
		OrderBy:      &orderBy,
	}

	res, err := cc.cwLogs.DescribeLogStreams(in)
	if err != nil {
		panic(err)
	}

	stream := *res.LogStreams[0].LogStreamName
	fmt.Println("stream Name: ", stream)
	eventInput := &cw.GetLogEventsInput{
		LogGroupName:  lEvent.LogGroupName,
		LogStreamName: &stream,
		NextToken:     lEvent.NextToken,
		Limit:         lEvent.Limit,
	}

	lres, err := cc.cwLogs.GetLogEvents(eventInput)
	fmt.Println(lres)
	return lres, err
}

func (cc *CloudWatchClient) FindLogStream(lEvent *LogStream) (*cw.DescribeLogStreamsOutput, error) {
	in := &cw.DescribeLogStreamsInput{
		LogGroupName:        lEvent.LogGroupName,
		Descending:          lEvent.Descending,
		Limit:               lEvent.Limit,
		LogStreamNamePrefix: lEvent.LogStreamNamePrefix,
		NextToken:           lEvent.NextToken,
		OrderBy:             lEvent.OrderBy,
	}

	res, err := cc.cwLogs.DescribeLogStreams(in)
	fmt.Print(res)

	return res, err
}

func (cc *CloudWatchClient) FindLogEvent(lEvent *LogEvent) (*cw.GetLogEventsOutput, error) {
	in := &cw.GetLogEventsInput{
		LogGroupName:  lEvent.LogGroupName,
		LogStreamName: lEvent.LogStreamName,
		NextToken:     lEvent.NextToken,
		StartFromHead: lEvent.StartFromHead,
		Limit:         lEvent.Limit,
		StartTime:     lEvent.StartTime,
		EndTime:       lEvent.EndTime,
	}

	res, err := cc.cwLogs.GetLogEvents(in)
	fmt.Print(res)
	return res, err
}
