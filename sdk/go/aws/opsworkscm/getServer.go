// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opsworkscm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::OpsWorksCM::Server
func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServerResult
	err := ctx.Invoke("aws-native:opsworkscm:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	// The name of the server. The server name must be unique within your AWS account, within each region. Server names must start with a letter; then letters, numbers, or hyphens (-) are allowed, up to a maximum of 40 characters.
	ServerName string `pulumi:"serverName"`
}

type LookupServerResult struct {
	// The Amazon Resource Name (ARN) of the server, such as `arn:aws:OpsWorksCM:us-east-1:123456789012:server/server-a1bzhi` .
	Arn *string `pulumi:"arn"`
	// The number of automated backups that you want to keep. Whenever a new backup is created, AWS OpsWorks CM deletes the oldest backups if this number is exceeded. The default value is `1` .
	BackupRetentionCount *int `pulumi:"backupRetentionCount"`
	// Enable or disable scheduled backups. Valid values are `true` or `false` . The default value is `true` .
	DisableAutomatedBackup *bool `pulumi:"disableAutomatedBackup"`
	// A DNS name that can be used to access the engine. Example: `myserver-asdfghjkl.us-east-1.opsworks.io` .
	Endpoint *string `pulumi:"endpoint"`
	// The start time for a one-hour period during which AWS OpsWorks CM backs up application-level data on your server if automated backups are enabled. Valid values must be specified in one of the following formats:
	//
	// - `HH:MM` for daily backups
	// - `DDD:HH:MM` for weekly backups
	//
	// `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random, daily start time.
	//
	// *Example:* `08:00` , which represents a daily start time of 08:00 UTC.
	//
	// *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
	PreferredBackupWindow *string `pulumi:"preferredBackupWindow"`
	// The start time for a one-hour period each week during which AWS OpsWorks CM performs maintenance on the instance. Valid values must be specified in the following format: `DDD:HH:MM` . `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random one-hour period on Tuesday, Wednesday, or Friday. See `TimeWindowDefinition` for more information.
	//
	// *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
}

func LookupServerOutput(ctx *pulumi.Context, args LookupServerOutputArgs, opts ...pulumi.InvokeOption) LookupServerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupServerResultOutput, error) {
			args := v.(LookupServerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:opsworkscm:getServer", args, LookupServerResultOutput{}, options).(LookupServerResultOutput), nil
		}).(LookupServerResultOutput)
}

type LookupServerOutputArgs struct {
	// The name of the server. The server name must be unique within your AWS account, within each region. Server names must start with a letter; then letters, numbers, or hyphens (-) are allowed, up to a maximum of 40 characters.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerArgs)(nil)).Elem()
}

type LookupServerResultOutput struct{ *pulumi.OutputState }

func (LookupServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerResult)(nil)).Elem()
}

func (o LookupServerResultOutput) ToLookupServerResultOutput() LookupServerResultOutput {
	return o
}

func (o LookupServerResultOutput) ToLookupServerResultOutputWithContext(ctx context.Context) LookupServerResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the server, such as `arn:aws:OpsWorksCM:us-east-1:123456789012:server/server-a1bzhi` .
func (o LookupServerResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The number of automated backups that you want to keep. Whenever a new backup is created, AWS OpsWorks CM deletes the oldest backups if this number is exceeded. The default value is `1` .
func (o LookupServerResultOutput) BackupRetentionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *int { return v.BackupRetentionCount }).(pulumi.IntPtrOutput)
}

// Enable or disable scheduled backups. Valid values are `true` or `false` . The default value is `true` .
func (o LookupServerResultOutput) DisableAutomatedBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *bool { return v.DisableAutomatedBackup }).(pulumi.BoolPtrOutput)
}

// A DNS name that can be used to access the engine. Example: `myserver-asdfghjkl.us-east-1.opsworks.io` .
func (o LookupServerResultOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// The start time for a one-hour period during which AWS OpsWorks CM backs up application-level data on your server if automated backups are enabled. Valid values must be specified in one of the following formats:
//
// - `HH:MM` for daily backups
// - `DDD:HH:MM` for weekly backups
//
// `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random, daily start time.
//
// *Example:* `08:00` , which represents a daily start time of 08:00 UTC.
//
// *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
func (o LookupServerResultOutput) PreferredBackupWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.PreferredBackupWindow }).(pulumi.StringPtrOutput)
}

// The start time for a one-hour period each week during which AWS OpsWorks CM performs maintenance on the instance. Valid values must be specified in the following format: `DDD:HH:MM` . `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random one-hour period on Tuesday, Wednesday, or Friday. See `TimeWindowDefinition` for more information.
//
// *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
func (o LookupServerResultOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerResult) *string { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerResultOutput{})
}
