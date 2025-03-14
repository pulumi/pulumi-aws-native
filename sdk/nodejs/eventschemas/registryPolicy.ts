// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EventSchemas::RegistryPolicy
 *
 * ## Example Usage
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const registryPolicy = new aws_native.eventschemas.RegistryPolicy("registryPolicy", {
 *     registryName: "registryName",
 *     policy: {
 *         version: "2012-10-17",
 *         statement: {
 *             sid: 1,
 *             effect: "Allow",
 *             principal: {
 *                 aws: "arn:aws:iam::012345678901:user/TestAccountForRegistryPolicy",
 *             },
 *             action: [
 *                 "schemas:DescribeRegistry",
 *                 "schemas:CreateSchema",
 *             ],
 *             resource: "registryArn",
 *         },
 *     },
 * });
 *
 * ```
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const registryPolicy = new aws_native.eventschemas.RegistryPolicy("registryPolicy", {
 *     registryName: "MyRegistry",
 *     policy: {
 *         version: "2012-10-17",
 *         statement: [{
 *             sid: "Test",
 *             effect: "Allow",
 *             action: ["schemas:*"],
 *             principal: {
 *                 aws: ["109876543210"],
 *             },
 *             resource: [
 *                 "arn:aws:schemas:us-east-1:012345678901:registry/MyRegistry",
 *                 "arn:aws:schemas:us-east-1:012345678901:schema/MyRegistry*",
 *             ],
 *         }],
 *     },
 * });
 *
 * ```
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const registryPolicy = new aws_native.eventschemas.RegistryPolicy("registryPolicy", {
 *     registryName: "MyRegistry",
 *     policy: {
 *         version: "2012-10-17",
 *         statement: [{
 *             sid: "Test",
 *             effect: "Allow",
 *             action: ["schemas:*"],
 *             principal: {
 *                 aws: ["109876543210"],
 *             },
 *             resource: [
 *                 "arn:aws:schemas:us-east-1:012345678901:registry/MyRegistry",
 *                 "arn:aws:schemas:us-east-1:012345678901:schema/MyRegistry*",
 *             ],
 *         }],
 *     },
 * });
 *
 * ```
 */
export class RegistryPolicy extends pulumi.CustomResource {
    /**
     * Get an existing RegistryPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegistryPolicy {
        return new RegistryPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:eventschemas:RegistryPolicy';

    /**
     * Returns true if the given object is an instance of RegistryPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegistryPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegistryPolicy.__pulumiType;
    }

    /**
     * The ID of the policy.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * A resource-based policy.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EventSchemas::RegistryPolicy` for more information about the expected schema for this property.
     */
    public readonly policy!: pulumi.Output<any>;
    /**
     * The name of the registry.
     */
    public readonly registryName!: pulumi.Output<string>;
    /**
     * The revision ID of the policy.
     */
    public readonly revisionId!: pulumi.Output<string | undefined>;

    /**
     * Create a RegistryPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistryPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.registryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryName'");
            }
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["registryName"] = args ? args.registryName : undefined;
            resourceInputs["revisionId"] = args ? args.revisionId : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["policy"] = undefined /*out*/;
            resourceInputs["registryName"] = undefined /*out*/;
            resourceInputs["revisionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RegistryPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegistryPolicy resource.
 */
export interface RegistryPolicyArgs {
    /**
     * A resource-based policy.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EventSchemas::RegistryPolicy` for more information about the expected schema for this property.
     */
    policy: any;
    /**
     * The name of the registry.
     */
    registryName: pulumi.Input<string>;
    /**
     * The revision ID of the policy.
     */
    revisionId?: pulumi.Input<string>;
}
