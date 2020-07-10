package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ram"
)

func (g *Getter) getRAMResourceShareAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.ram.GetResourceSharesWithContext(ctx, &ram.GetResourceSharesInput{
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
