package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (g *Getter) getDynamoDBTableAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.dynamodb.DescribeTableWithContext(ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String(physicalResourceID),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == dynamodb.ErrCodeResourceNotFoundException {
			return nil, NewNotFoundError(err)
		}
		return nil, err
	}
	table := resp.Table

	arn := aws.StringValue(table.TableArn)
	streamArn := aws.StringValue(table.LatestStreamArn)

	attrs := map[string]interface{}{
		"Arn":       arn,
		"StreamArn": streamArn,
	}
	return attrs, nil
}
