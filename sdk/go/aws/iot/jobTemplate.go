// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoT::JobTemplate. Job templates enable you to preconfigure jobs so that you can deploy them to multiple sets of target devices.
type JobTemplate struct {
	pulumi.CustomResourceState

	// The criteria that determine when and how a job abort takes place.
	AbortConfig AbortConfigPropertiesPtrOutput `pulumi:"abortConfig"`
	// The ARN of the job to use as the basis for the job template.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description of the Job Template.
	Description pulumi.StringOutput `pulumi:"description"`
	// The package version Amazon Resource Names (ARNs) that are installed on the device’s reserved named shadow ( `$package` ) when the job successfully completes.
	//
	// *Note:* Up to 25 package version ARNS are allowed.
	DestinationPackageVersions pulumi.StringArrayOutput `pulumi:"destinationPackageVersions"`
	// The job document. Required if you don't specify a value for documentSource.
	Document pulumi.StringPtrOutput `pulumi:"document"`
	// An S3 link to the job document to use in the template. Required if you don't specify a value for document.
	DocumentSource pulumi.StringPtrOutput `pulumi:"documentSource"`
	// Optional for copying a JobTemplate from a pre-existing Job configuration.
	JobArn pulumi.StringPtrOutput `pulumi:"jobArn"`
	// Allows you to create the criteria to retry a job.
	JobExecutionsRetryConfig JobExecutionsRetryConfigPropertiesPtrOutput `pulumi:"jobExecutionsRetryConfig"`
	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig JobExecutionsRolloutConfigPropertiesPtrOutput `pulumi:"jobExecutionsRolloutConfig"`
	// A unique identifier for the job template. We recommend using a UUID. Alpha-numeric characters, "-", and "_" are valid for use here.
	JobTemplateId pulumi.StringOutput `pulumi:"jobTemplateId"`
	// An optional configuration within the SchedulingConfig to setup a recurring maintenance window with a predetermined start time and duration for the rollout of a job document to all devices in a target group for a job.
	MaintenanceWindows JobTemplateMaintenanceWindowArrayOutput `pulumi:"maintenanceWindows"`
	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig PresignedUrlConfigPropertiesPtrOutput `pulumi:"presignedUrlConfig"`
	// Metadata that can be used to manage the JobTemplate.
	Tags aws.CreateOnlyTagArrayOutput `pulumi:"tags"`
	// Specifies the amount of time each device has to finish its execution of the job.
	TimeoutConfig TimeoutConfigPropertiesPtrOutput `pulumi:"timeoutConfig"`
}

