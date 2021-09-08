package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/servicecatalog"
)

func (g *Getter) getServiceCatalogPortfolioAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.servicecatalog.DescribePortfolioWithContext(ctx, &servicecatalog.DescribePortfolioInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"PortfolioName": nil,
	}
	return attrs, nil
}
