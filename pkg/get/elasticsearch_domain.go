package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/elasticsearchservice"
)

func (g *Getter) getElasticsearchDomainAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.elasticsearchservice.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"DomainArn": nil,
		"DomainEndpoint": nil,
	}
	return attrs, nil
}
