package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/appmesh"
)

func (g *Getter) getAppMeshVirtualNodeAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.appmesh.DescribeVirtualNodeWithContext(ctx, &appmesh.DescribeVirtualNodeInput{
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
		"VirtualNodeName": nil,
	}
	return attrs, nil
}
