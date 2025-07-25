// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::Project
//
// ## Example Usage
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/sagemaker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sagemaker.NewProject(ctx, "sampleProject", &sagemaker.ProjectArgs{
//				ProjectName:        pulumi.String("project1"),
//				ProjectDescription: pulumi.String("Project Description"),
//				ServiceCatalogProvisioningDetails: &sagemaker.ServiceCatalogProvisioningDetailsPropertiesArgs{
//					ProductId:              pulumi.String("prod-53ibyqbj2cgmo"),
//					ProvisioningArtifactId: pulumi.String("pa-sm4pjfuzictpe"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/sagemaker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sagemaker.NewProject(ctx, "sampleProject", &sagemaker.ProjectArgs{
//				ProjectName:        pulumi.String("SampleProject"),
//				ProjectDescription: pulumi.String("Project Description"),
//				ServiceCatalogProvisioningDetails: &sagemaker.ServiceCatalogProvisioningDetailsPropertiesArgs{
//					ProductId:              pulumi.String("prod-53ibyqbj2cgmo"),
//					ProvisioningArtifactId: pulumi.String("pa-sm4pjfuzictpe"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Project struct {
	pulumi.CustomResourceState

	// The time at which the project was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The Amazon Resource Name (ARN) of the project.
	ProjectArn pulumi.StringOutput `pulumi:"projectArn"`
	// The description of the project.
	ProjectDescription pulumi.StringPtrOutput `pulumi:"projectDescription"`
	// The ID of the project. This ID is prepended to all entities associated with this project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The name of the project.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The status of a project.
	ProjectStatus ProjectStatusOutput `pulumi:"projectStatus"`
	// Provisioned ServiceCatalog  Details
	ServiceCatalogProvisionedProductDetails ServiceCatalogProvisionedProductDetailsPropertiesPtrOutput `pulumi:"serviceCatalogProvisionedProductDetails"`
	// Input ServiceCatalog Provisioning Details
	ServiceCatalogProvisioningDetails ServiceCatalogProvisioningDetailsPropertiesPtrOutput `pulumi:"serviceCatalogProvisioningDetails"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.CreateOnlyTagArrayOutput `pulumi:"tags"`
	// An array of template providers associated with the project.
	TemplateProviderDetails ProjectTemplateProviderDetailArrayOutput `pulumi:"templateProviderDetails"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"projectDescription",
		"projectName",
		"serviceCatalogProvisioningDetails",
		"tags[*]",
		"templateProviderDetails[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("aws-native:sagemaker:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("aws-native:sagemaker:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
}

type ProjectState struct {
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// The description of the project.
	ProjectDescription *string `pulumi:"projectDescription"`
	// The name of the project.
	ProjectName *string `pulumi:"projectName"`
	// Provisioned ServiceCatalog  Details
	ServiceCatalogProvisionedProductDetails *ServiceCatalogProvisionedProductDetailsProperties `pulumi:"serviceCatalogProvisionedProductDetails"`
	// Input ServiceCatalog Provisioning Details
	ServiceCatalogProvisioningDetails *ServiceCatalogProvisioningDetailsProperties `pulumi:"serviceCatalogProvisioningDetails"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.CreateOnlyTag `pulumi:"tags"`
	// An array of template providers associated with the project.
	TemplateProviderDetails []ProjectTemplateProviderDetail `pulumi:"templateProviderDetails"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// The description of the project.
	ProjectDescription pulumi.StringPtrInput
	// The name of the project.
	ProjectName pulumi.StringPtrInput
	// Provisioned ServiceCatalog  Details
	ServiceCatalogProvisionedProductDetails ServiceCatalogProvisionedProductDetailsPropertiesPtrInput
	// Input ServiceCatalog Provisioning Details
	ServiceCatalogProvisioningDetails ServiceCatalogProvisioningDetailsPropertiesPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.CreateOnlyTagArrayInput
	// An array of template providers associated with the project.
	TemplateProviderDetails ProjectTemplateProviderDetailArrayInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// The time at which the project was created.
func (o ProjectOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the project.
func (o ProjectOutput) ProjectArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectArn }).(pulumi.StringOutput)
}

// The description of the project.
func (o ProjectOutput) ProjectDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.ProjectDescription }).(pulumi.StringPtrOutput)
}

// The ID of the project. This ID is prepended to all entities associated with this project.
func (o ProjectOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The name of the project.
func (o ProjectOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// The status of a project.
func (o ProjectOutput) ProjectStatus() ProjectStatusOutput {
	return o.ApplyT(func(v *Project) ProjectStatusOutput { return v.ProjectStatus }).(ProjectStatusOutput)
}

// Provisioned ServiceCatalog  Details
func (o ProjectOutput) ServiceCatalogProvisionedProductDetails() ServiceCatalogProvisionedProductDetailsPropertiesPtrOutput {
	return o.ApplyT(func(v *Project) ServiceCatalogProvisionedProductDetailsPropertiesPtrOutput {
		return v.ServiceCatalogProvisionedProductDetails
	}).(ServiceCatalogProvisionedProductDetailsPropertiesPtrOutput)
}

// Input ServiceCatalog Provisioning Details
func (o ProjectOutput) ServiceCatalogProvisioningDetails() ServiceCatalogProvisioningDetailsPropertiesPtrOutput {
	return o.ApplyT(func(v *Project) ServiceCatalogProvisioningDetailsPropertiesPtrOutput {
		return v.ServiceCatalogProvisioningDetails
	}).(ServiceCatalogProvisioningDetailsPropertiesPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ProjectOutput) Tags() aws.CreateOnlyTagArrayOutput {
	return o.ApplyT(func(v *Project) aws.CreateOnlyTagArrayOutput { return v.Tags }).(aws.CreateOnlyTagArrayOutput)
}

// An array of template providers associated with the project.
func (o ProjectOutput) TemplateProviderDetails() ProjectTemplateProviderDetailArrayOutput {
	return o.ApplyT(func(v *Project) ProjectTemplateProviderDetailArrayOutput { return v.TemplateProviderDetails }).(ProjectTemplateProviderDetailArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterOutputType(ProjectOutput{})
}
