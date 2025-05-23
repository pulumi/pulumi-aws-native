// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Neptune::DBClusterParameterGroup resource creates a new Amazon Neptune DB cluster parameter group
type DbClusterParameterGroup struct {
	pulumi.CustomResourceState

	// Provides the customer-specified description for this DB cluster parameter group.
	Description pulumi.StringOutput `pulumi:"description"`
	// Must be neptune1 for engine versions prior to 1.2.0.0, or neptune1.2 for engine version 1.2.0.0 and higher.
	Family pulumi.StringOutput `pulumi:"family"`
	// Provides the name of the DB cluster parameter group.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Neptune::DBClusterParameterGroup` for more information about the expected schema for this property.
	Parameters pulumi.AnyOutput `pulumi:"parameters"`
	// The list of tags for the cluster parameter group.
	Tags aws.TagArrayOutput `pulumi:"tags"`
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
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"description",
		"family",
		"name",
	})
	opts = append(opts, replaceOnChanges)
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
	// Provides the customer-specified description for this DB cluster parameter group.
	Description string `pulumi:"description"`
	// Must be neptune1 for engine versions prior to 1.2.0.0, or neptune1.2 for engine version 1.2.0.0 and higher.
	Family string `pulumi:"family"`
	// Provides the name of the DB cluster parameter group.
	Name *string `pulumi:"name"`
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Neptune::DBClusterParameterGroup` for more information about the expected schema for this property.
	Parameters interface{} `pulumi:"parameters"`
	// The list of tags for the cluster parameter group.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a DbClusterParameterGroup resource.
type DbClusterParameterGroupArgs struct {
	// Provides the customer-specified description for this DB cluster parameter group.
	Description pulumi.StringInput
	// Must be neptune1 for engine versions prior to 1.2.0.0, or neptune1.2 for engine version 1.2.0.0 and higher.
	Family pulumi.StringInput
	// Provides the name of the DB cluster parameter group.
	Name pulumi.StringPtrInput
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Neptune::DBClusterParameterGroup` for more information about the expected schema for this property.
	Parameters pulumi.Input
	// The list of tags for the cluster parameter group.
	Tags aws.TagArrayInput
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

// Provides the customer-specified description for this DB cluster parameter group.
func (o DbClusterParameterGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DbClusterParameterGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Must be neptune1 for engine versions prior to 1.2.0.0, or neptune1.2 for engine version 1.2.0.0 and higher.
func (o DbClusterParameterGroupOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v *DbClusterParameterGroup) pulumi.StringOutput { return v.Family }).(pulumi.StringOutput)
}

// Provides the name of the DB cluster parameter group.
func (o DbClusterParameterGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DbClusterParameterGroup) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Neptune::DBClusterParameterGroup` for more information about the expected schema for this property.
func (o DbClusterParameterGroupOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *DbClusterParameterGroup) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

// The list of tags for the cluster parameter group.
func (o DbClusterParameterGroupOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *DbClusterParameterGroup) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterParameterGroupInput)(nil)).Elem(), &DbClusterParameterGroup{})
	pulumi.RegisterOutputType(DbClusterParameterGroupOutput{})
}
