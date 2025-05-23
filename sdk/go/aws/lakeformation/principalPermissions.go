// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lakeformation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::LakeFormation::PrincipalPermissions“ resource represents the permissions that a principal has on a GLUDC resource (such as GLUlong databases or GLUlong tables). When you create a “PrincipalPermissions“ resource, the permissions are granted via the LFlong“GrantPermissions“ API operation. When you delete a “PrincipalPermissions“ resource, the permissions on principal-resource pair are revoked via the LFlong“RevokePermissions“ API operation.
type PrincipalPermissions struct {
	pulumi.CustomResourceState

	// The identifier for the GLUDC. By default, the account ID. The GLUDC is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	Catalog pulumi.StringPtrOutput `pulumi:"catalog"`
	// The permissions granted or revoked.
	Permissions PrincipalPermissionsPermissionArrayOutput `pulumi:"permissions"`
	// Indicates the ability to grant permissions (as a subset of permissions granted).
	PermissionsWithGrantOption PrincipalPermissionsPermissionArrayOutput `pulumi:"permissionsWithGrantOption"`
	// The principal to be granted a permission.
	Principal PrincipalPermissionsDataLakePrincipalOutput `pulumi:"principal"`
	// Json encoding of the input principal. For example: `{"DataLakePrincipalIdentifier":"arn:aws:iam::123456789012:role/ExampleRole"}`
	PrincipalIdentifier pulumi.StringOutput `pulumi:"principalIdentifier"`
	// The resource to be granted or revoked permissions.
	Resource PrincipalPermissionsResourceOutput `pulumi:"resource"`
	// Json encoding of the input resource. For example: `{"Catalog":null,"Database":null,"Table":null,"TableWithColumns":null,"DataLocation":null,"DataCellsFilter":{"TableCatalogId":"123456789012","DatabaseName":"ExampleDatabase","TableName":"ExampleTable","Name":"ExampleFilter"},"LFTag":null,"LFTagPolicy":null}`
	ResourceIdentifier pulumi.StringOutput `pulumi:"resourceIdentifier"`
}

// NewPrincipalPermissions registers a new resource with the given unique name, arguments, and options.
func NewPrincipalPermissions(ctx *pulumi.Context,
	name string, args *PrincipalPermissionsArgs, opts ...pulumi.ResourceOption) (*PrincipalPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.PermissionsWithGrantOption == nil {
		return nil, errors.New("invalid value for required argument 'PermissionsWithGrantOption'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"catalog",
		"permissions[*]",
		"permissionsWithGrantOption[*]",
		"principal",
		"resource",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrincipalPermissions
	err := ctx.RegisterResource("aws-native:lakeformation:PrincipalPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrincipalPermissions gets an existing PrincipalPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrincipalPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrincipalPermissionsState, opts ...pulumi.ResourceOption) (*PrincipalPermissions, error) {
	var resource PrincipalPermissions
	err := ctx.ReadResource("aws-native:lakeformation:PrincipalPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrincipalPermissions resources.
type principalPermissionsState struct {
}

type PrincipalPermissionsState struct {
}

func (PrincipalPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*principalPermissionsState)(nil)).Elem()
}

type principalPermissionsArgs struct {
	// The identifier for the GLUDC. By default, the account ID. The GLUDC is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	Catalog *string `pulumi:"catalog"`
	// The permissions granted or revoked.
	Permissions []PrincipalPermissionsPermission `pulumi:"permissions"`
	// Indicates the ability to grant permissions (as a subset of permissions granted).
	PermissionsWithGrantOption []PrincipalPermissionsPermission `pulumi:"permissionsWithGrantOption"`
	// The principal to be granted a permission.
	Principal PrincipalPermissionsDataLakePrincipal `pulumi:"principal"`
	// The resource to be granted or revoked permissions.
	Resource PrincipalPermissionsResource `pulumi:"resource"`
}

// The set of arguments for constructing a PrincipalPermissions resource.
type PrincipalPermissionsArgs struct {
	// The identifier for the GLUDC. By default, the account ID. The GLUDC is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	Catalog pulumi.StringPtrInput
	// The permissions granted or revoked.
	Permissions PrincipalPermissionsPermissionArrayInput
	// Indicates the ability to grant permissions (as a subset of permissions granted).
	PermissionsWithGrantOption PrincipalPermissionsPermissionArrayInput
	// The principal to be granted a permission.
	Principal PrincipalPermissionsDataLakePrincipalInput
	// The resource to be granted or revoked permissions.
	Resource PrincipalPermissionsResourceInput
}

func (PrincipalPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*principalPermissionsArgs)(nil)).Elem()
}

type PrincipalPermissionsInput interface {
	pulumi.Input

	ToPrincipalPermissionsOutput() PrincipalPermissionsOutput
	ToPrincipalPermissionsOutputWithContext(ctx context.Context) PrincipalPermissionsOutput
}

func (*PrincipalPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalPermissions)(nil)).Elem()
}

