package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ssm"
)

func (g *Getter) getSSMParameterAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ssm.DescribeParametersWithContext(ctx, &ssm.DescribeParametersInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Type": nil,
		"Value": nil,
	}
	return attrs, nil
}
