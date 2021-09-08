package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/appsync"
)

func (g *Getter) getAppSyncDataSourceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.appsync.GetDataSourceWithContext(ctx, &appsync.GetDataSourceInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"DataSourceArn": nil,
		"Name": nil,
	}
	return attrs, nil
}
