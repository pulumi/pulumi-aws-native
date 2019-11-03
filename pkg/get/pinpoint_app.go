package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/pinpoint"
)

func (g *Getter) getPinpointAppAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.pinpoint.GetAppWithContext(ctx, &pinpoint.GetAppInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
	}
	return attrs, nil
}
