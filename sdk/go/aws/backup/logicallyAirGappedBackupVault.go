// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Backup::LogicallyAirGappedBackupVault
type LogicallyAirGappedBackupVault struct {
	pulumi.CustomResourceState

	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Backup::LogicallyAirGappedBackupVault` for more information about the expected schema for this property.
	AccessPolicy     pulumi.AnyOutput                                             `pulumi:"accessPolicy"`
	BackupVaultArn   pulumi.StringOutput                                          `pulumi:"backupVaultArn"`
	BackupVaultName  pulumi.StringOutput                                          `pulumi:"backupVaultName"`
	BackupVaultTags  pulumi.StringMapOutput                                       `pulumi:"backupVaultTags"`
	EncryptionKeyArn pulumi.StringOutput                                          `pulumi:"encryptionKeyArn"`
	MaxRetentionDays pulumi.IntOutput                                             `pulumi:"maxRetentionDays"`
	MinRetentionDays pulumi.IntOutput                                             `pulumi:"minRetentionDays"`
	Notifications    LogicallyAirGappedBackupVaultNotificationObjectTypePtrOutput `pulumi:"notifications"`
	VaultState       pulumi.StringPtrOutput                                       `pulumi:"vaultState"`
	VaultType        pulumi.StringPtrOutput                                       `pulumi:"vaultType"`
}

// NewLogicallyAirGappedBackupVault registers a new resource with the given unique name, arguments, and options.
func NewLogicallyAirGappedBackupVault(ctx *pulumi.Context,
	name string, args *LogicallyAirGappedBackupVaultArgs, opts ...pulumi.ResourceOption) (*LogicallyAirGappedBackupVault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxRetentionDays == nil {
		return nil, errors.New("invalid value for required argument 'MaxRetentionDays'")
	}
	if args.MinRetentionDays == nil {
		return nil, errors.New("invalid value for required argument 'MinRetentionDays'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"backupVaultName",
		"maxRetentionDays",
		"minRetentionDays",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogicallyAirGappedBackupVault
	err := ctx.RegisterResource("aws-native:backup:LogicallyAirGappedBackupVault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogicallyAirGappedBackupVault gets an existing LogicallyAirGappedBackupVault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogicallyAirGappedBackupVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogicallyAirGappedBackupVaultState, opts ...pulumi.ResourceOption) (*LogicallyAirGappedBackupVault, error) {
	var resource LogicallyAirGappedBackupVault
	err := ctx.ReadResource("aws-native:backup:LogicallyAirGappedBackupVault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogicallyAirGappedBackupVault resources.
type logicallyAirGappedBackupVaultState struct {
}

type LogicallyAirGappedBackupVaultState struct {
}

func (LogicallyAirGappedBackupVaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*logicallyAirGappedBackupVaultState)(nil)).Elem()
}

type logicallyAirGappedBackupVaultArgs struct {
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Backup::LogicallyAirGappedBackupVault` for more information about the expected schema for this property.
	AccessPolicy     interface{}                                          `pulumi:"accessPolicy"`
	BackupVaultName  *string                                              `pulumi:"backupVaultName"`
	BackupVaultTags  map[string]string                                    `pulumi:"backupVaultTags"`
	MaxRetentionDays int                                                  `pulumi:"maxRetentionDays"`
	MinRetentionDays int                                                  `pulumi:"minRetentionDays"`
	Notifications    *LogicallyAirGappedBackupVaultNotificationObjectType `pulumi:"notifications"`
	VaultState       *string                                              `pulumi:"vaultState"`
	VaultType        *string                                              `pulumi:"vaultType"`
}

