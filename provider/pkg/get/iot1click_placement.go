package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
)

func (g *Getter) getIoT1ClickPlacementAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.iot1clickprojects.DescribePlacementWithContext(ctx, &iot1clickprojects.DescribePlacementInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"PlacementName": nil,
		"ProjectName": nil,
	}
	return attrs, nil
}
