// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ServiceDiscovery::PrivateDnsNamespace
 *
 * @deprecated PrivateDnsNamespace is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class PrivateDnsNamespace extends pulumi.CustomResource {
    /**
     * Get an existing PrivateDnsNamespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateDnsNamespace {
        pulumi.log.warn("PrivateDnsNamespace is deprecated: PrivateDnsNamespace is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new PrivateDnsNamespace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:servicediscovery:PrivateDnsNamespace';

    /**
     * Returns true if the given object is an instance of PrivateDnsNamespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateDnsNamespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateDnsNamespace.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly hostedZoneId!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly properties!: pulumi.Output<outputs.servicediscovery.PrivateDnsNamespaceProperties | undefined>;
    public readonly tags!: pulumi.Output<outputs.servicediscovery.PrivateDnsNamespaceTag[] | undefined>;
    public readonly vpc!: pulumi.Output<string>;

    /**
     * Create a PrivateDnsNamespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated PrivateDnsNamespace is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: PrivateDnsNamespaceArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("PrivateDnsNamespace is deprecated: PrivateDnsNamespace is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.vpc === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpc'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpc"] = args ? args.vpc : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["hostedZoneId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["hostedZoneId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["vpc"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivateDnsNamespace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateDnsNamespace resource.
 */
export interface PrivateDnsNamespaceArgs {
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    properties?: pulumi.Input<inputs.servicediscovery.PrivateDnsNamespacePropertiesArgs>;
    tags?: pulumi.Input<pulumi.Input<inputs.servicediscovery.PrivateDnsNamespaceTagArgs>[]>;
    vpc: pulumi.Input<string>;
}