// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::SecurityHub::ConfigurationPolicy resource represents the Central Configuration Policy in your account.
 */
export class ConfigurationPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ConfigurationPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConfigurationPolicy {
        return new ConfigurationPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:securityhub:ConfigurationPolicy';

    /**
     * Returns true if the given object is an instance of ConfigurationPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigurationPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigurationPolicy.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the configuration policy.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The universally unique identifier (UUID) of the configuration policy.
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * An object that defines how AWS Security Hub is configured. It includes whether Security Hub is enabled or disabled, a list of enabled security standards, a list of enabled or disabled security controls, and a list of custom parameter values for specified controls. If you provide a list of security controls that are enabled in the configuration policy, Security Hub disables all other controls (including newly released controls). If you provide a list of security controls that are disabled in the configuration policy, Security Hub enables all other controls (including newly released controls).
     */
    public readonly configurationPolicy!: pulumi.Output<outputs.securityhub.ConfigurationPolicyPolicy>;
    /**
     * The date and time, in UTC and ISO 8601 format.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The description of the configuration policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the configuration policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Indicates whether the service that the configuration policy applies to is enabled in the policy.
     */
    public /*out*/ readonly serviceEnabled!: pulumi.Output<boolean>;
    /**
     * User-defined tags associated with a configuration policy. For more information, see [Tagging AWS Security Hub resources](https://docs.aws.amazon.com/securityhub/latest/userguide/tagging-resources.html) in the *Security Hub user guide* .
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The date and time, in UTC and ISO 8601 format.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ConfigurationPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigurationPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.configurationPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configurationPolicy'");
            }
            resourceInputs["configurationPolicy"] = args ? args.configurationPolicy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["serviceEnabled"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["configurationPolicy"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serviceEnabled"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigurationPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConfigurationPolicy resource.
 */
export interface ConfigurationPolicyArgs {
    /**
     * An object that defines how AWS Security Hub is configured. It includes whether Security Hub is enabled or disabled, a list of enabled security standards, a list of enabled or disabled security controls, and a list of custom parameter values for specified controls. If you provide a list of security controls that are enabled in the configuration policy, Security Hub disables all other controls (including newly released controls). If you provide a list of security controls that are disabled in the configuration policy, Security Hub enables all other controls (including newly released controls).
     */
    configurationPolicy: pulumi.Input<inputs.securityhub.ConfigurationPolicyPolicyArgs>;
    /**
     * The description of the configuration policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the configuration policy.
     */
    name?: pulumi.Input<string>;
    /**
     * User-defined tags associated with a configuration policy. For more information, see [Tagging AWS Security Hub resources](https://docs.aws.amazon.com/securityhub/latest/userguide/tagging-resources.html) in the *Security Hub user guide* .
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
