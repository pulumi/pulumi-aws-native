package provider

import (
	"context"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func (p *cfnProvider) getAccountID(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	return resource.NewPropertyMapFromMap(map[string]interface{}{
		"accountId": p.accountID,
	}), nil
}

func (p *cfnProvider) getRegion(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	return resource.NewPropertyMapFromMap(map[string]interface{}{
		"region": p.region,
	}), nil
}
