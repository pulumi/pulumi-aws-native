// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DBClusterParameterGroupTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DBClusterParameterGroupTagInput is an input type that accepts DBClusterParameterGroupTagArgs and DBClusterParameterGroupTagOutput values.
// You can construct a concrete instance of `DBClusterParameterGroupTagInput` via:
//
//	DBClusterParameterGroupTagArgs{...}
type DBClusterParameterGroupTagInput interface {
	pulumi.Input

	ToDBClusterParameterGroupTagOutput() DBClusterParameterGroupTagOutput
	ToDBClusterParameterGroupTagOutputWithContext(context.Context) DBClusterParameterGroupTagOutput
}

type DBClusterParameterGroupTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DBClusterParameterGroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterParameterGroupTag)(nil)).Elem()
}

func (i DBClusterParameterGroupTagArgs) ToDBClusterParameterGroupTagOutput() DBClusterParameterGroupTagOutput {
	return i.ToDBClusterParameterGroupTagOutputWithContext(context.Background())
}

func (i DBClusterParameterGroupTagArgs) ToDBClusterParameterGroupTagOutputWithContext(ctx context.Context) DBClusterParameterGroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterParameterGroupTagOutput)
}

// DBClusterParameterGroupTagArrayInput is an input type that accepts DBClusterParameterGroupTagArray and DBClusterParameterGroupTagArrayOutput values.
// You can construct a concrete instance of `DBClusterParameterGroupTagArrayInput` via:
//
//	DBClusterParameterGroupTagArray{ DBClusterParameterGroupTagArgs{...} }
type DBClusterParameterGroupTagArrayInput interface {
	pulumi.Input

	ToDBClusterParameterGroupTagArrayOutput() DBClusterParameterGroupTagArrayOutput
	ToDBClusterParameterGroupTagArrayOutputWithContext(context.Context) DBClusterParameterGroupTagArrayOutput
}

type DBClusterParameterGroupTagArray []DBClusterParameterGroupTagInput

func (DBClusterParameterGroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBClusterParameterGroupTag)(nil)).Elem()
}

func (i DBClusterParameterGroupTagArray) ToDBClusterParameterGroupTagArrayOutput() DBClusterParameterGroupTagArrayOutput {
	return i.ToDBClusterParameterGroupTagArrayOutputWithContext(context.Background())
}

func (i DBClusterParameterGroupTagArray) ToDBClusterParameterGroupTagArrayOutputWithContext(ctx context.Context) DBClusterParameterGroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterParameterGroupTagArrayOutput)
}

type DBClusterParameterGroupTagOutput struct{ *pulumi.OutputState }

func (DBClusterParameterGroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterParameterGroupTag)(nil)).Elem()
}

func (o DBClusterParameterGroupTagOutput) ToDBClusterParameterGroupTagOutput() DBClusterParameterGroupTagOutput {
	return o
}

func (o DBClusterParameterGroupTagOutput) ToDBClusterParameterGroupTagOutputWithContext(ctx context.Context) DBClusterParameterGroupTagOutput {
	return o
}

func (o DBClusterParameterGroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DBClusterParameterGroupTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DBClusterParameterGroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DBClusterParameterGroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type DBClusterParameterGroupTagArrayOutput struct{ *pulumi.OutputState }

func (DBClusterParameterGroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBClusterParameterGroupTag)(nil)).Elem()
}

func (o DBClusterParameterGroupTagArrayOutput) ToDBClusterParameterGroupTagArrayOutput() DBClusterParameterGroupTagArrayOutput {
	return o
}

func (o DBClusterParameterGroupTagArrayOutput) ToDBClusterParameterGroupTagArrayOutputWithContext(ctx context.Context) DBClusterParameterGroupTagArrayOutput {
	return o
}

func (o DBClusterParameterGroupTagArrayOutput) Index(i pulumi.IntInput) DBClusterParameterGroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DBClusterParameterGroupTag {
		return vs[0].([]DBClusterParameterGroupTag)[vs[1].(int)]
	}).(DBClusterParameterGroupTagOutput)
}

type DBClusterTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DBClusterTagInput is an input type that accepts DBClusterTagArgs and DBClusterTagOutput values.
// You can construct a concrete instance of `DBClusterTagInput` via:
//
//	DBClusterTagArgs{...}
type DBClusterTagInput interface {
	pulumi.Input

	ToDBClusterTagOutput() DBClusterTagOutput
	ToDBClusterTagOutputWithContext(context.Context) DBClusterTagOutput
}

type DBClusterTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DBClusterTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterTag)(nil)).Elem()
}

func (i DBClusterTagArgs) ToDBClusterTagOutput() DBClusterTagOutput {
	return i.ToDBClusterTagOutputWithContext(context.Background())
}

func (i DBClusterTagArgs) ToDBClusterTagOutputWithContext(ctx context.Context) DBClusterTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterTagOutput)
}

