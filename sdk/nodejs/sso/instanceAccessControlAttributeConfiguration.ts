// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for SSO InstanceAccessControlAttributeConfiguration
 */
export class InstanceAccessControlAttributeConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing InstanceAccessControlAttributeConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InstanceAccessControlAttributeConfiguration {
        return new InstanceAccessControlAttributeConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sso:InstanceAccessControlAttributeConfiguration';

    /**
     * Returns true if the given object is an instance of InstanceAccessControlAttributeConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceAccessControlAttributeConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceAccessControlAttributeConfiguration.__pulumiType;
    }

    public readonly accessControlAttributes!: pulumi.Output<outputs.sso.InstanceAccessControlAttributeConfigurationAccessControlAttribute[] | undefined>;
    /**
     * The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes. We recomend that you use  AccessControlAttributes property instead.
     */
    public readonly instanceAccessControlAttributeConfiguration!: pulumi.Output<outputs.sso.InstanceAccessControlAttributeConfigurationProperties | undefined>;
    /**
     * The ARN of the AWS SSO instance under which the operation will be executed.
     */
    public readonly instanceArn!: pulumi.Output<string>;

    /**
     * Create a InstanceAccessControlAttributeConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceAccessControlAttributeConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceArn'");
            }
            resourceInputs["accessControlAttributes"] = args ? args.accessControlAttributes : undefined;
            resourceInputs["instanceAccessControlAttributeConfiguration"] = args ? args.instanceAccessControlAttributeConfiguration : undefined;
            resourceInputs["instanceArn"] = args ? args.instanceArn : undefined;
        } else {
            resourceInputs["accessControlAttributes"] = undefined /*out*/;
            resourceInputs["instanceAccessControlAttributeConfiguration"] = undefined /*out*/;
            resourceInputs["instanceArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceAccessControlAttributeConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a InstanceAccessControlAttributeConfiguration resource.
 */
export interface InstanceAccessControlAttributeConfigurationArgs {
    accessControlAttributes?: pulumi.Input<pulumi.Input<inputs.sso.InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs>[]>;
    /**
     * The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes. We recomend that you use  AccessControlAttributes property instead.
     */
    instanceAccessControlAttributeConfiguration?: pulumi.Input<inputs.sso.InstanceAccessControlAttributeConfigurationPropertiesArgs>;
    /**
     * The ARN of the AWS SSO instance under which the operation will be executed.
     */
    instanceArn: pulumi.Input<string>;
}