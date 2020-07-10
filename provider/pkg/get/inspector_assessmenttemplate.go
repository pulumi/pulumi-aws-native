package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/inspector"
)

func (g *Getter) getInspectorAssessmentTemplateAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.inspector.DescribeAssessmentTemplatesWithContext(ctx, &inspector.DescribeAssessmentTemplatesInput{
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
