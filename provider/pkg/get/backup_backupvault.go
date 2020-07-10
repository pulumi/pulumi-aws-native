package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/backup"
)

func (g *Getter) getBackupBackupVaultAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.backup.DescribeBackupVaultWithContext(ctx, &backup.DescribeBackupVaultInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"BackupVaultArn": nil,
		"BackupVaultName": nil,
	}
	return attrs, nil
}
