package get

import (
	"context"
	"strings"

	"github.com/pkg/errors"
)

func (g *Getter) getECSServiceAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	// TODO(pdg): a proper implementation of this that can detect the deletion of a service requires the ARN of
	// the cluster in which the service is running. That requires the ECS service's properties in addition to its
	// physical resource ID. This is possible; it just requires some plumbing. For now, parse the service name out
	// of the ARN.

	components := strings.SplitN(physicalResourceID, ":service/", 2)
	if len(components) != 2 {
		return nil, NewNotFoundError(errors.Errorf("could not find ECS service %v", physicalResourceID))
	}
	name := components[1]

	attrs := map[string]interface{}{
		"Name": name,
	}
	return attrs, nil
}
