package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/emr"
)

func (g *Getter) getEMRClusterAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.emr.DescribeClusterWithContext(ctx, &emr.DescribeClusterInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"MasterPublicDNS": nil,
	}
	return attrs, nil
}
