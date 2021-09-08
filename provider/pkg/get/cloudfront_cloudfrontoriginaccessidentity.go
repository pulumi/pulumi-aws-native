package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/cloudfront"
)

func (g *Getter) getCloudFrontCloudFrontOriginAccessIdentityAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.cloudfront.GetCloudFrontOriginAccessIdentityWithContext(ctx, &cloudfront.GetCloudFrontOriginAccessIdentityInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"S3CanonicalUserId": nil,
	}
	return attrs, nil
}
