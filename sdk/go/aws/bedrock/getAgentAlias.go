// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bedrock

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Bedrock::AgentAlias Resource Type
func LookupAgentAlias(ctx *pulumi.Context, args *LookupAgentAliasArgs, opts ...pulumi.InvokeOption) (*LookupAgentAliasResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAgentAliasResult
	err := ctx.Invoke("aws-native:bedrock:getAgentAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAgentAliasArgs struct {
	// Id for an Agent Alias generated at the server side.
	AgentAliasId string `pulumi:"agentAliasId"`
	// Identifier for a resource.
	AgentId string `pulumi:"agentId"`
}

type LookupAgentAliasResult struct {
	// Arn representation of the Agent Alias.
	AgentAliasArn *string `pulumi:"agentAliasArn"`
	// The list of history events for an alias for an Agent.
	AgentAliasHistoryEvents []AgentAliasHistoryEvent `pulumi:"agentAliasHistoryEvents"`
	// Id for an Agent Alias generated at the server side.
	AgentAliasId *string `pulumi:"agentAliasId"`
	// Name for a resource.
	AgentAliasName *string `pulumi:"agentAliasName"`
	// The status of the alias of the agent and whether it is ready for use. The following statuses are possible:
	//
	// - CREATING – The agent alias is being created.
	// - PREPARED – The agent alias is finished being created or updated and is ready to be invoked.
	// - FAILED – The agent alias API operation failed.
	// - UPDATING – The agent alias is being updated.
	// - DELETING – The agent alias is being deleted.
	AgentAliasStatus *AgentAliasStatus `pulumi:"agentAliasStatus"`
	// Time Stamp.
	CreatedAt *string `pulumi:"createdAt"`
	// Description of the Resource.
	Description *string `pulumi:"description"`
	// Routing configuration for an Agent alias.
	RoutingConfiguration []AgentAliasRoutingConfigurationListItem `pulumi:"routingConfiguration"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	Tags map[string]string `pulumi:"tags"`
	// Time Stamp.
	UpdatedAt *string `pulumi:"updatedAt"`
}

func LookupAgentAliasOutput(ctx *pulumi.Context, args LookupAgentAliasOutputArgs, opts ...pulumi.InvokeOption) LookupAgentAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAgentAliasResult, error) {
			args := v.(LookupAgentAliasArgs)
			r, err := LookupAgentAlias(ctx, &args, opts...)
			var s LookupAgentAliasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAgentAliasResultOutput)
}

type LookupAgentAliasOutputArgs struct {
	// Id for an Agent Alias generated at the server side.
	AgentAliasId pulumi.StringInput `pulumi:"agentAliasId"`
	// Identifier for a resource.
	AgentId pulumi.StringInput `pulumi:"agentId"`
}

func (LookupAgentAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentAliasArgs)(nil)).Elem()
}

type LookupAgentAliasResultOutput struct{ *pulumi.OutputState }

func (LookupAgentAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentAliasResult)(nil)).Elem()
}

func (o LookupAgentAliasResultOutput) ToLookupAgentAliasResultOutput() LookupAgentAliasResultOutput {
	return o
}

func (o LookupAgentAliasResultOutput) ToLookupAgentAliasResultOutputWithContext(ctx context.Context) LookupAgentAliasResultOutput {
	return o
}

// Arn representation of the Agent Alias.
func (o LookupAgentAliasResultOutput) AgentAliasArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentAliasResult) *string { return v.AgentAliasArn }).(pulumi.StringPtrOutput)
}

// The list of history events for an alias for an Agent.
func (o LookupAgentAliasResultOutput) AgentAliasHistoryEvents() AgentAliasHistoryEventArrayOutput {
	return o.ApplyT(func(v LookupAgentAliasResult) []AgentAliasHistoryEvent { return v.AgentAliasHistoryEvents }).(AgentAliasHistoryEventArrayOutput)
}

// Id for an Agent Alias generated at the server side.
func (o LookupAgentAliasResultOutput) AgentAliasId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentAliasResult) *string { return v.AgentAliasId }).(pulumi.StringPtrOutput)
}

// Name for a resource.
func (o LookupAgentAliasResultOutput) AgentAliasName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentAliasResult) *string { return v.AgentAliasName }).(pulumi.StringPtrOutput)
}

// The status of the alias of the agent and whether it is ready for use. The following statuses are possible:
//
// - CREATING – The agent alias is being created.
// - PREPARED – The agent alias is finished being created or updated and is ready to be invoked.
// - FAILED – The agent alias API operation failed.
// - UPDATING – The agent alias is being updated.
// - DELETING – The agent alias is being deleted.
func (o LookupAgentAliasResultOutput) AgentAliasStatus() AgentAliasStatusPtrOutput {
	return o.ApplyT(func(v LookupAgentAliasResult) *AgentAliasStatus { return v.AgentAliasStatus }).(AgentAliasStatusPtrOutput)
}

// Time Stamp.
func (o LookupAgentAliasResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentAliasResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// Description of the Resource.
func (o LookupAgentAliasResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentAliasResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Routing configuration for an Agent alias.
func (o LookupAgentAliasResultOutput) RoutingConfiguration() AgentAliasRoutingConfigurationListItemArrayOutput {
	return o.ApplyT(func(v LookupAgentAliasResult) []AgentAliasRoutingConfigurationListItem { return v.RoutingConfiguration }).(AgentAliasRoutingConfigurationListItemArrayOutput)
}

// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
//
// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
func (o LookupAgentAliasResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAgentAliasResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Time Stamp.
func (o LookupAgentAliasResultOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentAliasResult) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAgentAliasResultOutput{})
}