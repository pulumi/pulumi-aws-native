package get

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

func (g *Getter) getCloudFrontDistributionAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.cloudfront.GetDistributionWithContext(ctx, &cloudfront.GetDistributionInput{
		Id: aws.String(physicalResourceID),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NoSuchDistribution" {
			return nil, NewNotFoundError(err)
		}
		return nil, err
	}

	attrs := map[string]interface{}{
		"DomainName": aws.StringValue(resp.Distribution.DomainName),
	}
	return attrs, nil
}
