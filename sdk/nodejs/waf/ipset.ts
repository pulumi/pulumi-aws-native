// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAF::IPSet
 *
 * @deprecated IPSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class IPSet extends pulumi.CustomResource {
    /**
     * Get an existing IPSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IPSet {
        pulumi.log.warn("IPSet is deprecated: IPSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new IPSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:waf:IPSet';

    /**
     * Returns true if the given object is an instance of IPSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IPSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IPSet.__pulumiType;
    }

    public readonly iPSetDescriptors!: pulumi.Output<outputs.waf.IPSetDescriptor[] | undefined>;
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a IPSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated IPSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: IPSetArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("IPSet is deprecated: IPSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["iPSetDescriptors"] = args ? args.iPSetDescriptors : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        } else {
            resourceInputs["iPSetDescriptors"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IPSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a IPSet resource.
 */
export interface IPSetArgs {
    iPSetDescriptors?: pulumi.Input<pulumi.Input<inputs.waf.IPSetDescriptorArgs>[]>;
    name?: pulumi.Input<string>;
}