package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/mediastore"
)

func (g *Getter) getMediaStoreContainerAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.mediastore.DescribeContainerWithContext(ctx, &mediastore.DescribeContainerInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Endpoint": nil,
	}
	return attrs, nil
}
