package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/backup"
)

func (g *Getter) getBackupBackupPlanAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.backup.GetBackupPlanWithContext(ctx, &backup.GetBackupPlanInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"BackupPlanArn": nil,
		"BackupPlanId": nil,
		"VersionId": nil,
	}
	return attrs, nil
}
