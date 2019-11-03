package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/kms"
)

func (g *Getter) getKMSKeyAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.kms.DescribeKeyWithContext(ctx, &kms.DescribeKeyInput{
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