// DBClusterTagArrayInput is an input type that accepts DBClusterTagArray and DBClusterTagArrayOutput values.
// You can construct a concrete instance of `DBClusterTagArrayInput` via:
//
//	DBClusterTagArray{ DBClusterTagArgs{...} }
type DBClusterTagArrayInput interface {
	pulumi.Input

	ToDBClusterTagArrayOutput() DBClusterTagArrayOutput
	ToDBClusterTagArrayOutputWithContext(context.Context) DBClusterTagArrayOutput
}

type DBClusterTagArray []DBClusterTagInput

func (DBClusterTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBClusterTag)(nil)).Elem()
}

func (i DBClusterTagArray) ToDBClusterTagArrayOutput() DBClusterTagArrayOutput {
	return i.ToDBClusterTagArrayOutputWithContext(context.Background())
}

func (i DBClusterTagArray) ToDBClusterTagArrayOutputWithContext(ctx context.Context) DBClusterTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterTagArrayOutput)
}

type DBClusterTagOutput struct{ *pulumi.OutputState }

func (DBClusterTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterTag)(nil)).Elem()
}

func (o DBClusterTagOutput) ToDBClusterTagOutput() DBClusterTagOutput {
	return o
}

func (o DBClusterTagOutput) ToDBClusterTagOutputWithContext(ctx context.Context) DBClusterTagOutput {
	return o
}

func (o DBClusterTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DBClusterTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DBClusterTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DBClusterTag) string { return v.Value }).(pulumi.StringOutput)
}

type DBClusterTagArrayOutput struct{ *pulumi.OutputState }

func (DBClusterTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBClusterTag)(nil)).Elem()
}

func (o DBClusterTagArrayOutput) ToDBClusterTagArrayOutput() DBClusterTagArrayOutput {
	return o
}

func (o DBClusterTagArrayOutput) ToDBClusterTagArrayOutputWithContext(ctx context.Context) DBClusterTagArrayOutput {
	return o
}

func (o DBClusterTagArrayOutput) Index(i pulumi.IntInput) DBClusterTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DBClusterTag {
		return vs[0].([]DBClusterTag)[vs[1].(int)]
	}).(DBClusterTagOutput)
}

type DBInstanceTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DBInstanceTagInput is an input type that accepts DBInstanceTagArgs and DBInstanceTagOutput values.
// You can construct a concrete instance of `DBInstanceTagInput` via:
//
//	DBInstanceTagArgs{...}
type DBInstanceTagInput interface {
	pulumi.Input

	ToDBInstanceTagOutput() DBInstanceTagOutput
	ToDBInstanceTagOutputWithContext(context.Context) DBInstanceTagOutput
}

type DBInstanceTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DBInstanceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBInstanceTag)(nil)).Elem()
}

func (i DBInstanceTagArgs) ToDBInstanceTagOutput() DBInstanceTagOutput {
	return i.ToDBInstanceTagOutputWithContext(context.Background())
}

func (i DBInstanceTagArgs) ToDBInstanceTagOutputWithContext(ctx context.Context) DBInstanceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBInstanceTagOutput)
}

// DBInstanceTagArrayInput is an input type that accepts DBInstanceTagArray and DBInstanceTagArrayOutput values.
// You can construct a concrete instance of `DBInstanceTagArrayInput` via:
//
//	DBInstanceTagArray{ DBInstanceTagArgs{...} }
type DBInstanceTagArrayInput interface {
	pulumi.Input

	ToDBInstanceTagArrayOutput() DBInstanceTagArrayOutput
	ToDBInstanceTagArrayOutputWithContext(context.Context) DBInstanceTagArrayOutput
}

type DBInstanceTagArray []DBInstanceTagInput

func (DBInstanceTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBInstanceTag)(nil)).Elem()
}

func (i DBInstanceTagArray) ToDBInstanceTagArrayOutput() DBInstanceTagArrayOutput {
	return i.ToDBInstanceTagArrayOutputWithContext(context.Background())
}

func (i DBInstanceTagArray) ToDBInstanceTagArrayOutputWithContext(ctx context.Context) DBInstanceTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBInstanceTagArrayOutput)
}

type DBInstanceTagOutput struct{ *pulumi.OutputState }

func (DBInstanceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBInstanceTag)(nil)).Elem()
}

func (o DBInstanceTagOutput) ToDBInstanceTagOutput() DBInstanceTagOutput {
	return o
}

func (o DBInstanceTagOutput) ToDBInstanceTagOutputWithContext(ctx context.Context) DBInstanceTagOutput {
	return o
}

func (o DBInstanceTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DBInstanceTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DBInstanceTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DBInstanceTag) string { return v.Value }).(pulumi.StringOutput)
}

type DBInstanceTagArrayOutput struct{ *pulumi.OutputState }

func (DBInstanceTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBInstanceTag)(nil)).Elem()
}

func (o DBInstanceTagArrayOutput) ToDBInstanceTagArrayOutput() DBInstanceTagArrayOutput {
	return o
}

func (o DBInstanceTagArrayOutput) ToDBInstanceTagArrayOutputWithContext(ctx context.Context) DBInstanceTagArrayOutput {
	return o
}

