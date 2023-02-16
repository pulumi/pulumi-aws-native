// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DocDB::DBInstance
//
// Deprecated: DBInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DBInstance struct {
	pulumi.CustomResourceState

	AutoMinorVersionUpgrade    pulumi.BoolPtrOutput     `pulumi:"autoMinorVersionUpgrade"`
	AvailabilityZone           pulumi.StringPtrOutput   `pulumi:"availabilityZone"`
	DBClusterIdentifier        pulumi.StringOutput      `pulumi:"dBClusterIdentifier"`
	DBInstanceClass            pulumi.StringOutput      `pulumi:"dBInstanceClass"`
	DBInstanceIdentifier       pulumi.StringPtrOutput   `pulumi:"dBInstanceIdentifier"`
	EnablePerformanceInsights  pulumi.BoolPtrOutput     `pulumi:"enablePerformanceInsights"`
	Endpoint                   pulumi.StringOutput      `pulumi:"endpoint"`
	Port                       pulumi.StringOutput      `pulumi:"port"`
	PreferredMaintenanceWindow pulumi.StringPtrOutput   `pulumi:"preferredMaintenanceWindow"`
	Tags                       DBInstanceTagArrayOutput `pulumi:"tags"`
}

// NewDBInstance registers a new resource with the given unique name, arguments, and options.
func NewDBInstance(ctx *pulumi.Context,
	name string, args *DBInstanceArgs, opts ...pulumi.ResourceOption) (*DBInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DBClusterIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'DBClusterIdentifier'")
	}
	if args.DBInstanceClass == nil {
		return nil, errors.New("invalid value for required argument 'DBInstanceClass'")
	}
	var resource DBInstance
	err := ctx.RegisterResource("aws-native:docdb:DBInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBInstance gets an existing DBInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBInstanceState, opts ...pulumi.ResourceOption) (*DBInstance, error) {
	var resource DBInstance
	err := ctx.ReadResource("aws-native:docdb:DBInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBInstance resources.
type dbinstanceState struct {
}

type DBInstanceState struct {
}

func (DBInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbinstanceState)(nil)).Elem()
}

type dbinstanceArgs struct {
	AutoMinorVersionUpgrade    *bool           `pulumi:"autoMinorVersionUpgrade"`
	AvailabilityZone           *string         `pulumi:"availabilityZone"`
	DBClusterIdentifier        string          `pulumi:"dBClusterIdentifier"`
	DBInstanceClass            string          `pulumi:"dBInstanceClass"`
	DBInstanceIdentifier       *string         `pulumi:"dBInstanceIdentifier"`
	EnablePerformanceInsights  *bool           `pulumi:"enablePerformanceInsights"`
	PreferredMaintenanceWindow *string         `pulumi:"preferredMaintenanceWindow"`
	Tags                       []DBInstanceTag `pulumi:"tags"`
}

// The set of arguments for constructing a DBInstance resource.
type DBInstanceArgs struct {
	AutoMinorVersionUpgrade    pulumi.BoolPtrInput
	AvailabilityZone           pulumi.StringPtrInput
	DBClusterIdentifier        pulumi.StringInput
	DBInstanceClass            pulumi.StringInput
	DBInstanceIdentifier       pulumi.StringPtrInput
	EnablePerformanceInsights  pulumi.BoolPtrInput
	PreferredMaintenanceWindow pulumi.StringPtrInput
	Tags                       DBInstanceTagArrayInput
}

func (DBInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbinstanceArgs)(nil)).Elem()
}

type DBInstanceInput interface {
	pulumi.Input

	ToDBInstanceOutput() DBInstanceOutput
	ToDBInstanceOutputWithContext(ctx context.Context) DBInstanceOutput
}

func (*DBInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**DBInstance)(nil)).Elem()
}

func (i *DBInstance) ToDBInstanceOutput() DBInstanceOutput {
	return i.ToDBInstanceOutputWithContext(context.Background())
}

func (i *DBInstance) ToDBInstanceOutputWithContext(ctx context.Context) DBInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBInstanceOutput)
}

type DBInstanceOutput struct{ *pulumi.OutputState }

func (DBInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DBInstance)(nil)).Elem()
}

func (o DBInstanceOutput) ToDBInstanceOutput() DBInstanceOutput {
	return o
}

func (o DBInstanceOutput) ToDBInstanceOutputWithContext(ctx context.Context) DBInstanceOutput {
	return o
}

func (o DBInstanceOutput) AutoMinorVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.AutoMinorVersionUpgrade }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) DBClusterIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringOutput { return v.DBClusterIdentifier }).(pulumi.StringOutput)
}

func (o DBInstanceOutput) DBInstanceClass() pulumi.StringOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringOutput { return v.DBInstanceClass }).(pulumi.StringOutput)
}

func (o DBInstanceOutput) DBInstanceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.DBInstanceIdentifier }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) EnablePerformanceInsights() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.BoolPtrOutput { return v.EnablePerformanceInsights }).(pulumi.BoolPtrOutput)
}

func (o DBInstanceOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

func (o DBInstanceOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

func (o DBInstanceOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBInstance) pulumi.StringPtrOutput { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o DBInstanceOutput) Tags() DBInstanceTagArrayOutput {
	return o.ApplyT(func(v *DBInstance) DBInstanceTagArrayOutput { return v.Tags }).(DBInstanceTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DBInstanceInput)(nil)).Elem(), &DBInstance{})
	pulumi.RegisterOutputType(DBInstanceOutput{})
}