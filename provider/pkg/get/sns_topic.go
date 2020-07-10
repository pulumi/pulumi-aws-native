package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/sns"
)

func (g *Getter) getSNSTopicAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.sns.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"TopicName": nil,
	}
	return attrs, nil
}
