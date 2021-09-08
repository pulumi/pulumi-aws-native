package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func (g *Getter) getLambdaFunctionAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.lambda.GetFunctionWithContext(ctx, &lambda.GetFunctionInput{
		FunctionName: aws.String(physicalResourceID),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "ResourceNotFoundException" {
			return nil, NewNotFoundError(err)
		}
		return nil, err
	}
	function := resp.Configuration

	arn := aws.StringValue(function.FunctionArn)

	attrs := map[string]interface{}{
		"Arn": arn,
	}
	return attrs, nil
}
