package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/amplify"
)

func (g *Getter) getAmplifyAppAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.amplify.GetAppWithContext(ctx, &amplify.GetAppInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"AppId": nil,
		"AppName": nil,
		"Arn": nil,
		"DefaultDomain": nil,
	}
	return attrs, nil
}
