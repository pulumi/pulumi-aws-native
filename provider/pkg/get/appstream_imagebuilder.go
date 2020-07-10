package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/appstream"
)

func (g *Getter) getAppStreamImageBuilderAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.appstream.DescribeImageBuildersWithContext(ctx, &appstream.DescribeImageBuildersInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"StreamingUrl": nil,
	}
	return attrs, nil
}
