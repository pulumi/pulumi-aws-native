package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/greengrass"
)

func (g *Getter) getGreengrassResourceDefinitionAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.greengrass.GetResourceDefinitionWithContext(ctx, &greengrass.GetResourceDefinitionInput{
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
