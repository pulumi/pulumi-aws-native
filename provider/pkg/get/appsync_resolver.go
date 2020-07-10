package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/appsync"
)

func (g *Getter) getAppSyncResolverAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.appsync.GetResolverWithContext(ctx, &appsync.GetResolverInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"FieldName": nil,
		"ResolverArn": nil,
		"TypeName": nil,
	}
	return attrs, nil
}
