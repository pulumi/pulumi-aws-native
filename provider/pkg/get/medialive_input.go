package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/medialive"
)

func (g *Getter) getMediaLiveInputAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.medialive.DescribeInputWithContext(ctx, &medialive.DescribeInputInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"Destinations": nil,
		"Sources": nil,
	}
	return attrs, nil
}
