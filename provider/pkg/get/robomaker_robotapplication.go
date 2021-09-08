package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/robomaker"
)

func (g *Getter) getRoboMakerRobotApplicationAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.robomaker.DescribeRobotApplicationWithContext(ctx, &robomaker.DescribeRobotApplicationInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"CurrentRevisionId": nil,
	}
	return attrs, nil
}
