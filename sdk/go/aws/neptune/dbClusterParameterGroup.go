// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Neptune::DBClusterParameterGroup
//
// Deprecated: DbClusterParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DbClusterParameterGroup struct {
	pulumi.CustomResourceState

	Description pulumi.StringOutput                   `pulumi:"description"`
	Family      pulumi.StringOutput                   `pulumi:"family"`
	Name        pulumi.StringPtrOutput                `pulumi:"name"`
	Parameters  pulumi.AnyOutput                      `pulumi:"parameters"`
	Tags        DbClusterParameterGroupTagArrayOutput `pulumi:"tags"`
}

// NewDbClusterParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewDbClusterParameterGroup(ctx *pulumi.Context,
	name string, args *DbClusterParameterGroupArgs, opts ...pulumi.ResourceOption) (*DbClusterParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Family == nil {
		return nil, errors.New("invalid value for required argument 'Family'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DbClusterParameterGroup
	err := ctx.RegisterResource("aws-native:neptune:DbClusterParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbClusterParameterGroup gets an existing DbClusterParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbClusterParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbClusterParameterGroupState, opts ...pulumi.ResourceOption) (*DbClusterParameterGroup, error) {
	var resource DbClusterParameterGroup
	err := ctx.ReadResource("aws-native:neptune:DbClusterParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbClusterParameterGroup resources.
type dbClusterParameterGroupState struct {
}

type DbClusterParameterGroupState struct {
}

func (DbClusterParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbClusterParameterGroupState)(nil)).Elem()
}

type dbClusterParameterGroupArgs struct {
	Description string                       `pulumi:"description"`
	Family      string                       `pulumi:"family"`
	Name        *string                      `pulumi:"name"`
	Parameters  interface{}                  `pulumi:"parameters"`
	Tags        []DbClusterParameterGroupTag `pulumi:"tags"`
}

// The set of arguments for constructing a DbClusterParameterGroup resource.
type DbClusterParameterGroupArgs struct {
	Description pulumi.StringInput
	Family      pulumi.StringInput
	Name        pulumi.StringPtrInput
	Parameters  pulumi.Input
	Tags        DbClusterParameterGroupTagArrayInput
}

func (DbClusterParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbClusterParameterGroupArgs)(nil)).Elem()
}

type DbClusterParameterGroupInput interface {
	pulumi.Input

	ToDbClusterParameterGroupOutput() DbClusterParameterGroupOutput
	ToDbClusterParameterGroupOutputWithContext(ctx context.Context) DbClusterParameterGroupOutput
}

func (*DbClusterParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DbClusterParameterGroup)(nil)).Elem()
}

func (i *DbClusterParameterGroup) ToDbClusterParameterGroupOutput() DbClusterParameterGroupOutput {
	return i.ToDbClusterParameterGroupOutputWithContext(context.Background())
}

func (i *DbClusterParameterGroup) ToDbClusterParameterGroupOutputWithContext(ctx context.Context) DbClusterParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterParameterGroupOutput)
}

type DbClusterParameterGroupOutput struct{ *pulumi.OutputState }

func (DbClusterParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbClusterParameterGroup)(nil)).Elem()
}

func (o DbClusterParameterGroupOutput) ToDbClusterParameterGroupOutput() DbClusterParameterGroupOutput {
	return o
}

func (o DbClusterParameterGroupOutput) ToDbClusterParameterGroupOutputWithContext(ctx context.Context) DbClusterParameterGroupOutput {
	return o
}

func (o DbClusterParameterGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DbClusterParameterGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o DbClusterParameterGroupOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v *DbClusterParameterGroup) pulumi.StringOutput { return v.Family }).(pulumi.StringOutput)
}

func (o DbClusterParameterGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DbClusterParameterGroup) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DbClusterParameterGroupOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *DbClusterParameterGroup) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

func (o DbClusterParameterGroupOutput) Tags() DbClusterParameterGroupTagArrayOutput {
	return o.ApplyT(func(v *DbClusterParameterGroup) DbClusterParameterGroupTagArrayOutput { return v.Tags }).(DbClusterParameterGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterParameterGroupInput)(nil)).Elem(), &DbClusterParameterGroup{})
	pulumi.RegisterOutputType(DbClusterParameterGroupOutput{})
}