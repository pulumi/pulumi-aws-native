// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package secretsmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new secret. A *secret* can be a password, a set of credentials such as a user name and password, an OAuth token, or other secret information that you store in an encrypted form in Secrets Manager.
//
//	For RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-masterusersecret.html).
//	For RS admin user credentials, see [AWS::Redshift::Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html).
//	To retrieve a secret in a CFNshort template, use a *dynamic reference*. For more information, see [Retrieve a secret in an resource](https://docs.aws.amazon.com/secretsmanager/latest/userguide/cfn-example_reference-secret.html).
//	For information about creating a secret in the console, see [Create a secret](https://docs.aws.amazon.com/secretsmanager/latest/userguide/manage_create-basic-secret.html). For information about creating a secret using the CLI or SDK, see [CreateSecret](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_CreateSecret.html).
//	For information about retrieving a secret in code, see [Retrieve secrets from Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieving-secrets.html).
func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecretResult
	err := ctx.Invoke("aws-native:secretsmanager:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecretArgs struct {
	// The ARN of the secret.
	Id string `pulumi:"id"`
}

type LookupSecretResult struct {
	// The description of the secret.
	Description *string `pulumi:"description"`
	// The ARN of the secret.
	Id *string `pulumi:"id"`
	// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by ``alias/``, for example ``alias/aws/secretsmanager``. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).
	//  To use a KMS key in a different account, use the key ARN or the alias ARN.
	//  If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	//  If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A custom type that specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.
	ReplicaRegions []SecretReplicaRegion `pulumi:"replicaRegions"`
	// A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:
	//   ``[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]``
	//  Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
	//  Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret.
	//  If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2).
	//  For information about how to format a JSON parameter for the various command line tool environments, see [Using JSON for Parameters](https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json). If your command-line tool or SDK requires quotation marks around the parameter, you should use single quotes to avoid confusion with the double quotes required in the JSON text.
	//  The following restrictions apply to tags:
	//   +  Maximum number of tags per secret: 50
	//   +  Maximum key length: 127 Unicode characters in UTF-8
	//   +  Maximum value length: 255 Unicode characters in UTF-8
	//   +  Tag keys and values are case sensitive.
	//   +  Do not use the ``aws:`` prefix in your tag names or values because AWS reserves it for AWS use. You can't edit or delete tag names or values with this prefix. Tags with this prefix do not count against your tags per secret limit.
	//   +  If you use your tagging schema across multiple services and resources, other services might have restrictions on allowed characters. Generally allowed characters: letters, spaces, and numbers representable in UTF-8, plus the following special characters: + - = . _ : / @.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupSecretOutput(ctx *pulumi.Context, args LookupSecretOutputArgs, opts ...pulumi.InvokeOption) LookupSecretResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSecretResultOutput, error) {
			args := v.(LookupSecretArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:secretsmanager:getSecret", args, LookupSecretResultOutput{}, options).(LookupSecretResultOutput), nil
		}).(LookupSecretResultOutput)
}

type LookupSecretOutputArgs struct {
	// The ARN of the secret.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSecretOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretArgs)(nil)).Elem()
}

type LookupSecretResultOutput struct{ *pulumi.OutputState }

func (LookupSecretResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretResult)(nil)).Elem()
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutput() LookupSecretResultOutput {
	return o
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutputWithContext(ctx context.Context) LookupSecretResultOutput {
	return o
}

// The description of the secret.
func (o LookupSecretResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The ARN of the secret.
func (o LookupSecretResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by “alias/“, for example “alias/aws/secretsmanager“. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).
//
//	To use a KMS key in a different account, use the key ARN or the alias ARN.
//	If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
//	If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.
func (o LookupSecretResultOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// A custom type that specifies a “Region“ and the “KmsKeyId“ for a replica secret.
func (o LookupSecretResultOutput) ReplicaRegions() SecretReplicaRegionArrayOutput {
	return o.ApplyT(func(v LookupSecretResult) []SecretReplicaRegion { return v.ReplicaRegions }).(SecretReplicaRegionArrayOutput)
}

// A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:
//
//	 ``[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]``
//	Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
//	Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret.
//	If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2).
//	For information about how to format a JSON parameter for the various command line tool environments, see [Using JSON for Parameters](https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json). If your command-line tool or SDK requires quotation marks around the parameter, you should use single quotes to avoid confusion with the double quotes required in the JSON text.
//	The following restrictions apply to tags:
//	 +  Maximum number of tags per secret: 50
//	 +  Maximum key length: 127 Unicode characters in UTF-8
//	 +  Maximum value length: 255 Unicode characters in UTF-8
//	 +  Tag keys and values are case sensitive.
//	 +  Do not use the ``aws:`` prefix in your tag names or values because AWS reserves it for AWS use. You can't edit or delete tag names or values with this prefix. Tags with this prefix do not count against your tags per secret limit.
//	 +  If you use your tagging schema across multiple services and resources, other services might have restrictions on allowed characters. Generally allowed characters: letters, spaces, and numbers representable in UTF-8, plus the following special characters: + - = . _ : / @.
func (o LookupSecretResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupSecretResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretResultOutput{})
}
