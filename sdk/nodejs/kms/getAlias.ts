// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The ``AWS::KMS::Alias`` resource specifies a display name for a [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys). You can use an alias to identify a KMS key in the KMS console, in the [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html) operation, and in [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations), such as [Decrypt](https://docs.aws.amazon.com/kms/latest/APIReference/API_Decrypt.html) and [GenerateDataKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKey.html).
 *   Adding, deleting, or updating an alias can allow or deny permission to the KMS key. For details, see [ABAC for](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *Developer Guide*.
 *   Using an alias to refer to a KMS key can help you simplify key management. For example, an alias in your code can be associated with different KMS keys in different AWS-Regions. For more information, see [Using aliases](https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html) in the *Developer Guide*.
 *  When specifying an alias, observe the following rules.
 *   +  Each alias is associated with one KMS key, but multiple aliases can be associated with the same KMS key.
 *   +  The alias and its associated KMS key must be in the same AWS-account and Region.
 *   +  The alias name must be unique in the AWS-account and Region. However, you can create aliases with the same name in different AWS-Regions. For example, you can have an ``alias/projectKey`` in multiple Regions, each of which is associated with a KMS key in its Region.
 *   +  Each alias name must begin with ``alias/`` followed by a name, such as ``alias/exampleKey``. The alias name can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). Alias names cannot begin with ``alias/aws/``. That alias name prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
 *
 *   *Regions*
 *  KMS CloudFormation resources are available in all AWS-Regions in which KMS and CFN are supported.
 */
export function getAlias(args: GetAliasArgs, opts?: pulumi.InvokeOptions): Promise<GetAliasResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:kms:getAlias", {
        "aliasName": args.aliasName,
    }, opts);
}

export interface GetAliasArgs {
    /**
     * Specifies the alias name. This value must begin with ``alias/`` followed by a name, such as ``alias/ExampleAlias``. 
     *   If you change the value of the ``AliasName`` property, the existing alias is deleted and a new alias is created for the specified KMS key. This change can disrupt applications that use the alias. It can also allow or deny access to a KMS key affected by attribute-based access control (ABAC).
     *   The alias must be string of 1-256 characters. It can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). The alias name cannot begin with ``alias/aws/``. The ``alias/aws/`` prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
     */
    aliasName: string;
}

export interface GetAliasResult {
    /**
     * Associates the alias with the specified [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk). The KMS key must be in the same AWS-account and Region.
     *  A valid key ID is required. If you supply a null or empty string value, this operation returns an error.
     *  For help finding the key ID and ARN, see [Finding the key ID and ARN](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html#find-cmk-id-arn) in the *Developer Guide*.
     *  Specify the key ID or the key ARN of the KMS key.
     *  For example:
     *   +  Key ID: ``1234abcd-12ab-34cd-56ef-1234567890ab``
     *   +  Key ARN: ``arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab``
     *   
     *  To get the key ID and key ARN for a KMS key, use [ListKeys](https://docs.aws.amazon.com/kms/latest/APIReference/API_ListKeys.html) or [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html).
     */
    readonly targetKeyId?: string;
}
/**
 * The ``AWS::KMS::Alias`` resource specifies a display name for a [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys). You can use an alias to identify a KMS key in the KMS console, in the [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html) operation, and in [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations), such as [Decrypt](https://docs.aws.amazon.com/kms/latest/APIReference/API_Decrypt.html) and [GenerateDataKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKey.html).
 *   Adding, deleting, or updating an alias can allow or deny permission to the KMS key. For details, see [ABAC for](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *Developer Guide*.
 *   Using an alias to refer to a KMS key can help you simplify key management. For example, an alias in your code can be associated with different KMS keys in different AWS-Regions. For more information, see [Using aliases](https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html) in the *Developer Guide*.
 *  When specifying an alias, observe the following rules.
 *   +  Each alias is associated with one KMS key, but multiple aliases can be associated with the same KMS key.
 *   +  The alias and its associated KMS key must be in the same AWS-account and Region.
 *   +  The alias name must be unique in the AWS-account and Region. However, you can create aliases with the same name in different AWS-Regions. For example, you can have an ``alias/projectKey`` in multiple Regions, each of which is associated with a KMS key in its Region.
 *   +  Each alias name must begin with ``alias/`` followed by a name, such as ``alias/exampleKey``. The alias name can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). Alias names cannot begin with ``alias/aws/``. That alias name prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
 *
 *   *Regions*
 *  KMS CloudFormation resources are available in all AWS-Regions in which KMS and CFN are supported.
 */
export function getAliasOutput(args: GetAliasOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAliasResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:kms:getAlias", {
        "aliasName": args.aliasName,
    }, opts);
}

export interface GetAliasOutputArgs {
    /**
     * Specifies the alias name. This value must begin with ``alias/`` followed by a name, such as ``alias/ExampleAlias``. 
     *   If you change the value of the ``AliasName`` property, the existing alias is deleted and a new alias is created for the specified KMS key. This change can disrupt applications that use the alias. It can also allow or deny access to a KMS key affected by attribute-based access control (ABAC).
     *   The alias must be string of 1-256 characters. It can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). The alias name cannot begin with ``alias/aws/``. The ``alias/aws/`` prefix is reserved for [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
     */
    aliasName: pulumi.Input<string>;
}
