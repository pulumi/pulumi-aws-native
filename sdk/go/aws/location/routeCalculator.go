// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Location::RouteCalculator Resource Type
type RouteCalculator struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) for the route calculator resource. Use the ARN when you specify a resource across all AWS .
	//
	// - Format example: `arn:aws:geo:region:account-id:route-calculator/ExampleCalculator`
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Synonym for `Arn` . The Amazon Resource Name (ARN) for the route calculator resource. Use the ARN when you specify a resource across all AWS .
	//
	// - Format example: `arn:aws:geo:region:account-id:route-calculator/ExampleCalculator`
	CalculatorArn pulumi.StringOutput `pulumi:"calculatorArn"`
	// The name of the route calculator resource.
	//
	// Requirements:
	//
	// - Can use alphanumeric characters (A–Z, a–z, 0–9) , hyphens (-), periods (.), and underscores (_).
	// - Must be a unique Route calculator resource name.
	// - No spaces allowed. For example, `ExampleRouteCalculator` .
	CalculatorName pulumi.StringOutput `pulumi:"calculatorName"`
	// The timestamp for when the route calculator resource was created in [ISO 8601](https://docs.aws.amazon.com/https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ` .
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Specifies the data provider of traffic and road network data.
	//
	// > This field is case-sensitive. Enter the valid values as shown. For example, entering `HERE` returns an error.
	//
	// Valid values include:
	//
	// - `Esri` – For additional information about [Esri](https://docs.aws.amazon.com/location/previous/developerguide/esri.html) 's coverage in your region of interest, see [Esri details on street networks and traffic coverage](https://docs.aws.amazon.com/https://doc.arcgis.com/en/arcgis-online/reference/network-coverage.htm) .
	//
	// Route calculators that use Esri as a data source only calculate routes that are shorter than 400 km.
	// - `Grab` – Grab provides routing functionality for Southeast Asia. For additional information about [GrabMaps](https://docs.aws.amazon.com/location/previous/developerguide/grab.html) ' coverage, see [GrabMaps countries and areas covered](https://docs.aws.amazon.com/location/previous/developerguide/grab.html#grab-coverage-area) .
	// - `Here` – For additional information about [HERE Technologies](https://docs.aws.amazon.com/location/previous/developerguide/HERE.html) ' coverage in your region of interest, see [HERE car routing coverage](https://docs.aws.amazon.com/https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/car-routing.html) and [HERE truck routing coverage](https://docs.aws.amazon.com/https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/truck-routing.html) .
	//
	// For additional information , see [Data providers](https://docs.aws.amazon.com/location/previous/developerguide/what-is-data-provider.html) on the *Amazon Location Service Developer Guide* .
	DataSource pulumi.StringOutput `pulumi:"dataSource"`
	// The optional description for the route calculator resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// No longer used. If included, the only allowed value is `RequestBasedUsage` .
	//
	// *Allowed Values* : `RequestBasedUsage`
	PricingPlan RouteCalculatorPricingPlanPtrOutput `pulumi:"pricingPlan"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The timestamp for when the route calculator resource was last updated in [ISO 8601](https://docs.aws.amazon.com/https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ` .
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewRouteCalculator registers a new resource with the given unique name, arguments, and options.
func NewRouteCalculator(ctx *pulumi.Context,
	name string, args *RouteCalculatorArgs, opts ...pulumi.ResourceOption) (*RouteCalculator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataSource == nil {
		return nil, errors.New("invalid value for required argument 'DataSource'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"calculatorName",
		"dataSource",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouteCalculator
	err := ctx.RegisterResource("aws-native:location:RouteCalculator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteCalculator gets an existing RouteCalculator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteCalculator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteCalculatorState, opts ...pulumi.ResourceOption) (*RouteCalculator, error) {
	var resource RouteCalculator
	err := ctx.ReadResource("aws-native:location:RouteCalculator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteCalculator resources.
type routeCalculatorState struct {
}

type RouteCalculatorState struct {
}

func (RouteCalculatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeCalculatorState)(nil)).Elem()
}

type routeCalculatorArgs struct {
	// The name of the route calculator resource.
	//
	// Requirements:
	//
	// - Can use alphanumeric characters (A–Z, a–z, 0–9) , hyphens (-), periods (.), and underscores (_).
	// - Must be a unique Route calculator resource name.
	// - No spaces allowed. For example, `ExampleRouteCalculator` .
	CalculatorName *string `pulumi:"calculatorName"`
	// Specifies the data provider of traffic and road network data.
	//
	// > This field is case-sensitive. Enter the valid values as shown. For example, entering `HERE` returns an error.
	//
	// Valid values include:
	//
	// - `Esri` – For additional information about [Esri](https://docs.aws.amazon.com/location/previous/developerguide/esri.html) 's coverage in your region of interest, see [Esri details on street networks and traffic coverage](https://docs.aws.amazon.com/https://doc.arcgis.com/en/arcgis-online/reference/network-coverage.htm) .
	//
	// Route calculators that use Esri as a data source only calculate routes that are shorter than 400 km.
	// - `Grab` – Grab provides routing functionality for Southeast Asia. For additional information about [GrabMaps](https://docs.aws.amazon.com/location/previous/developerguide/grab.html) ' coverage, see [GrabMaps countries and areas covered](https://docs.aws.amazon.com/location/previous/developerguide/grab.html#grab-coverage-area) .
	// - `Here` – For additional information about [HERE Technologies](https://docs.aws.amazon.com/location/previous/developerguide/HERE.html) ' coverage in your region of interest, see [HERE car routing coverage](https://docs.aws.amazon.com/https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/car-routing.html) and [HERE truck routing coverage](https://docs.aws.amazon.com/https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/truck-routing.html) .
	//
	// For additional information , see [Data providers](https://docs.aws.amazon.com/location/previous/developerguide/what-is-data-provider.html) on the *Amazon Location Service Developer Guide* .
	DataSource string `pulumi:"dataSource"`
	// The optional description for the route calculator resource.
	Description *string `pulumi:"description"`
	// No longer used. If included, the only allowed value is `RequestBasedUsage` .
	//
	// *Allowed Values* : `RequestBasedUsage`
	PricingPlan *RouteCalculatorPricingPlan `pulumi:"pricingPlan"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a RouteCalculator resource.
type RouteCalculatorArgs struct {
	// The name of the route calculator resource.
	//
	// Requirements:
	//
	// - Can use alphanumeric characters (A–Z, a–z, 0–9) , hyphens (-), periods (.), and underscores (_).
	// - Must be a unique Route calculator resource name.
	// - No spaces allowed. For example, `ExampleRouteCalculator` .
	CalculatorName pulumi.StringPtrInput
	// Specifies the data provider of traffic and road network data.
	//
	// > This field is case-sensitive. Enter the valid values as shown. For example, entering `HERE` returns an error.
	//
	// Valid values include:
	//
	// - `Esri` – For additional information about [Esri](https://docs.aws.amazon.com/location/previous/developerguide/esri.html) 's coverage in your region of interest, see [Esri details on street networks and traffic coverage](https://docs.aws.amazon.com/https://doc.arcgis.com/en/arcgis-online/reference/network-coverage.htm) .
	//
	// Route calculators that use Esri as a data source only calculate routes that are shorter than 400 km.
	// - `Grab` – Grab provides routing functionality for Southeast Asia. For additional information about [GrabMaps](https://docs.aws.amazon.com/location/previous/developerguide/grab.html) ' coverage, see [GrabMaps countries and areas covered](https://docs.aws.amazon.com/location/previous/developerguide/grab.html#grab-coverage-area) .
	// - `Here` – For additional information about [HERE Technologies](https://docs.aws.amazon.com/location/previous/developerguide/HERE.html) ' coverage in your region of interest, see [HERE car routing coverage](https://docs.aws.amazon.com/https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/car-routing.html) and [HERE truck routing coverage](https://docs.aws.amazon.com/https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/truck-routing.html) .
	//
	// For additional information , see [Data providers](https://docs.aws.amazon.com/location/previous/developerguide/what-is-data-provider.html) on the *Amazon Location Service Developer Guide* .
	DataSource pulumi.StringInput
	// The optional description for the route calculator resource.
	Description pulumi.StringPtrInput
	// No longer used. If included, the only allowed value is `RequestBasedUsage` .
	//
	// *Allowed Values* : `RequestBasedUsage`
	PricingPlan RouteCalculatorPricingPlanPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayInput
}

func (RouteCalculatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeCalculatorArgs)(nil)).Elem()
}

type RouteCalculatorInput interface {
	pulumi.Input

	ToRouteCalculatorOutput() RouteCalculatorOutput
	ToRouteCalculatorOutputWithContext(ctx context.Context) RouteCalculatorOutput
}

func (*RouteCalculator) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteCalculator)(nil)).Elem()
}

func (i *RouteCalculator) ToRouteCalculatorOutput() RouteCalculatorOutput {
	return i.ToRouteCalculatorOutputWithContext(context.Background())
}

func (i *RouteCalculator) ToRouteCalculatorOutputWithContext(ctx context.Context) RouteCalculatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteCalculatorOutput)
}

type RouteCalculatorOutput struct{ *pulumi.OutputState }

func (RouteCalculatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteCalculator)(nil)).Elem()
}

func (o RouteCalculatorOutput) ToRouteCalculatorOutput() RouteCalculatorOutput {
	return o
}

func (o RouteCalculatorOutput) ToRouteCalculatorOutputWithContext(ctx context.Context) RouteCalculatorOutput {
	return o
}

// The Amazon Resource Name (ARN) for the route calculator resource. Use the ARN when you specify a resource across all AWS .
//
// - Format example: `arn:aws:geo:region:account-id:route-calculator/ExampleCalculator`
func (o RouteCalculatorOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteCalculator) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Synonym for `Arn` . The Amazon Resource Name (ARN) for the route calculator resource. Use the ARN when you specify a resource across all AWS .
//
// - Format example: `arn:aws:geo:region:account-id:route-calculator/ExampleCalculator`
func (o RouteCalculatorOutput) CalculatorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteCalculator) pulumi.StringOutput { return v.CalculatorArn }).(pulumi.StringOutput)
}

// The name of the route calculator resource.
//
// Requirements:
//
// - Can use alphanumeric characters (A–Z, a–z, 0–9) , hyphens (-), periods (.), and underscores (_).
// - Must be a unique Route calculator resource name.
// - No spaces allowed. For example, `ExampleRouteCalculator` .
func (o RouteCalculatorOutput) CalculatorName() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteCalculator) pulumi.StringOutput { return v.CalculatorName }).(pulumi.StringOutput)
}

// The timestamp for when the route calculator resource was created in [ISO 8601](https://docs.aws.amazon.com/https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ` .
func (o RouteCalculatorOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteCalculator) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Specifies the data provider of traffic and road network data.
//
// > This field is case-sensitive. Enter the valid values as shown. For example, entering `HERE` returns an error.
//
// Valid values include:
//
// - `Esri` – For additional information about [Esri](https://docs.aws.amazon.com/location/previous/developerguide/esri.html) 's coverage in your region of interest, see [Esri details on street networks and traffic coverage](https://docs.aws.amazon.com/https://doc.arcgis.com/en/arcgis-online/reference/network-coverage.htm) .
//
// Route calculators that use Esri as a data source only calculate routes that are shorter than 400 km.
// - `Grab` – Grab provides routing functionality for Southeast Asia. For additional information about [GrabMaps](https://docs.aws.amazon.com/location/previous/developerguide/grab.html) ' coverage, see [GrabMaps countries and areas covered](https://docs.aws.amazon.com/location/previous/developerguide/grab.html#grab-coverage-area) .
// - `Here` – For additional information about [HERE Technologies](https://docs.aws.amazon.com/location/previous/developerguide/HERE.html) ' coverage in your region of interest, see [HERE car routing coverage](https://docs.aws.amazon.com/https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/car-routing.html) and [HERE truck routing coverage](https://docs.aws.amazon.com/https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/truck-routing.html) .
//
// For additional information , see [Data providers](https://docs.aws.amazon.com/location/previous/developerguide/what-is-data-provider.html) on the *Amazon Location Service Developer Guide* .
func (o RouteCalculatorOutput) DataSource() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteCalculator) pulumi.StringOutput { return v.DataSource }).(pulumi.StringOutput)
}

// The optional description for the route calculator resource.
func (o RouteCalculatorOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteCalculator) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// No longer used. If included, the only allowed value is `RequestBasedUsage` .
//
// *Allowed Values* : `RequestBasedUsage`
func (o RouteCalculatorOutput) PricingPlan() RouteCalculatorPricingPlanPtrOutput {
	return o.ApplyT(func(v *RouteCalculator) RouteCalculatorPricingPlanPtrOutput { return v.PricingPlan }).(RouteCalculatorPricingPlanPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o RouteCalculatorOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *RouteCalculator) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The timestamp for when the route calculator resource was last updated in [ISO 8601](https://docs.aws.amazon.com/https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ` .
func (o RouteCalculatorOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteCalculator) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteCalculatorInput)(nil)).Elem(), &RouteCalculator{})
	pulumi.RegisterOutputType(RouteCalculatorOutput{})
}