func (o DBInstanceTagArrayOutput) Index(i pulumi.IntInput) DBInstanceTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DBInstanceTag {
		return vs[0].([]DBInstanceTag)[vs[1].(int)]
	}).(DBInstanceTagOutput)
}

type DBSubnetGroupTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DBSubnetGroupTagInput is an input type that accepts DBSubnetGroupTagArgs and DBSubnetGroupTagOutput values.
// You can construct a concrete instance of `DBSubnetGroupTagInput` via:
//
//	DBSubnetGroupTagArgs{...}
type DBSubnetGroupTagInput interface {
	pulumi.Input

	ToDBSubnetGroupTagOutput() DBSubnetGroupTagOutput
	ToDBSubnetGroupTagOutputWithContext(context.Context) DBSubnetGroupTagOutput
}

type DBSubnetGroupTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DBSubnetGroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBSubnetGroupTag)(nil)).Elem()
}

func (i DBSubnetGroupTagArgs) ToDBSubnetGroupTagOutput() DBSubnetGroupTagOutput {
	return i.ToDBSubnetGroupTagOutputWithContext(context.Background())
}

func (i DBSubnetGroupTagArgs) ToDBSubnetGroupTagOutputWithContext(ctx context.Context) DBSubnetGroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBSubnetGroupTagOutput)
}

// DBSubnetGroupTagArrayInput is an input type that accepts DBSubnetGroupTagArray and DBSubnetGroupTagArrayOutput values.
// You can construct a concrete instance of `DBSubnetGroupTagArrayInput` via:
//
//	DBSubnetGroupTagArray{ DBSubnetGroupTagArgs{...} }
type DBSubnetGroupTagArrayInput interface {
	pulumi.Input

	ToDBSubnetGroupTagArrayOutput() DBSubnetGroupTagArrayOutput
	ToDBSubnetGroupTagArrayOutputWithContext(context.Context) DBSubnetGroupTagArrayOutput
}

type DBSubnetGroupTagArray []DBSubnetGroupTagInput

func (DBSubnetGroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBSubnetGroupTag)(nil)).Elem()
}

func (i DBSubnetGroupTagArray) ToDBSubnetGroupTagArrayOutput() DBSubnetGroupTagArrayOutput {
	return i.ToDBSubnetGroupTagArrayOutputWithContext(context.Background())
}

func (i DBSubnetGroupTagArray) ToDBSubnetGroupTagArrayOutputWithContext(ctx context.Context) DBSubnetGroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBSubnetGroupTagArrayOutput)
}

type DBSubnetGroupTagOutput struct{ *pulumi.OutputState }

func (DBSubnetGroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBSubnetGroupTag)(nil)).Elem()
}

func (o DBSubnetGroupTagOutput) ToDBSubnetGroupTagOutput() DBSubnetGroupTagOutput {
	return o
}

func (o DBSubnetGroupTagOutput) ToDBSubnetGroupTagOutputWithContext(ctx context.Context) DBSubnetGroupTagOutput {
	return o
}

func (o DBSubnetGroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DBSubnetGroupTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DBSubnetGroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DBSubnetGroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type DBSubnetGroupTagArrayOutput struct{ *pulumi.OutputState }

func (DBSubnetGroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBSubnetGroupTag)(nil)).Elem()
}

func (o DBSubnetGroupTagArrayOutput) ToDBSubnetGroupTagArrayOutput() DBSubnetGroupTagArrayOutput {
	return o
}

func (o DBSubnetGroupTagArrayOutput) ToDBSubnetGroupTagArrayOutputWithContext(ctx context.Context) DBSubnetGroupTagArrayOutput {
	return o
}

func (o DBSubnetGroupTagArrayOutput) Index(i pulumi.IntInput) DBSubnetGroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DBSubnetGroupTag {
		return vs[0].([]DBSubnetGroupTag)[vs[1].(int)]
	}).(DBSubnetGroupTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterParameterGroupTagInput)(nil)).Elem(), DBClusterParameterGroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterParameterGroupTagArrayInput)(nil)).Elem(), DBClusterParameterGroupTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterTagInput)(nil)).Elem(), DBClusterTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterTagArrayInput)(nil)).Elem(), DBClusterTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBInstanceTagInput)(nil)).Elem(), DBInstanceTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBInstanceTagArrayInput)(nil)).Elem(), DBInstanceTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBSubnetGroupTagInput)(nil)).Elem(), DBSubnetGroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBSubnetGroupTagArrayInput)(nil)).Elem(), DBSubnetGroupTagArray{})
	pulumi.RegisterOutputType(DBClusterParameterGroupTagOutput{})
	pulumi.RegisterOutputType(DBClusterParameterGroupTagArrayOutput{})
	pulumi.RegisterOutputType(DBClusterTagOutput{})
	pulumi.RegisterOutputType(DBClusterTagArrayOutput{})
	pulumi.RegisterOutputType(DBInstanceTagOutput{})
	pulumi.RegisterOutputType(DBInstanceTagArrayOutput{})
	pulumi.RegisterOutputType(DBSubnetGroupTagOutput{})
	pulumi.RegisterOutputType(DBSubnetGroupTagArrayOutput{})
}