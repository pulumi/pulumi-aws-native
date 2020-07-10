package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/backup"
)

func (g *Getter) getBackupBackupSelectionAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.backup.GetBackupSelectionWithContext(ctx, &backup.GetBackupSelectionInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"BackupPlanId": nil,
		"SelectionId": nil,
	}
	return attrs, nil
}
