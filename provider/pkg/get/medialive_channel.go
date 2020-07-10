package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/medialive"
)

func (g *Getter) getMediaLiveChannelAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.medialive.DescribeChannelWithContext(ctx, &medialive.DescribeChannelInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"Inputs": nil,
	}
	return attrs, nil
}
