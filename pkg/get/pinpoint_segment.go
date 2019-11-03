package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/pinpoint"
)

func (g *Getter) getPinpointSegmentAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.pinpoint.GetSegmentWithContext(ctx, &pinpoint.GetSegmentInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"SegmentId": nil,
	}
	return attrs, nil
}
