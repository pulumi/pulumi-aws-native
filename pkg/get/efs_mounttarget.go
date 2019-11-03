package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/efs"
)

func (g *Getter) getEFSMountTargetAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.efs.DescribeMountTargetsWithContext(ctx, &efs.DescribeMountTargetsInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"IpAddress": nil,
	}
	return attrs, nil
}
