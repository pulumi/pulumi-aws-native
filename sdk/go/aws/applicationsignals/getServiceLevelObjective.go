// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package applicationsignals

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApplicationSignals::ServiceLevelObjective
func LookupServiceLevelObjective(ctx *pulumi.Context, args *LookupServiceLevelObjectiveArgs, opts ...pulumi.InvokeOption) (*LookupServiceLevelObjectiveResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceLevelObjectiveResult
	err := ctx.Invoke("aws-native:applicationsignals:getServiceLevelObjective", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceLevelObjectiveArgs struct {
	// The ARN of this SLO.
	Arn string `pulumi:"arn"`
}

type LookupServiceLevelObjectiveResult struct {
	// The ARN of this SLO.
	Arn *string `pulumi:"arn"`
	// Epoch time in seconds of the time that this SLO was created
	CreatedTime *int `pulumi:"createdTime"`
	// An optional description for this SLO. Default is 'No description'
	Description *string `pulumi:"description"`
	// This structure contains the attributes that determine the goal of an SLO. This includes the time period for evaluation and the attainment threshold.
	Goal *ServiceLevelObjectiveGoal `pulumi:"goal"`
	// Epoch time in seconds of the time that this SLO was most recently updated
	LastUpdatedTime *int `pulumi:"lastUpdatedTime"`
	// A structure containing information about the performance metric that this SLO monitors.
	Sli *ServiceLevelObjectiveSli `pulumi:"sli"`
	// A list of key-value pairs to associate with the SLO. You can associate as many as 50 tags with an SLO. To be able to associate tags with the SLO when you create the SLO, you must have the cloudwatch:TagResource permission.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupServiceLevelObjectiveOutput(ctx *pulumi.Context, args LookupServiceLevelObjectiveOutputArgs, opts ...pulumi.InvokeOption) LookupServiceLevelObjectiveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceLevelObjectiveResult, error) {
			args := v.(LookupServiceLevelObjectiveArgs)
			r, err := LookupServiceLevelObjective(ctx, &args, opts...)
			var s LookupServiceLevelObjectiveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceLevelObjectiveResultOutput)
}

type LookupServiceLevelObjectiveOutputArgs struct {
	// The ARN of this SLO.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupServiceLevelObjectiveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceLevelObjectiveArgs)(nil)).Elem()
}

type LookupServiceLevelObjectiveResultOutput struct{ *pulumi.OutputState }

func (LookupServiceLevelObjectiveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceLevelObjectiveResult)(nil)).Elem()
}

func (o LookupServiceLevelObjectiveResultOutput) ToLookupServiceLevelObjectiveResultOutput() LookupServiceLevelObjectiveResultOutput {
	return o
}

func (o LookupServiceLevelObjectiveResultOutput) ToLookupServiceLevelObjectiveResultOutputWithContext(ctx context.Context) LookupServiceLevelObjectiveResultOutput {
	return o
}

// The ARN of this SLO.
func (o LookupServiceLevelObjectiveResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Epoch time in seconds of the time that this SLO was created
func (o LookupServiceLevelObjectiveResultOutput) CreatedTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) *int { return v.CreatedTime }).(pulumi.IntPtrOutput)
}

// An optional description for this SLO. Default is 'No description'
func (o LookupServiceLevelObjectiveResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// This structure contains the attributes that determine the goal of an SLO. This includes the time period for evaluation and the attainment threshold.
func (o LookupServiceLevelObjectiveResultOutput) Goal() ServiceLevelObjectiveGoalPtrOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) *ServiceLevelObjectiveGoal { return v.Goal }).(ServiceLevelObjectiveGoalPtrOutput)
}

// Epoch time in seconds of the time that this SLO was most recently updated
func (o LookupServiceLevelObjectiveResultOutput) LastUpdatedTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) *int { return v.LastUpdatedTime }).(pulumi.IntPtrOutput)
}

// A structure containing information about the performance metric that this SLO monitors.
func (o LookupServiceLevelObjectiveResultOutput) Sli() ServiceLevelObjectiveSliPtrOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) *ServiceLevelObjectiveSli { return v.Sli }).(ServiceLevelObjectiveSliPtrOutput)
}

// A list of key-value pairs to associate with the SLO. You can associate as many as 50 tags with an SLO. To be able to associate tags with the SLO when you create the SLO, you must have the cloudwatch:TagResource permission.
//
// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
func (o LookupServiceLevelObjectiveResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupServiceLevelObjectiveResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceLevelObjectiveResultOutput{})
}