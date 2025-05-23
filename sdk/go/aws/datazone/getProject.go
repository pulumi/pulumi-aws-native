// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datazone

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Amazon DataZone projects are business use case–based groupings of people, assets (data), and tools used to simplify access to the AWS analytics.
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("aws-native:datazone:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	// The identifier of the Amazon DataZone domain in which the project was created.
	DomainId string `pulumi:"domainId"`
	// The ID of the Amazon DataZone project.
	Id string `pulumi:"id"`
}

type LookupProjectResult struct {
	// The timestamp of when the project was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The Amazon DataZone user who created the project.
	CreatedBy *string `pulumi:"createdBy"`
	// The description of the Amazon DataZone project.
	Description *string `pulumi:"description"`
	// The identifier of the Amazon DataZone domain in which the project was created.
	DomainId *string `pulumi:"domainId"`
	// The glossary terms that can be used in this Amazon DataZone project.
	GlossaryTerms []string `pulumi:"glossaryTerms"`
	// The ID of the Amazon DataZone project.
	Id *string `pulumi:"id"`
	// The timestamp of when the project was last updated.
	LastUpdatedAt *string `pulumi:"lastUpdatedAt"`
	// The name of the Amazon DataZone project.
	Name *string `pulumi:"name"`
	// The status of the project.
	ProjectStatus *ProjectStatus `pulumi:"projectStatus"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectResultOutput, error) {
			args := v.(LookupProjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:datazone:getProject", args, LookupProjectResultOutput{}, options).(LookupProjectResultOutput), nil
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	// The identifier of the Amazon DataZone domain in which the project was created.
	DomainId pulumi.StringInput `pulumi:"domainId"`
	// The ID of the Amazon DataZone project.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// The timestamp of when the project was created.
func (o LookupProjectResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The Amazon DataZone user who created the project.
func (o LookupProjectResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The description of the Amazon DataZone project.
func (o LookupProjectResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier of the Amazon DataZone domain in which the project was created.
func (o LookupProjectResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

// The glossary terms that can be used in this Amazon DataZone project.
func (o LookupProjectResultOutput) GlossaryTerms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []string { return v.GlossaryTerms }).(pulumi.StringArrayOutput)
}

// The ID of the Amazon DataZone project.
func (o LookupProjectResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The timestamp of when the project was last updated.
func (o LookupProjectResultOutput) LastUpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.LastUpdatedAt }).(pulumi.StringPtrOutput)
}

// The name of the Amazon DataZone project.
func (o LookupProjectResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The status of the project.
func (o LookupProjectResultOutput) ProjectStatus() ProjectStatusPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectStatus { return v.ProjectStatus }).(ProjectStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
