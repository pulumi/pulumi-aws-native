package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/pkg/errors"
)

func (g *Getter) getApiGatewayRestApiAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	_, err := g.apigateway.GetRestApiWithContext(ctx, &apigateway.GetRestApiInput{
		RestApiId: aws.String(physicalResourceID),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == apigateway.ErrCodeNotFoundException {
			return nil, NewNotFoundError(err)
		}
		return nil, err
	}

	rootResourceID := ""
	err = g.apigateway.GetResourcesPagesWithContext(ctx, &apigateway.GetResourcesInput{
		RestApiId: aws.String(physicalResourceID),
	}, func(resp *apigateway.GetResourcesOutput, _ bool) bool {
		for _, item := range resp.Items {
			if aws.StringValue(item.Path) == "/" {
				rootResourceID = aws.StringValue(item.Id)
				return false
			}
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	if rootResourceID == "" {
		return nil, NewNotFoundError(errors.Errorf("could not find root resource for REST API %v", physicalResourceID))
	}

	attrs := map[string]interface{}{
		"RootResourceId": rootResourceID,
	}
	return attrs, nil
}