// NewJobTemplate registers a new resource with the given unique name, arguments, and options.
func NewJobTemplate(ctx *pulumi.Context,
	name string, args *JobTemplateArgs, opts ...pulumi.ResourceOption) (*JobTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.JobTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'JobTemplateId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"abortConfig",
		"description",
		"destinationPackageVersions[*]",
		"document",
		"documentSource",
		"jobArn",
		"jobExecutionsRetryConfig",
		"jobExecutionsRolloutConfig",
		"jobTemplateId",
		"maintenanceWindows[*]",
		"presignedUrlConfig",
		"tags[*]",
		"timeoutConfig",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource JobTemplate
	err := ctx.RegisterResource("aws-native:iot:JobTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobTemplate gets an existing JobTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobTemplateState, opts ...pulumi.ResourceOption) (*JobTemplate, error) {
	var resource JobTemplate
	err := ctx.ReadResource("aws-native:iot:JobTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobTemplate resources.
type jobTemplateState struct {
}

type JobTemplateState struct {
}

func (JobTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTemplateState)(nil)).Elem()
}

type jobTemplateArgs struct {
	// The criteria that determine when and how a job abort takes place.
	AbortConfig *AbortConfigProperties `pulumi:"abortConfig"`
	// A description of the Job Template.
	Description string `pulumi:"description"`
	// The package version Amazon Resource Names (ARNs) that are installed on the device’s reserved named shadow ( `$package` ) when the job successfully completes.
	//
	// *Note:* Up to 25 package version ARNS are allowed.
	DestinationPackageVersions []string `pulumi:"destinationPackageVersions"`
	// The job document. Required if you don't specify a value for documentSource.
	Document *string `pulumi:"document"`
	// An S3 link to the job document to use in the template. Required if you don't specify a value for document.
	DocumentSource *string `pulumi:"documentSource"`
	// Optional for copying a JobTemplate from a pre-existing Job configuration.
	JobArn *string `pulumi:"jobArn"`
	// Allows you to create the criteria to retry a job.
	JobExecutionsRetryConfig *JobExecutionsRetryConfigProperties `pulumi:"jobExecutionsRetryConfig"`
	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig *JobExecutionsRolloutConfigProperties `pulumi:"jobExecutionsRolloutConfig"`
	// A unique identifier for the job template. We recommend using a UUID. Alpha-numeric characters, "-", and "_" are valid for use here.
	JobTemplateId string `pulumi:"jobTemplateId"`
	// An optional configuration within the SchedulingConfig to setup a recurring maintenance window with a predetermined start time and duration for the rollout of a job document to all devices in a target group for a job.
	MaintenanceWindows []JobTemplateMaintenanceWindow `pulumi:"maintenanceWindows"`
	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig *PresignedUrlConfigProperties `pulumi:"presignedUrlConfig"`
	// Metadata that can be used to manage the JobTemplate.
	Tags []aws.CreateOnlyTag `pulumi:"tags"`
	// Specifies the amount of time each device has to finish its execution of the job.
	TimeoutConfig *TimeoutConfigProperties `pulumi:"timeoutConfig"`
}

// The set of arguments for constructing a JobTemplate resource.
type JobTemplateArgs struct {
	// The criteria that determine when and how a job abort takes place.
	AbortConfig AbortConfigPropertiesPtrInput
	// A description of the Job Template.
	Description pulumi.StringInput
	// The package version Amazon Resource Names (ARNs) that are installed on the device’s reserved named shadow ( `$package` ) when the job successfully completes.
	//
	// *Note:* Up to 25 package version ARNS are allowed.
	DestinationPackageVersions pulumi.StringArrayInput
	// The job document. Required if you don't specify a value for documentSource.
	Document pulumi.StringPtrInput
	// An S3 link to the job document to use in the template. Required if you don't specify a value for document.
	DocumentSource pulumi.StringPtrInput
	// Optional for copying a JobTemplate from a pre-existing Job configuration.
	JobArn pulumi.StringPtrInput
	// Allows you to create the criteria to retry a job.
	JobExecutionsRetryConfig JobExecutionsRetryConfigPropertiesPtrInput
	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig JobExecutionsRolloutConfigPropertiesPtrInput
	// A unique identifier for the job template. We recommend using a UUID. Alpha-numeric characters, "-", and "_" are valid for use here.
	JobTemplateId pulumi.StringInput
	// An optional configuration within the SchedulingConfig to setup a recurring maintenance window with a predetermined start time and duration for the rollout of a job document to all devices in a target group for a job.
	MaintenanceWindows JobTemplateMaintenanceWindowArrayInput
	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig PresignedUrlConfigPropertiesPtrInput
	// Metadata that can be used to manage the JobTemplate.
	Tags aws.CreateOnlyTagArrayInput
	// Specifies the amount of time each device has to finish its execution of the job.
	TimeoutConfig TimeoutConfigPropertiesPtrInput
}

func (JobTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTemplateArgs)(nil)).Elem()
}

type JobTemplateInput interface {
	pulumi.Input

	ToJobTemplateOutput() JobTemplateOutput
	ToJobTemplateOutputWithContext(ctx context.Context) JobTemplateOutput
}

func (*JobTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTemplate)(nil)).Elem()
}

func (i *JobTemplate) ToJobTemplateOutput() JobTemplateOutput {
	return i.ToJobTemplateOutputWithContext(context.Background())
}

func (i *JobTemplate) ToJobTemplateOutputWithContext(ctx context.Context) JobTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTemplateOutput)
}

type JobTemplateOutput struct{ *pulumi.OutputState }

