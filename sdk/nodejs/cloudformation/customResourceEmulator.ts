// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The Custom Resource Emulator allows you to use AWS CloudFormation Custom Resources directly in your Pulumi programs. It provides a way to invoke AWS Lambda functions that implement custom provisioning logic following the CloudFormation Custom Resource protocol.
 *
 * > **Note**: Currently, only Lambda-backed Custom Resources are supported. SNS-backed Custom Resources are not supported at this time.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as aws from "@pulumi/aws-native";
 *
 * const bucket = new aws.s3.Bucket('custom-resource-emulator');
 *
 * // Create a Custom Resource that invokes a Lambda function
 * const cr = new aws.cloudformation.CustomResourceEmulator('cr', {
 *     bucketName: bucket.id,
 *     bucketKeyPrefix: 'custom-resource-emulator',
 *     customResourceProperties: {
 *         hello: "world"
 *     },
 *     serviceToken: "arn:aws:lambda:us-west-2:123456789012:function:my-custom-resource",
 *     resourceType: 'Custom::MyResource',
 * }, { customTimeouts: { create: '5m', update: '5m', delete: '5m' } });
 *
 * // Access the response data
 * export const customResourceData = customResource.data;
 * ```
 *
 * A full example of creating a CloudFormation Custom Resource Lambda function and using it in Pulumi can be found [here](https://github.com/pulumi/pulumi-aws-native/tree/master/examples/cfn-custom-resource).
 *
 * ## About CloudFormation Custom Resources
 *
 * CloudFormation Custom Resources allow you to write custom provisioning logic for resources that aren't directly available as AWS CloudFormation resource types. Common use cases include:
 *
 * - Implementing complex provisioning logic
 * - Performing custom validations or transformations
 * - Integrating with third-party services
 * - Implementing organization-specific infrastructure patterns
 *
 * For more information about CloudFormation Custom Resources, see [Custom Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html) in the AWS CloudFormation User Guide.
 *
 * ## Permissions
 *
 * The IAM principal used by your Pulumi program must have the following permissions:
 *
 * 1. `lambda:InvokeFunction` on the Lambda function specified in `serviceToken`
 * 2. S3 permissions on the bucket specified in `bucketName`:
 *    - `s3:PutObject`
 *    - `s3:GetObject`
 *    - `s3:HeadObject`
 *
 * ## Lambda Function Requirements
 *
 * The Lambda function specified in `serviceToken` must implement the CloudFormation Custom Resource lifecycle.
 * For detailed information about implementing Lambda-backed Custom Resources, see [AWS Lambda-backed Custom Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources-lambda.html) in the AWS CloudFormation User Guide.
 *
 * ## Timeouts
 *
 * Custom Resources have a default timeout of 60 minutes, matching the CloudFormation timeout for custom resource operations. You can customize it using the [`customTimeouts`](https://www.pulumi.com/docs/iac/concepts/options/customtimeouts/) resource option.
 */
export class CustomResourceEmulator extends pulumi.CustomResource {
    /**
     * Get an existing CustomResourceEmulator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomResourceEmulator {
        return new CustomResourceEmulator(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudformation:CustomResourceEmulator';

    /**
     * Returns true if the given object is an instance of CustomResourceEmulator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomResourceEmulator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomResourceEmulator.__pulumiType;
    }

    /**
     * The name of the S3 bucket to use for storing the response from the Custom Resource.
     */
    public /*out*/ readonly bucket!: pulumi.Output<string>;
    /**
     * The response data returned by invoking the Custom Resource.
     */
    public /*out*/ readonly data!: pulumi.Output<{[key: string]: any}>;
    /**
     * Whether the response data contains sensitive information that should be marked as secret and not logged.
     */
    public /*out*/ readonly noEcho!: pulumi.Output<boolean>;
    /**
     * The name or unique identifier that corresponds to the `PhysicalResourceId` included in the Custom Resource response. If no `PhysicalResourceId` is provided in the response, a random ID will be generated.
     */
    public /*out*/ readonly physicalResourceId!: pulumi.Output<string>;
    /**
     * The CloudFormation type of the Custom Resource provider. For example, `Custom::MyCustomResource`.
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * The service token, such as a Lambda function ARN, that is invoked when the resource is created, updated, or deleted.
     */
    public readonly serviceToken!: pulumi.Output<string>;
    /**
     * A stand-in value for the CloudFormation stack ID.
     */
    public readonly stackId!: pulumi.Output<string>;

    /**
     * Create a CustomResourceEmulator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomResourceEmulatorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bucketKeyPrefix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucketKeyPrefix'");
            }
            if ((!args || args.bucketName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucketName'");
            }
            if ((!args || args.customResourceProperties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customResourceProperties'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            if ((!args || args.serviceToken === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceToken'");
            }
            resourceInputs["bucketKeyPrefix"] = args ? args.bucketKeyPrefix : undefined;
            resourceInputs["bucketName"] = args ? args.bucketName : undefined;
            resourceInputs["customResourceProperties"] = args ? args.customResourceProperties : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["serviceToken"] = args ? args.serviceToken : undefined;
            resourceInputs["stackId"] = args ? args.stackId : undefined;
            resourceInputs["bucket"] = undefined /*out*/;
            resourceInputs["data"] = undefined /*out*/;
            resourceInputs["noEcho"] = undefined /*out*/;
            resourceInputs["physicalResourceId"] = undefined /*out*/;
        } else {
            resourceInputs["bucket"] = undefined /*out*/;
            resourceInputs["data"] = undefined /*out*/;
            resourceInputs["noEcho"] = undefined /*out*/;
            resourceInputs["physicalResourceId"] = undefined /*out*/;
            resourceInputs["resourceType"] = undefined /*out*/;
            resourceInputs["serviceToken"] = undefined /*out*/;
            resourceInputs["stackId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomResourceEmulator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomResourceEmulator resource.
 */
export interface CustomResourceEmulatorArgs {
    /**
     * The prefix to use for the bucket key when storing the response from the Custom Resource provider.
     */
    bucketKeyPrefix: pulumi.Input<string>;
    /**
     * The name of the S3 bucket to use for storing the response from the Custom Resource.
     *
     * The IAM principal configured for the provider must have `s3:PutObject`, `s3:HeadObject` and `s3:GetObject` permissions on this bucket.
     */
    bucketName: pulumi.Input<string>;
    /**
     * The properties to pass as an input to the Custom Resource.
     * The properties are passed as a map of key-value pairs whereas all primitive values (number, boolean) are converted to strings for CloudFormation interoperability.
     */
    customResourceProperties: pulumi.Input<{[key: string]: any}>;
    /**
     * The CloudFormation type of the Custom Resource. For example, `Custom::MyCustomResource`.
     * This is required for CloudFormation interoperability.
     */
    resourceType: pulumi.Input<string>;
    /**
     * The service token to use for the Custom Resource. The service token is invoked when the resource is created, updated, or deleted.
     * This can be a Lambda Function ARN with optional version or alias identifiers.
     *
     * The IAM principal configured for the provider must have `lambda:InvokeFunction` permissions on this service token.
     */
    serviceToken: pulumi.Input<string>;
    /**
     * A stand-in value for the CloudFormation stack ID. This is required for CloudFormation interoperability.
     * If not provided, the Pulumi Stack ID is used.
     */
    stackId?: pulumi.Input<string>;
}
