package provider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func (p *cfnProvider) getAccountID(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	return resource.NewPropertyMapFromMap(map[string]interface{}{
		"accountId": p.accountID,
	}), nil
}

func (p *cfnProvider) getPartition(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	partition, ok := endpoints.PartitionForRegion(endpoints.DefaultPartitions(), p.region)
	if !ok {
		return nil, fmt.Errorf("no partition for region %v", p.region)
	}
	return resource.NewPropertyMapFromMap(map[string]interface{}{
		"partition": partition.ID(),
	}), nil
}

func (p *cfnProvider) getRegion(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	return resource.NewPropertyMapFromMap(map[string]interface{}{
		"region": p.region,
	}), nil
}

func (p *cfnProvider) getURLSuffix(ctx context.Context, inputs resource.PropertyMap) (resource.PropertyMap, error) {
	partition, ok := endpoints.PartitionForRegion(endpoints.DefaultPartitions(), p.region)
	if !ok {
		return nil, fmt.Errorf("no partition for region %v", p.region)
	}
	return resource.NewPropertyMapFromMap(map[string]interface{}{
		"urlSuffix": partition.DNSSuffix(),
	}), nil
}
