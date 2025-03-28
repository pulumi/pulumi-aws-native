// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lookoutmetrics

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Amazon Lookout for Metrics Detector
type AnomalyDetector struct {
	pulumi.CustomResourceState

	// Configuration options for the AnomalyDetector
	AnomalyDetectorConfig AnomalyDetectorConfigOutput `pulumi:"anomalyDetectorConfig"`
	// A description for the AnomalyDetector.
	AnomalyDetectorDescription pulumi.StringPtrOutput `pulumi:"anomalyDetectorDescription"`
	// Name for the Amazon Lookout for Metrics Anomaly Detector
	AnomalyDetectorName pulumi.StringPtrOutput `pulumi:"anomalyDetectorName"`
	// The Amazon Resource Name (ARN) of the detector. For example, `arn:aws:lookoutmetrics:us-east-2:123456789012:AnomalyDetector:my-detector`
	Arn pulumi.StringOutput `pulumi:"arn"`
	// KMS key used to encrypt the AnomalyDetector data
	KmsKeyArn pulumi.StringPtrOutput `pulumi:"kmsKeyArn"`
	// List of metric sets for anomaly detection
	MetricSetList AnomalyDetectorMetricSetArrayOutput `pulumi:"metricSetList"`
}

// NewAnomalyDetector registers a new resource with the given unique name, arguments, and options.
func NewAnomalyDetector(ctx *pulumi.Context,
	name string, args *AnomalyDetectorArgs, opts ...pulumi.ResourceOption) (*AnomalyDetector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AnomalyDetectorConfig == nil {
		return nil, errors.New("invalid value for required argument 'AnomalyDetectorConfig'")
	}
	if args.MetricSetList == nil {
		return nil, errors.New("invalid value for required argument 'MetricSetList'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"anomalyDetectorName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AnomalyDetector
	err := ctx.RegisterResource("aws-native:lookoutmetrics:AnomalyDetector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnomalyDetector gets an existing AnomalyDetector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnomalyDetector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnomalyDetectorState, opts ...pulumi.ResourceOption) (*AnomalyDetector, error) {
	var resource AnomalyDetector
	err := ctx.ReadResource("aws-native:lookoutmetrics:AnomalyDetector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnomalyDetector resources.
type anomalyDetectorState struct {
}

type AnomalyDetectorState struct {
}

func (AnomalyDetectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalyDetectorState)(nil)).Elem()
}

type anomalyDetectorArgs struct {
	// Configuration options for the AnomalyDetector
	AnomalyDetectorConfig AnomalyDetectorConfig `pulumi:"anomalyDetectorConfig"`
	// A description for the AnomalyDetector.
	AnomalyDetectorDescription *string `pulumi:"anomalyDetectorDescription"`
	// Name for the Amazon Lookout for Metrics Anomaly Detector
	AnomalyDetectorName *string `pulumi:"anomalyDetectorName"`
	// KMS key used to encrypt the AnomalyDetector data
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// List of metric sets for anomaly detection
	MetricSetList []AnomalyDetectorMetricSet `pulumi:"metricSetList"`
}

// The set of arguments for constructing a AnomalyDetector resource.
type AnomalyDetectorArgs struct {
	// Configuration options for the AnomalyDetector
	AnomalyDetectorConfig AnomalyDetectorConfigInput
	// A description for the AnomalyDetector.
	AnomalyDetectorDescription pulumi.StringPtrInput
	// Name for the Amazon Lookout for Metrics Anomaly Detector
	AnomalyDetectorName pulumi.StringPtrInput
	// KMS key used to encrypt the AnomalyDetector data
	KmsKeyArn pulumi.StringPtrInput
	// List of metric sets for anomaly detection
	MetricSetList AnomalyDetectorMetricSetArrayInput
}

func (AnomalyDetectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalyDetectorArgs)(nil)).Elem()
}

type AnomalyDetectorInput interface {
	pulumi.Input

	ToAnomalyDetectorOutput() AnomalyDetectorOutput
	ToAnomalyDetectorOutputWithContext(ctx context.Context) AnomalyDetectorOutput
}

func (*AnomalyDetector) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalyDetector)(nil)).Elem()
}

func (i *AnomalyDetector) ToAnomalyDetectorOutput() AnomalyDetectorOutput {
	return i.ToAnomalyDetectorOutputWithContext(context.Background())
}

func (i *AnomalyDetector) ToAnomalyDetectorOutputWithContext(ctx context.Context) AnomalyDetectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnomalyDetectorOutput)
}

type AnomalyDetectorOutput struct{ *pulumi.OutputState }

func (AnomalyDetectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalyDetector)(nil)).Elem()
}

func (o AnomalyDetectorOutput) ToAnomalyDetectorOutput() AnomalyDetectorOutput {
	return o
}

func (o AnomalyDetectorOutput) ToAnomalyDetectorOutputWithContext(ctx context.Context) AnomalyDetectorOutput {
	return o
}

// Configuration options for the AnomalyDetector
func (o AnomalyDetectorOutput) AnomalyDetectorConfig() AnomalyDetectorConfigOutput {
	return o.ApplyT(func(v *AnomalyDetector) AnomalyDetectorConfigOutput { return v.AnomalyDetectorConfig }).(AnomalyDetectorConfigOutput)
}

// A description for the AnomalyDetector.
func (o AnomalyDetectorOutput) AnomalyDetectorDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnomalyDetector) pulumi.StringPtrOutput { return v.AnomalyDetectorDescription }).(pulumi.StringPtrOutput)
}

// Name for the Amazon Lookout for Metrics Anomaly Detector
func (o AnomalyDetectorOutput) AnomalyDetectorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnomalyDetector) pulumi.StringPtrOutput { return v.AnomalyDetectorName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the detector. For example, `arn:aws:lookoutmetrics:us-east-2:123456789012:AnomalyDetector:my-detector`
func (o AnomalyDetectorOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalyDetector) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// KMS key used to encrypt the AnomalyDetector data
func (o AnomalyDetectorOutput) KmsKeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnomalyDetector) pulumi.StringPtrOutput { return v.KmsKeyArn }).(pulumi.StringPtrOutput)
}

// List of metric sets for anomaly detection
func (o AnomalyDetectorOutput) MetricSetList() AnomalyDetectorMetricSetArrayOutput {
	return o.ApplyT(func(v *AnomalyDetector) AnomalyDetectorMetricSetArrayOutput { return v.MetricSetList }).(AnomalyDetectorMetricSetArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnomalyDetectorInput)(nil)).Elem(), &AnomalyDetector{})
	pulumi.RegisterOutputType(AnomalyDetectorOutput{})
}