// The set of arguments for constructing a LogicallyAirGappedBackupVault resource.
type LogicallyAirGappedBackupVaultArgs struct {
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Backup::LogicallyAirGappedBackupVault` for more information about the expected schema for this property.
	AccessPolicy     pulumi.Input
	BackupVaultName  pulumi.StringPtrInput
	BackupVaultTags  pulumi.StringMapInput
	MaxRetentionDays pulumi.IntInput
	MinRetentionDays pulumi.IntInput
	Notifications    LogicallyAirGappedBackupVaultNotificationObjectTypePtrInput
	VaultState       pulumi.StringPtrInput
	VaultType        pulumi.StringPtrInput
}

func (LogicallyAirGappedBackupVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logicallyAirGappedBackupVaultArgs)(nil)).Elem()
}

type LogicallyAirGappedBackupVaultInput interface {
	pulumi.Input

	ToLogicallyAirGappedBackupVaultOutput() LogicallyAirGappedBackupVaultOutput
	ToLogicallyAirGappedBackupVaultOutputWithContext(ctx context.Context) LogicallyAirGappedBackupVaultOutput
}

func (*LogicallyAirGappedBackupVault) ElementType() reflect.Type {
	return reflect.TypeOf((**LogicallyAirGappedBackupVault)(nil)).Elem()
}

func (i *LogicallyAirGappedBackupVault) ToLogicallyAirGappedBackupVaultOutput() LogicallyAirGappedBackupVaultOutput {
	return i.ToLogicallyAirGappedBackupVaultOutputWithContext(context.Background())
}

func (i *LogicallyAirGappedBackupVault) ToLogicallyAirGappedBackupVaultOutputWithContext(ctx context.Context) LogicallyAirGappedBackupVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogicallyAirGappedBackupVaultOutput)
}

type LogicallyAirGappedBackupVaultOutput struct{ *pulumi.OutputState }

func (LogicallyAirGappedBackupVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogicallyAirGappedBackupVault)(nil)).Elem()
}

func (o LogicallyAirGappedBackupVaultOutput) ToLogicallyAirGappedBackupVaultOutput() LogicallyAirGappedBackupVaultOutput {
	return o
}

func (o LogicallyAirGappedBackupVaultOutput) ToLogicallyAirGappedBackupVaultOutputWithContext(ctx context.Context) LogicallyAirGappedBackupVaultOutput {
	return o
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Backup::LogicallyAirGappedBackupVault` for more information about the expected schema for this property.
func (o LogicallyAirGappedBackupVaultOutput) AccessPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *LogicallyAirGappedBackupVault) pulumi.AnyOutput { return v.AccessPolicy }).(pulumi.AnyOutput)
}

func (o LogicallyAirGappedBackupVaultOutput) BackupVaultArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LogicallyAirGappedBackupVault) pulumi.StringOutput { return v.BackupVaultArn }).(pulumi.StringOutput)
}

func (o LogicallyAirGappedBackupVaultOutput) BackupVaultName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogicallyAirGappedBackupVault) pulumi.StringOutput { return v.BackupVaultName }).(pulumi.StringOutput)
}

func (o LogicallyAirGappedBackupVaultOutput) BackupVaultTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LogicallyAirGappedBackupVault) pulumi.StringMapOutput { return v.BackupVaultTags }).(pulumi.StringMapOutput)
}

func (o LogicallyAirGappedBackupVaultOutput) EncryptionKeyArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LogicallyAirGappedBackupVault) pulumi.StringOutput { return v.EncryptionKeyArn }).(pulumi.StringOutput)
}

func (o LogicallyAirGappedBackupVaultOutput) MaxRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *LogicallyAirGappedBackupVault) pulumi.IntOutput { return v.MaxRetentionDays }).(pulumi.IntOutput)
}

func (o LogicallyAirGappedBackupVaultOutput) MinRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *LogicallyAirGappedBackupVault) pulumi.IntOutput { return v.MinRetentionDays }).(pulumi.IntOutput)
}

func (o LogicallyAirGappedBackupVaultOutput) Notifications() LogicallyAirGappedBackupVaultNotificationObjectTypePtrOutput {
	return o.ApplyT(func(v *LogicallyAirGappedBackupVault) LogicallyAirGappedBackupVaultNotificationObjectTypePtrOutput {
		return v.Notifications
	}).(LogicallyAirGappedBackupVaultNotificationObjectTypePtrOutput)
}

func (o LogicallyAirGappedBackupVaultOutput) VaultState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogicallyAirGappedBackupVault) pulumi.StringPtrOutput { return v.VaultState }).(pulumi.StringPtrOutput)
}

func (o LogicallyAirGappedBackupVaultOutput) VaultType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogicallyAirGappedBackupVault) pulumi.StringPtrOutput { return v.VaultType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogicallyAirGappedBackupVaultInput)(nil)).Elem(), &LogicallyAirGappedBackupVault{})
	pulumi.RegisterOutputType(LogicallyAirGappedBackupVaultOutput{})
}
