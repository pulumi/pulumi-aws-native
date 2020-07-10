package get

import (
	"context"

	_ "github.com/aws/aws-sdk-go/service/iot1clickprojects"
)

func (g *Getter) getIoT1ClickDeviceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// resp, err := g.iot1clickprojects.
	var err error
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"DeviceId": nil,
		"Enabled": nil,
	}
	return attrs, nil
}
