package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/sfn"
)

func (g *Getter) getStepFunctionsStateMachineAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.sfn.DescribeStateMachineWithContext(ctx, &sfn.DescribeStateMachineInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Name": nil,
	}
	return attrs, nil
}
