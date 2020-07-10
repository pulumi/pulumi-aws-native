package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/opsworkscm"
)

func (g *Getter) getOpsWorksCMServerAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.opsworkscm.DescribeServersWithContext(ctx, &opsworkscm.DescribeServersInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"Endpoint": nil,
	}
	return attrs, nil
}
