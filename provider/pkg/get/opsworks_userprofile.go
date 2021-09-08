package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/opsworks"
)

func (g *Getter) getOpsWorksUserProfileAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.opsworks.DescribeUserProfilesWithContext(ctx, &opsworks.DescribeUserProfilesInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"SshUsername": nil,
	}
	return attrs, nil
}
