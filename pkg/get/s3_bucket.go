package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/s3"
)

func (g *Getter) getS3BucketAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.s3.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"DomainName": nil,
		"DualStackDomainName": nil,
		"RegionalDomainName": nil,
		"WebsiteURL": nil,
	}
	return attrs, nil
}
