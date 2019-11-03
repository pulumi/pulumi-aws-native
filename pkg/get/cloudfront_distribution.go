package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/cloudfront"
)

func (g *Getter) getCloudFrontDistributionAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.cloudfront.GetDistributionWithContext(ctx, &cloudfront.GetDistributionInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"DomainName": nil,
	}
	return attrs, nil
}