func (i *PrincipalPermissions) ToPrincipalPermissionsOutput() PrincipalPermissionsOutput {
	return i.ToPrincipalPermissionsOutputWithContext(context.Background())
}

func (i *PrincipalPermissions) ToPrincipalPermissionsOutputWithContext(ctx context.Context) PrincipalPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrincipalPermissionsOutput)
}

type PrincipalPermissionsOutput struct{ *pulumi.OutputState }

func (PrincipalPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalPermissions)(nil)).Elem()
}

func (o PrincipalPermissionsOutput) ToPrincipalPermissionsOutput() PrincipalPermissionsOutput {
	return o
}

func (o PrincipalPermissionsOutput) ToPrincipalPermissionsOutputWithContext(ctx context.Context) PrincipalPermissionsOutput {
	return o
}

// The identifier for the GLUDC. By default, the account ID. The GLUDC is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
func (o PrincipalPermissionsOutput) Catalog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrincipalPermissions) pulumi.StringPtrOutput { return v.Catalog }).(pulumi.StringPtrOutput)
}

// The permissions granted or revoked.
func (o PrincipalPermissionsOutput) Permissions() PrincipalPermissionsPermissionArrayOutput {
	return o.ApplyT(func(v *PrincipalPermissions) PrincipalPermissionsPermissionArrayOutput { return v.Permissions }).(PrincipalPermissionsPermissionArrayOutput)
}

// Indicates the ability to grant permissions (as a subset of permissions granted).
func (o PrincipalPermissionsOutput) PermissionsWithGrantOption() PrincipalPermissionsPermissionArrayOutput {
	return o.ApplyT(func(v *PrincipalPermissions) PrincipalPermissionsPermissionArrayOutput {
		return v.PermissionsWithGrantOption
	}).(PrincipalPermissionsPermissionArrayOutput)
}

// The principal to be granted a permission.
func (o PrincipalPermissionsOutput) Principal() PrincipalPermissionsDataLakePrincipalOutput {
	return o.ApplyT(func(v *PrincipalPermissions) PrincipalPermissionsDataLakePrincipalOutput { return v.Principal }).(PrincipalPermissionsDataLakePrincipalOutput)
}

// Json encoding of the input principal. For example: `{"DataLakePrincipalIdentifier":"arn:aws:iam::123456789012:role/ExampleRole"}`
func (o PrincipalPermissionsOutput) PrincipalIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *PrincipalPermissions) pulumi.StringOutput { return v.PrincipalIdentifier }).(pulumi.StringOutput)
}

// The resource to be granted or revoked permissions.
func (o PrincipalPermissionsOutput) Resource() PrincipalPermissionsResourceOutput {
	return o.ApplyT(func(v *PrincipalPermissions) PrincipalPermissionsResourceOutput { return v.Resource }).(PrincipalPermissionsResourceOutput)
}

// Json encoding of the input resource. For example: `{"Catalog":null,"Database":null,"Table":null,"TableWithColumns":null,"DataLocation":null,"DataCellsFilter":{"TableCatalogId":"123456789012","DatabaseName":"ExampleDatabase","TableName":"ExampleTable","Name":"ExampleFilter"},"LFTag":null,"LFTagPolicy":null}`
func (o PrincipalPermissionsOutput) ResourceIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *PrincipalPermissions) pulumi.StringOutput { return v.ResourceIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrincipalPermissionsInput)(nil)).Elem(), &PrincipalPermissions{})
	pulumi.RegisterOutputType(PrincipalPermissionsOutput{})
}
