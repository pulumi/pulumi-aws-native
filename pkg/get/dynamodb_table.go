package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (g *Getter) getDynamoDBTableAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.dynamodb.DescribeTableWithContext(ctx, &dynamodb.DescribeTableInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"StreamArn": nil,
	}
	return attrs, nil
}
