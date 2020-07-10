package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/inspector"
)

func (g *Getter) getInspectorAssessmentTargetAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.inspector.DescribeAssessmentTargetsWithContext(ctx, &inspector.DescribeAssessmentTargetsInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
	}
	return attrs, nil
}
