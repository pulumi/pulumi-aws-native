package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/appsync"
)

func (g *Getter) getAppSyncFunctionConfigurationAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.appsync.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"DataSourceName": nil,
		"FunctionArn": nil,
		"FunctionId": nil,
		"Name": nil,
	}
	return attrs, nil
}
