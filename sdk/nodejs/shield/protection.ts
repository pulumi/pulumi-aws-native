// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Enables AWS Shield Advanced for a specific AWS resource. The resource can be an Amazon CloudFront distribution, Amazon Route 53 hosted zone, AWS Global Accelerator standard accelerator, Elastic IP Address, Application Load Balancer, or a Classic Load Balancer. You can protect Amazon EC2 instances and Network Load Balancers by association with protected Amazon EC2 Elastic IP addresses.
 *
 * @deprecated Protection is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Protection extends pulumi.CustomResource {
    /**
     * Get an existing Protection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Protection {
        pulumi.log.warn("Protection is deprecated: Protection is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Protection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:shield:Protection';

    /**
     * Returns true if the given object is an instance of Protection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Protection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Protection.__pulumiType;
    }

    public readonly applicationLayerAutomaticResponseConfiguration!: pulumi.Output<outputs.shield.ProtectionApplicationLayerAutomaticResponseConfiguration | undefined>;
    /**
     * The Amazon Resource Names (ARNs) of the health check to associate with the protection.
     */
    public readonly healthCheckArns!: pulumi.Output<string[] | undefined>;
    /**
     * Friendly name for the Protection.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ARN (Amazon Resource Name) of the protection.
     */
    public /*out*/ readonly protectionArn!: pulumi.Output<string>;
    /**
     * The unique identifier (ID) of the protection.
     */
    public /*out*/ readonly protectionId!: pulumi.Output<string>;
    /**
     * The ARN (Amazon Resource Name) of the resource to be protected.
     */
    public readonly resourceArn!: pulumi.Output<string>;
    /**
     * One or more tag key-value pairs for the Protection object.
     */
    public readonly tags!: pulumi.Output<outputs.shield.ProtectionTag[] | undefined>;

    /**
     * Create a Protection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Protection is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ProtectionArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Protection is deprecated: Protection is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceArn'");
            }
            resourceInputs["applicationLayerAutomaticResponseConfiguration"] = args ? args.applicationLayerAutomaticResponseConfiguration : undefined;
            resourceInputs["healthCheckArns"] = args ? args.healthCheckArns : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceArn"] = args ? args.resourceArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["protectionArn"] = undefined /*out*/;
            resourceInputs["protectionId"] = undefined /*out*/;
        } else {
            resourceInputs["applicationLayerAutomaticResponseConfiguration"] = undefined /*out*/;
            resourceInputs["healthCheckArns"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["protectionArn"] = undefined /*out*/;
            resourceInputs["protectionId"] = undefined /*out*/;
            resourceInputs["resourceArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Protection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Protection resource.
 */
export interface ProtectionArgs {
    applicationLayerAutomaticResponseConfiguration?: pulumi.Input<inputs.shield.ProtectionApplicationLayerAutomaticResponseConfigurationArgs>;
    /**
     * The Amazon Resource Names (ARNs) of the health check to associate with the protection.
     */
    healthCheckArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Friendly name for the Protection.
     */
    name?: pulumi.Input<string>;
    /**
     * The ARN (Amazon Resource Name) of the resource to be protected.
     */
    resourceArn: pulumi.Input<string>;
    /**
     * One or more tag key-value pairs for the Protection object.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.shield.ProtectionTagArgs>[]>;
}