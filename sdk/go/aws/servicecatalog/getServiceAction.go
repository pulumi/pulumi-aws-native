// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema for AWS::ServiceCatalog::ServiceAction
func LookupServiceAction(ctx *pulumi.Context, args *LookupServiceActionArgs, opts ...pulumi.InvokeOption) (*LookupServiceActionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceActionResult
	err := ctx.Invoke("aws-native:servicecatalog:getServiceAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceActionArgs struct {
	// The self-service action identifier. For example, `act-fs7abcd89wxyz` .
	Id string `pulumi:"id"`
}

type LookupServiceActionResult struct {
	// A map that defines the self-service action.
	Definition []ServiceActionDefinitionParameter `pulumi:"definition"`
	// The self-service action definition type. For example, `SSM_AUTOMATION` .
	DefinitionType *ServiceActionDefinitionType `pulumi:"definitionType"`
	// The self-service action description.
	Description *string `pulumi:"description"`
	// The self-service action identifier. For example, `act-fs7abcd89wxyz` .
	Id *string `pulumi:"id"`
	// The self-service action name.
	Name *string `pulumi:"name"`
}

func LookupServiceActionOutput(ctx *pulumi.Context, args LookupServiceActionOutputArgs, opts ...pulumi.InvokeOption) LookupServiceActionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupServiceActionResultOutput, error) {
			args := v.(LookupServiceActionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:servicecatalog:getServiceAction", args, LookupServiceActionResultOutput{}, options).(LookupServiceActionResultOutput), nil
		}).(LookupServiceActionResultOutput)
}

type LookupServiceActionOutputArgs struct {
	// The self-service action identifier. For example, `act-fs7abcd89wxyz` .
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupServiceActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceActionArgs)(nil)).Elem()
}

type LookupServiceActionResultOutput struct{ *pulumi.OutputState }

func (LookupServiceActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceActionResult)(nil)).Elem()
}

func (o LookupServiceActionResultOutput) ToLookupServiceActionResultOutput() LookupServiceActionResultOutput {
	return o
}

func (o LookupServiceActionResultOutput) ToLookupServiceActionResultOutputWithContext(ctx context.Context) LookupServiceActionResultOutput {
	return o
}

// A map that defines the self-service action.
func (o LookupServiceActionResultOutput) Definition() ServiceActionDefinitionParameterArrayOutput {
	return o.ApplyT(func(v LookupServiceActionResult) []ServiceActionDefinitionParameter { return v.Definition }).(ServiceActionDefinitionParameterArrayOutput)
}

// The self-service action definition type. For example, `SSM_AUTOMATION` .
func (o LookupServiceActionResultOutput) DefinitionType() ServiceActionDefinitionTypePtrOutput {
	return o.ApplyT(func(v LookupServiceActionResult) *ServiceActionDefinitionType { return v.DefinitionType }).(ServiceActionDefinitionTypePtrOutput)
}

// The self-service action description.
func (o LookupServiceActionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceActionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The self-service action identifier. For example, `act-fs7abcd89wxyz` .
func (o LookupServiceActionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceActionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The self-service action name.
func (o LookupServiceActionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceActionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceActionResultOutput{})
}
