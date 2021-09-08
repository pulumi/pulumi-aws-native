package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/pinpoint"
)

func (g *Getter) getPinpointCampaignAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.pinpoint.GetCampaignWithContext(ctx, &pinpoint.GetCampaignInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"CampaignId": nil,
	}
	return attrs, nil
}
