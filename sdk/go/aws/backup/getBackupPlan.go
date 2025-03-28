// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Backup::BackupPlan
func LookupBackupPlan(ctx *pulumi.Context, args *LookupBackupPlanArgs, opts ...pulumi.InvokeOption) (*LookupBackupPlanResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupPlanResult
	err := ctx.Invoke("aws-native:backup:getBackupPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupPlanArgs struct {
	// Uniquely identifies a backup plan.
	BackupPlanId string `pulumi:"backupPlanId"`
}

type LookupBackupPlanResult struct {
	// Uniquely identifies the backup plan to be associated with the selection of resources.
	BackupPlan *BackupPlanResourceType `pulumi:"backupPlan"`
	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for example, `arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50` .
	BackupPlanArn *string `pulumi:"backupPlanArn"`
	// Uniquely identifies a backup plan.
	BackupPlanId *string `pulumi:"backupPlanId"`
	// The tags to assign to the backup plan.
	BackupPlanTags map[string]string `pulumi:"backupPlanTags"`
	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most 1,024 bytes long. Version Ids cannot be edited.
	VersionId *string `pulumi:"versionId"`
}

func LookupBackupPlanOutput(ctx *pulumi.Context, args LookupBackupPlanOutputArgs, opts ...pulumi.InvokeOption) LookupBackupPlanResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBackupPlanResultOutput, error) {
			args := v.(LookupBackupPlanArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:backup:getBackupPlan", args, LookupBackupPlanResultOutput{}, options).(LookupBackupPlanResultOutput), nil
		}).(LookupBackupPlanResultOutput)
}

type LookupBackupPlanOutputArgs struct {
	// Uniquely identifies a backup plan.
	BackupPlanId pulumi.StringInput `pulumi:"backupPlanId"`
}

func (LookupBackupPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPlanArgs)(nil)).Elem()
}

type LookupBackupPlanResultOutput struct{ *pulumi.OutputState }

func (LookupBackupPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPlanResult)(nil)).Elem()
}

func (o LookupBackupPlanResultOutput) ToLookupBackupPlanResultOutput() LookupBackupPlanResultOutput {
	return o
}

func (o LookupBackupPlanResultOutput) ToLookupBackupPlanResultOutputWithContext(ctx context.Context) LookupBackupPlanResultOutput {
	return o
}

// Uniquely identifies the backup plan to be associated with the selection of resources.
func (o LookupBackupPlanResultOutput) BackupPlan() BackupPlanResourceTypePtrOutput {
	return o.ApplyT(func(v LookupBackupPlanResult) *BackupPlanResourceType { return v.BackupPlan }).(BackupPlanResourceTypePtrOutput)
}

// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for example, `arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50` .
func (o LookupBackupPlanResultOutput) BackupPlanArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupPlanResult) *string { return v.BackupPlanArn }).(pulumi.StringPtrOutput)
}

// Uniquely identifies a backup plan.
func (o LookupBackupPlanResultOutput) BackupPlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupPlanResult) *string { return v.BackupPlanId }).(pulumi.StringPtrOutput)
}

// The tags to assign to the backup plan.
func (o LookupBackupPlanResultOutput) BackupPlanTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBackupPlanResult) map[string]string { return v.BackupPlanTags }).(pulumi.StringMapOutput)
}

// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most 1,024 bytes long. Version Ids cannot be edited.
func (o LookupBackupPlanResultOutput) VersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupPlanResult) *string { return v.VersionId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupPlanResultOutput{})
}
