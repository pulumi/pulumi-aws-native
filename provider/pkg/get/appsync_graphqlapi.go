package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/appsync"
)

func (g *Getter) getAppSyncGraphQLApiAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.appsync.GetGraphqlApiWithContext(ctx, &appsync.GetGraphqlApiInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ApiId": nil,
		"Arn": nil,
		"GraphQLUrl": nil,
	}
	return attrs, nil
}
