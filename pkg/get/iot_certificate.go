package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/iot"
)

func (g *Getter) getIoTCertificateAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.iot.DescribeCertificateWithContext(ctx, &iot.DescribeCertificateInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
	}
	return attrs, nil
}
