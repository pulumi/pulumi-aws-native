package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/greengrass"
)

func (g *Getter) getGreengrassSubscriptionDefinitionAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.greengrass.GetSubscriptionDefinitionWithContext(ctx, &greengrass.GetSubscriptionDefinitionInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"Id": nil,
		"LatestVersionArn": nil,
		"Name": nil,
	}
	return attrs, nil
}
