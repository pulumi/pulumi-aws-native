package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/servicecatalog"
)

func (g *Getter) getServiceCatalogCloudFormationProductAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.servicecatalog.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ProductName": nil,
		"ProvisioningArtifactIds": nil,
		"ProvisioningArtifactNames": nil,
	}
	return attrs, nil
}
