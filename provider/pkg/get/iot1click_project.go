package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
)

func (g *Getter) getIoT1ClickProjectAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.iot1clickprojects.DescribeProjectWithContext(ctx, &iot1clickprojects.DescribeProjectInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"ProjectName": nil,
	}
	return attrs, nil
}