func (JobTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTemplate)(nil)).Elem()
}

func (o JobTemplateOutput) ToJobTemplateOutput() JobTemplateOutput {
	return o
}

func (o JobTemplateOutput) ToJobTemplateOutputWithContext(ctx context.Context) JobTemplateOutput {
	return o
}

// The criteria that determine when and how a job abort takes place.
func (o JobTemplateOutput) AbortConfig() AbortConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *JobTemplate) AbortConfigPropertiesPtrOutput { return v.AbortConfig }).(AbortConfigPropertiesPtrOutput)
}

// The ARN of the job to use as the basis for the job template.
func (o JobTemplateOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A description of the Job Template.
func (o JobTemplateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The package version Amazon Resource Names (ARNs) that are installed on the device’s reserved named shadow ( `$package` ) when the job successfully completes.
//
// *Note:* Up to 25 package version ARNS are allowed.
func (o JobTemplateOutput) DestinationPackageVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringArrayOutput { return v.DestinationPackageVersions }).(pulumi.StringArrayOutput)
}

// The job document. Required if you don't specify a value for documentSource.
func (o JobTemplateOutput) Document() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringPtrOutput { return v.Document }).(pulumi.StringPtrOutput)
}

// An S3 link to the job document to use in the template. Required if you don't specify a value for document.
func (o JobTemplateOutput) DocumentSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringPtrOutput { return v.DocumentSource }).(pulumi.StringPtrOutput)
}

// Optional for copying a JobTemplate from a pre-existing Job configuration.
func (o JobTemplateOutput) JobArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringPtrOutput { return v.JobArn }).(pulumi.StringPtrOutput)
}

// Allows you to create the criteria to retry a job.
func (o JobTemplateOutput) JobExecutionsRetryConfig() JobExecutionsRetryConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *JobTemplate) JobExecutionsRetryConfigPropertiesPtrOutput { return v.JobExecutionsRetryConfig }).(JobExecutionsRetryConfigPropertiesPtrOutput)
}

// Allows you to create a staged rollout of a job.
func (o JobTemplateOutput) JobExecutionsRolloutConfig() JobExecutionsRolloutConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *JobTemplate) JobExecutionsRolloutConfigPropertiesPtrOutput {
		return v.JobExecutionsRolloutConfig
	}).(JobExecutionsRolloutConfigPropertiesPtrOutput)
}

// A unique identifier for the job template. We recommend using a UUID. Alpha-numeric characters, "-", and "_" are valid for use here.
func (o JobTemplateOutput) JobTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringOutput { return v.JobTemplateId }).(pulumi.StringOutput)
}

// An optional configuration within the SchedulingConfig to setup a recurring maintenance window with a predetermined start time and duration for the rollout of a job document to all devices in a target group for a job.
func (o JobTemplateOutput) MaintenanceWindows() JobTemplateMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v *JobTemplate) JobTemplateMaintenanceWindowArrayOutput { return v.MaintenanceWindows }).(JobTemplateMaintenanceWindowArrayOutput)
}

// Configuration for pre-signed S3 URLs.
func (o JobTemplateOutput) PresignedUrlConfig() PresignedUrlConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *JobTemplate) PresignedUrlConfigPropertiesPtrOutput { return v.PresignedUrlConfig }).(PresignedUrlConfigPropertiesPtrOutput)
}

// Metadata that can be used to manage the JobTemplate.
func (o JobTemplateOutput) Tags() aws.CreateOnlyTagArrayOutput {
	return o.ApplyT(func(v *JobTemplate) aws.CreateOnlyTagArrayOutput { return v.Tags }).(aws.CreateOnlyTagArrayOutput)
}

// Specifies the amount of time each device has to finish its execution of the job.
func (o JobTemplateOutput) TimeoutConfig() TimeoutConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *JobTemplate) TimeoutConfigPropertiesPtrOutput { return v.TimeoutConfig }).(TimeoutConfigPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobTemplateInput)(nil)).Elem(), &JobTemplate{})
	pulumi.RegisterOutputType(JobTemplateOutput{})
}
