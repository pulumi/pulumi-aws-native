// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::RAM::ResourceShare
 *
 * @deprecated ResourceShare is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class ResourceShare extends pulumi.CustomResource {
    /**
     * Get an existing ResourceShare resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ResourceShare {
        pulumi.log.warn("ResourceShare is deprecated: ResourceShare is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new ResourceShare(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ram:ResourceShare';

    /**
     * Returns true if the given object is an instance of ResourceShare.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceShare {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceShare.__pulumiType;
    }

    public readonly allowExternalPrincipals!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly permissionArns!: pulumi.Output<string[] | undefined>;
    public readonly principals!: pulumi.Output<string[] | undefined>;
    public readonly resourceArns!: pulumi.Output<string[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.ram.ResourceShareTag[] | undefined>;

    /**
     * Create a ResourceShare resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated ResourceShare is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: ResourceShareArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("ResourceShare is deprecated: ResourceShare is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["allowExternalPrincipals"] = args ? args.allowExternalPrincipals : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["permissionArns"] = args ? args.permissionArns : undefined;
            resourceInputs["principals"] = args ? args.principals : undefined;
            resourceInputs["resourceArns"] = args ? args.resourceArns : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["allowExternalPrincipals"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["permissionArns"] = undefined /*out*/;
            resourceInputs["principals"] = undefined /*out*/;
            resourceInputs["resourceArns"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ResourceShare.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ResourceShare resource.
 */
export interface ResourceShareArgs {
    allowExternalPrincipals?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    permissionArns?: pulumi.Input<pulumi.Input<string>[]>;
    principals?: pulumi.Input<pulumi.Input<string>[]>;
    resourceArns?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.ram.ResourceShareTagArgs>[]>;
}