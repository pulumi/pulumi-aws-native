// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::DevEndpoint
func LookupDevEndpoint(ctx *pulumi.Context, args *LookupDevEndpointArgs, opts ...pulumi.InvokeOption) (*LookupDevEndpointResult, error) {
	var rv LookupDevEndpointResult
	err := ctx.Invoke("aws-native:glue:getDevEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDevEndpointArgs struct {
	Id string `pulumi:"id"`
}

type LookupDevEndpointResult struct {
	Arguments             interface{} `pulumi:"arguments"`
	ExtraJarsS3Path       *string     `pulumi:"extraJarsS3Path"`
	ExtraPythonLibsS3Path *string     `pulumi:"extraPythonLibsS3Path"`
	GlueVersion           *string     `pulumi:"glueVersion"`
	Id                    *string     `pulumi:"id"`
	NumberOfNodes         *int        `pulumi:"numberOfNodes"`
	NumberOfWorkers       *int        `pulumi:"numberOfWorkers"`
	PublicKey             *string     `pulumi:"publicKey"`
	PublicKeys            []string    `pulumi:"publicKeys"`
	RoleArn               *string     `pulumi:"roleArn"`
	SecurityConfiguration *string     `pulumi:"securityConfiguration"`
	SecurityGroupIds      []string    `pulumi:"securityGroupIds"`
	SubnetId              *string     `pulumi:"subnetId"`
	Tags                  interface{} `pulumi:"tags"`
	WorkerType            *string     `pulumi:"workerType"`
}

func LookupDevEndpointOutput(ctx *pulumi.Context, args LookupDevEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupDevEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDevEndpointResult, error) {
			args := v.(LookupDevEndpointArgs)
			r, err := LookupDevEndpoint(ctx, &args, opts...)
			var s LookupDevEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDevEndpointResultOutput)
}

type LookupDevEndpointOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDevEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevEndpointArgs)(nil)).Elem()
}

type LookupDevEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupDevEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevEndpointResult)(nil)).Elem()
}

func (o LookupDevEndpointResultOutput) ToLookupDevEndpointResultOutput() LookupDevEndpointResultOutput {
	return o
}

func (o LookupDevEndpointResultOutput) ToLookupDevEndpointResultOutputWithContext(ctx context.Context) LookupDevEndpointResultOutput {
	return o
}

func (o LookupDevEndpointResultOutput) Arguments() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) interface{} { return v.Arguments }).(pulumi.AnyOutput)
}

func (o LookupDevEndpointResultOutput) ExtraJarsS3Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) *string { return v.ExtraJarsS3Path }).(pulumi.StringPtrOutput)
}

func (o LookupDevEndpointResultOutput) ExtraPythonLibsS3Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) *string { return v.ExtraPythonLibsS3Path }).(pulumi.StringPtrOutput)
}

func (o LookupDevEndpointResultOutput) GlueVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) *string { return v.GlueVersion }).(pulumi.StringPtrOutput)
}

func (o LookupDevEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupDevEndpointResultOutput) NumberOfNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) *int { return v.NumberOfNodes }).(pulumi.IntPtrOutput)
}

func (o LookupDevEndpointResultOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) *int { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o LookupDevEndpointResultOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) *string { return v.PublicKey }).(pulumi.StringPtrOutput)
}

func (o LookupDevEndpointResultOutput) PublicKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) []string { return v.PublicKeys }).(pulumi.StringArrayOutput)
}

func (o LookupDevEndpointResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupDevEndpointResultOutput) SecurityConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) *string { return v.SecurityConfiguration }).(pulumi.StringPtrOutput)
}

func (o LookupDevEndpointResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupDevEndpointResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o LookupDevEndpointResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupDevEndpointResultOutput) WorkerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevEndpointResult) *string { return v.WorkerType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDevEndpointResultOutput{})
}