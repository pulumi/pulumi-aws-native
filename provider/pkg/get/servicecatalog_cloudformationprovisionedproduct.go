package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/servicecatalog"
)

func (g *Getter) getServiceCatalogCloudFormationProvisionedProductAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.servicecatalog.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"CloudformationStackArn": nil,
		"RecordId": nil,
	}
	return attrs, nil
}
