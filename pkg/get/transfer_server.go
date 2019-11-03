package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/transfer"
)

func (g *Getter) getTransferServerAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.transfer.DescribeServerWithContext(ctx, &transfer.DescribeServerInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"ServerId": nil,
	}
	return attrs, nil
}
