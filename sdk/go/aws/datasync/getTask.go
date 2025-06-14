// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataSync::Task.
func LookupTask(ctx *pulumi.Context, args *LookupTaskArgs, opts ...pulumi.InvokeOption) (*LookupTaskResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTaskResult
	err := ctx.Invoke("aws-native:datasync:getTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTaskArgs struct {
	// The ARN of the task.
	TaskArn string `pulumi:"taskArn"`
}

type LookupTaskResult struct {
	// The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.
	CloudWatchLogGroupArn *string `pulumi:"cloudWatchLogGroupArn"`
	// The ARNs of the destination elastic network interfaces (ENIs) that were created for your subnet.
	DestinationNetworkInterfaceArns []string `pulumi:"destinationNetworkInterfaceArns"`
	// Specifies exclude filters that define the files, objects, and folders in your source location that you don't want DataSync to transfer. For more information and examples, see [Specifying what DataSync transfers by using filters](https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
	Excludes []TaskFilterRule `pulumi:"excludes"`
	// Specifies include filters that define the files, objects, and folders in your source location that you want DataSync to transfer. For more information and examples, see [Specifying what DataSync transfers by using filters](https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
	Includes []TaskFilterRule `pulumi:"includes"`
	// The configuration of the manifest that lists the files or objects that you want DataSync to transfer. For more information, see [Specifying what DataSync transfers by using a manifest](https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html) .
	ManifestConfig *TaskManifestConfig `pulumi:"manifestConfig"`
	// The name of a task. This value is a text reference that is used to identify the task in the console.
	Name *string `pulumi:"name"`
	// Specifies your task's settings, such as preserving file metadata, verifying data integrity, among other options.
	Options *TaskOptions `pulumi:"options"`
	// Specifies a schedule for when you want your task to run. For more information, see [Scheduling your task](https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html) .
	Schedule *TaskSchedule `pulumi:"schedule"`
	// The ARNs of the source ENIs that were created for your subnet.
	SourceNetworkInterfaceArns []string `pulumi:"sourceNetworkInterfaceArns"`
	// The status of the task that was described.
	Status *TaskStatus `pulumi:"status"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The ARN of the task.
	TaskArn *string `pulumi:"taskArn"`
	// The configuration of your task report, which provides detailed information about your DataSync transfer. For more information, see [Monitoring your DataSync transfers with task reports](https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html) .
	TaskReportConfig *TaskReportConfig `pulumi:"taskReportConfig"`
}

func LookupTaskOutput(ctx *pulumi.Context, args LookupTaskOutputArgs, opts ...pulumi.InvokeOption) LookupTaskResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTaskResultOutput, error) {
			args := v.(LookupTaskArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:datasync:getTask", args, LookupTaskResultOutput{}, options).(LookupTaskResultOutput), nil
		}).(LookupTaskResultOutput)
}

type LookupTaskOutputArgs struct {
	// The ARN of the task.
	TaskArn pulumi.StringInput `pulumi:"taskArn"`
}

func (LookupTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskArgs)(nil)).Elem()
}

type LookupTaskResultOutput struct{ *pulumi.OutputState }

func (LookupTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskResult)(nil)).Elem()
}

func (o LookupTaskResultOutput) ToLookupTaskResultOutput() LookupTaskResultOutput {
	return o
}

func (o LookupTaskResultOutput) ToLookupTaskResultOutputWithContext(ctx context.Context) LookupTaskResultOutput {
	return o
}

// The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.
func (o LookupTaskResultOutput) CloudWatchLogGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *string { return v.CloudWatchLogGroupArn }).(pulumi.StringPtrOutput)
}

// The ARNs of the destination elastic network interfaces (ENIs) that were created for your subnet.
func (o LookupTaskResultOutput) DestinationNetworkInterfaceArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTaskResult) []string { return v.DestinationNetworkInterfaceArns }).(pulumi.StringArrayOutput)
}

// Specifies exclude filters that define the files, objects, and folders in your source location that you don't want DataSync to transfer. For more information and examples, see [Specifying what DataSync transfers by using filters](https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
func (o LookupTaskResultOutput) Excludes() TaskFilterRuleArrayOutput {
	return o.ApplyT(func(v LookupTaskResult) []TaskFilterRule { return v.Excludes }).(TaskFilterRuleArrayOutput)
}

// Specifies include filters that define the files, objects, and folders in your source location that you want DataSync to transfer. For more information and examples, see [Specifying what DataSync transfers by using filters](https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
func (o LookupTaskResultOutput) Includes() TaskFilterRuleArrayOutput {
	return o.ApplyT(func(v LookupTaskResult) []TaskFilterRule { return v.Includes }).(TaskFilterRuleArrayOutput)
}

// The configuration of the manifest that lists the files or objects that you want DataSync to transfer. For more information, see [Specifying what DataSync transfers by using a manifest](https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html) .
func (o LookupTaskResultOutput) ManifestConfig() TaskManifestConfigPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *TaskManifestConfig { return v.ManifestConfig }).(TaskManifestConfigPtrOutput)
}

// The name of a task. This value is a text reference that is used to identify the task in the console.
func (o LookupTaskResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Specifies your task's settings, such as preserving file metadata, verifying data integrity, among other options.
func (o LookupTaskResultOutput) Options() TaskOptionsPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *TaskOptions { return v.Options }).(TaskOptionsPtrOutput)
}

// Specifies a schedule for when you want your task to run. For more information, see [Scheduling your task](https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html) .
func (o LookupTaskResultOutput) Schedule() TaskSchedulePtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *TaskSchedule { return v.Schedule }).(TaskSchedulePtrOutput)
}

// The ARNs of the source ENIs that were created for your subnet.
func (o LookupTaskResultOutput) SourceNetworkInterfaceArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTaskResult) []string { return v.SourceNetworkInterfaceArns }).(pulumi.StringArrayOutput)
}

// The status of the task that was described.
func (o LookupTaskResultOutput) Status() TaskStatusPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *TaskStatus { return v.Status }).(TaskStatusPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupTaskResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupTaskResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The ARN of the task.
func (o LookupTaskResultOutput) TaskArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *string { return v.TaskArn }).(pulumi.StringPtrOutput)
}

// The configuration of your task report, which provides detailed information about your DataSync transfer. For more information, see [Monitoring your DataSync transfers with task reports](https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html) .
func (o LookupTaskResultOutput) TaskReportConfig() TaskReportConfigPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *TaskReportConfig { return v.TaskReportConfig }).(TaskReportConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTaskResultOutput{})
}
