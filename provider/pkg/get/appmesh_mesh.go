package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/appmesh"
)

func (g *Getter) getAppMeshMeshAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.appmesh.DescribeMeshWithContext(ctx, &appmesh.DescribeMeshInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"MeshName": nil,
		"Uid": nil,
	}
	return attrs, nil
}
