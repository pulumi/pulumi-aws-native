// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::RDS::DBProxyTargetGroup
type DbProxyTargetGroup struct {
	pulumi.CustomResourceState

	// Displays the settings that control the size and behavior of the connection pool associated with a `DBProxyTarget` .
	ConnectionPoolConfigurationInfo DbProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrOutput `pulumi:"connectionPoolConfigurationInfo"`
	// One or more DB cluster identifiers.
	DbClusterIdentifiers pulumi.StringArrayOutput `pulumi:"dbClusterIdentifiers"`
	// One or more DB instance identifiers.
	DbInstanceIdentifiers pulumi.StringArrayOutput `pulumi:"dbInstanceIdentifiers"`
	// The identifier for the proxy.
	DbProxyName pulumi.StringOutput `pulumi:"dbProxyName"`
	// The Amazon Resource Name (ARN) representing the target group.
	TargetGroupArn pulumi.StringOutput `pulumi:"targetGroupArn"`
	// The identifier for the DBProxyTargetGroup
	TargetGroupName DbProxyTargetGroupTargetGroupNameOutput `pulumi:"targetGroupName"`
}

// NewDbProxyTargetGroup registers a new resource with the given unique name, arguments, and options.
func NewDbProxyTargetGroup(ctx *pulumi.Context,
	name string, args *DbProxyTargetGroupArgs, opts ...pulumi.ResourceOption) (*DbProxyTargetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbProxyName == nil {
		return nil, errors.New("invalid value for required argument 'DbProxyName'")
	}
	if args.TargetGroupName == nil {
		return nil, errors.New("invalid value for required argument 'TargetGroupName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"dbProxyName",
		"targetGroupName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DbProxyTargetGroup
	err := ctx.RegisterResource("aws-native:rds:DbProxyTargetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbProxyTargetGroup gets an existing DbProxyTargetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbProxyTargetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbProxyTargetGroupState, opts ...pulumi.ResourceOption) (*DbProxyTargetGroup, error) {
	var resource DbProxyTargetGroup
	err := ctx.ReadResource("aws-native:rds:DbProxyTargetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbProxyTargetGroup resources.
type dbProxyTargetGroupState struct {
}

type DbProxyTargetGroupState struct {
}

func (DbProxyTargetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbProxyTargetGroupState)(nil)).Elem()
}

type dbProxyTargetGroupArgs struct {
	// Displays the settings that control the size and behavior of the connection pool associated with a `DBProxyTarget` .
	ConnectionPoolConfigurationInfo *DbProxyTargetGroupConnectionPoolConfigurationInfoFormat `pulumi:"connectionPoolConfigurationInfo"`
	// One or more DB cluster identifiers.
	DbClusterIdentifiers []string `pulumi:"dbClusterIdentifiers"`
	// One or more DB instance identifiers.
	DbInstanceIdentifiers []string `pulumi:"dbInstanceIdentifiers"`
	// The identifier for the proxy.
	DbProxyName string `pulumi:"dbProxyName"`
	// The identifier for the DBProxyTargetGroup
	TargetGroupName DbProxyTargetGroupTargetGroupName `pulumi:"targetGroupName"`
}

// The set of arguments for constructing a DbProxyTargetGroup resource.
type DbProxyTargetGroupArgs struct {
	// Displays the settings that control the size and behavior of the connection pool associated with a `DBProxyTarget` .
	ConnectionPoolConfigurationInfo DbProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrInput
	// One or more DB cluster identifiers.
	DbClusterIdentifiers pulumi.StringArrayInput
	// One or more DB instance identifiers.
	DbInstanceIdentifiers pulumi.StringArrayInput
	// The identifier for the proxy.
	DbProxyName pulumi.StringInput
	// The identifier for the DBProxyTargetGroup
	TargetGroupName DbProxyTargetGroupTargetGroupNameInput
}

func (DbProxyTargetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbProxyTargetGroupArgs)(nil)).Elem()
}

type DbProxyTargetGroupInput interface {
	pulumi.Input

	ToDbProxyTargetGroupOutput() DbProxyTargetGroupOutput
	ToDbProxyTargetGroupOutputWithContext(ctx context.Context) DbProxyTargetGroupOutput
}

func (*DbProxyTargetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DbProxyTargetGroup)(nil)).Elem()
}

func (i *DbProxyTargetGroup) ToDbProxyTargetGroupOutput() DbProxyTargetGroupOutput {
	return i.ToDbProxyTargetGroupOutputWithContext(context.Background())
}

func (i *DbProxyTargetGroup) ToDbProxyTargetGroupOutputWithContext(ctx context.Context) DbProxyTargetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbProxyTargetGroupOutput)
}

type DbProxyTargetGroupOutput struct{ *pulumi.OutputState }

func (DbProxyTargetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbProxyTargetGroup)(nil)).Elem()
}

func (o DbProxyTargetGroupOutput) ToDbProxyTargetGroupOutput() DbProxyTargetGroupOutput {
	return o
}

func (o DbProxyTargetGroupOutput) ToDbProxyTargetGroupOutputWithContext(ctx context.Context) DbProxyTargetGroupOutput {
	return o
}

// Displays the settings that control the size and behavior of the connection pool associated with a `DBProxyTarget` .
func (o DbProxyTargetGroupOutput) ConnectionPoolConfigurationInfo() DbProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrOutput {
	return o.ApplyT(func(v *DbProxyTargetGroup) DbProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrOutput {
		return v.ConnectionPoolConfigurationInfo
	}).(DbProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrOutput)
}

// One or more DB cluster identifiers.
func (o DbProxyTargetGroupOutput) DbClusterIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DbProxyTargetGroup) pulumi.StringArrayOutput { return v.DbClusterIdentifiers }).(pulumi.StringArrayOutput)
}

// One or more DB instance identifiers.
func (o DbProxyTargetGroupOutput) DbInstanceIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DbProxyTargetGroup) pulumi.StringArrayOutput { return v.DbInstanceIdentifiers }).(pulumi.StringArrayOutput)
}

// The identifier for the proxy.
func (o DbProxyTargetGroupOutput) DbProxyName() pulumi.StringOutput {
	return o.ApplyT(func(v *DbProxyTargetGroup) pulumi.StringOutput { return v.DbProxyName }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) representing the target group.
func (o DbProxyTargetGroupOutput) TargetGroupArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DbProxyTargetGroup) pulumi.StringOutput { return v.TargetGroupArn }).(pulumi.StringOutput)
}

// The identifier for the DBProxyTargetGroup
func (o DbProxyTargetGroupOutput) TargetGroupName() DbProxyTargetGroupTargetGroupNameOutput {
	return o.ApplyT(func(v *DbProxyTargetGroup) DbProxyTargetGroupTargetGroupNameOutput { return v.TargetGroupName }).(DbProxyTargetGroupTargetGroupNameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbProxyTargetGroupInput)(nil)).Elem(), &DbProxyTargetGroup{})
	pulumi.RegisterOutputType(DbProxyTargetGroupOutput{})
}
