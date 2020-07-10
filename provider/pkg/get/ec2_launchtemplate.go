package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (g *Getter) getEC2LaunchTemplateAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ec2.DescribeLaunchTemplatesWithContext(ctx, &ec2.DescribeLaunchTemplatesInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"DefaultVersionNumber": nil,
		"LatestVersionNumber": nil,
	}
	return attrs, nil
}
