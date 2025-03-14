// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package simspaceweaver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type SimulationS3Location struct {
	// The Schema S3 bucket name.
	BucketName string `pulumi:"bucketName"`
	// This is the schema S3 object key, which includes the full path of "folders" from the bucket root to the schema.
	ObjectKey string `pulumi:"objectKey"`
}

// SimulationS3LocationInput is an input type that accepts SimulationS3LocationArgs and SimulationS3LocationOutput values.
// You can construct a concrete instance of `SimulationS3LocationInput` via:
//
//	SimulationS3LocationArgs{...}
type SimulationS3LocationInput interface {
	pulumi.Input

	ToSimulationS3LocationOutput() SimulationS3LocationOutput
	ToSimulationS3LocationOutputWithContext(context.Context) SimulationS3LocationOutput
}

type SimulationS3LocationArgs struct {
	// The Schema S3 bucket name.
	BucketName pulumi.StringInput `pulumi:"bucketName"`
	// This is the schema S3 object key, which includes the full path of "folders" from the bucket root to the schema.
	ObjectKey pulumi.StringInput `pulumi:"objectKey"`
}

func (SimulationS3LocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimulationS3Location)(nil)).Elem()
}

func (i SimulationS3LocationArgs) ToSimulationS3LocationOutput() SimulationS3LocationOutput {
	return i.ToSimulationS3LocationOutputWithContext(context.Background())
}

func (i SimulationS3LocationArgs) ToSimulationS3LocationOutputWithContext(ctx context.Context) SimulationS3LocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimulationS3LocationOutput)
}

func (i SimulationS3LocationArgs) ToSimulationS3LocationPtrOutput() SimulationS3LocationPtrOutput {
	return i.ToSimulationS3LocationPtrOutputWithContext(context.Background())
}

func (i SimulationS3LocationArgs) ToSimulationS3LocationPtrOutputWithContext(ctx context.Context) SimulationS3LocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimulationS3LocationOutput).ToSimulationS3LocationPtrOutputWithContext(ctx)
}

// SimulationS3LocationPtrInput is an input type that accepts SimulationS3LocationArgs, SimulationS3LocationPtr and SimulationS3LocationPtrOutput values.
// You can construct a concrete instance of `SimulationS3LocationPtrInput` via:
//
//	        SimulationS3LocationArgs{...}
//
//	or:
//
//	        nil
type SimulationS3LocationPtrInput interface {
	pulumi.Input

	ToSimulationS3LocationPtrOutput() SimulationS3LocationPtrOutput
	ToSimulationS3LocationPtrOutputWithContext(context.Context) SimulationS3LocationPtrOutput
}

type simulationS3LocationPtrType SimulationS3LocationArgs

func SimulationS3LocationPtr(v *SimulationS3LocationArgs) SimulationS3LocationPtrInput {
	return (*simulationS3LocationPtrType)(v)
}

func (*simulationS3LocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SimulationS3Location)(nil)).Elem()
}

func (i *simulationS3LocationPtrType) ToSimulationS3LocationPtrOutput() SimulationS3LocationPtrOutput {
	return i.ToSimulationS3LocationPtrOutputWithContext(context.Background())
}

func (i *simulationS3LocationPtrType) ToSimulationS3LocationPtrOutputWithContext(ctx context.Context) SimulationS3LocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimulationS3LocationPtrOutput)
}

type SimulationS3LocationOutput struct{ *pulumi.OutputState }

func (SimulationS3LocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimulationS3Location)(nil)).Elem()
}

func (o SimulationS3LocationOutput) ToSimulationS3LocationOutput() SimulationS3LocationOutput {
	return o
}

func (o SimulationS3LocationOutput) ToSimulationS3LocationOutputWithContext(ctx context.Context) SimulationS3LocationOutput {
	return o
}

func (o SimulationS3LocationOutput) ToSimulationS3LocationPtrOutput() SimulationS3LocationPtrOutput {
	return o.ToSimulationS3LocationPtrOutputWithContext(context.Background())
}

func (o SimulationS3LocationOutput) ToSimulationS3LocationPtrOutputWithContext(ctx context.Context) SimulationS3LocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SimulationS3Location) *SimulationS3Location {
		return &v
	}).(SimulationS3LocationPtrOutput)
}

// The Schema S3 bucket name.
func (o SimulationS3LocationOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v SimulationS3Location) string { return v.BucketName }).(pulumi.StringOutput)
}

// This is the schema S3 object key, which includes the full path of "folders" from the bucket root to the schema.
func (o SimulationS3LocationOutput) ObjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v SimulationS3Location) string { return v.ObjectKey }).(pulumi.StringOutput)
}

type SimulationS3LocationPtrOutput struct{ *pulumi.OutputState }

func (SimulationS3LocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimulationS3Location)(nil)).Elem()
}

func (o SimulationS3LocationPtrOutput) ToSimulationS3LocationPtrOutput() SimulationS3LocationPtrOutput {
	return o
}

func (o SimulationS3LocationPtrOutput) ToSimulationS3LocationPtrOutputWithContext(ctx context.Context) SimulationS3LocationPtrOutput {
	return o
}

func (o SimulationS3LocationPtrOutput) Elem() SimulationS3LocationOutput {
	return o.ApplyT(func(v *SimulationS3Location) SimulationS3Location {
		if v != nil {
			return *v
		}
		var ret SimulationS3Location
		return ret
	}).(SimulationS3LocationOutput)
}

// The Schema S3 bucket name.
func (o SimulationS3LocationPtrOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimulationS3Location) *string {
		if v == nil {
			return nil
		}
		return &v.BucketName
	}).(pulumi.StringPtrOutput)
}

// This is the schema S3 object key, which includes the full path of "folders" from the bucket root to the schema.
func (o SimulationS3LocationPtrOutput) ObjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimulationS3Location) *string {
		if v == nil {
			return nil
		}
		return &v.ObjectKey
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SimulationS3LocationInput)(nil)).Elem(), SimulationS3LocationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SimulationS3LocationPtrInput)(nil)).Elem(), SimulationS3LocationArgs{})
	pulumi.RegisterOutputType(SimulationS3LocationOutput{})
	pulumi.RegisterOutputType(SimulationS3LocationPtrOutput{})
}
