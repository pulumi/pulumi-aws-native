// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SES::VdmAttributes
 */
export class VdmAttributes extends pulumi.CustomResource {
    /**
     * Get an existing VdmAttributes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VdmAttributes {
        return new VdmAttributes(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ses:VdmAttributes';

    /**
     * Returns true if the given object is an instance of VdmAttributes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VdmAttributes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VdmAttributes.__pulumiType;
    }

    public readonly dashboardAttributes!: pulumi.Output<outputs.ses.VdmAttributesDashboardAttributes | undefined>;
    public readonly guardianAttributes!: pulumi.Output<outputs.ses.VdmAttributesGuardianAttributes | undefined>;
    /**
     * Unique identifier for this resource
     */
    public /*out*/ readonly vdmAttributesResourceId!: pulumi.Output<string>;

    /**
     * Create a VdmAttributes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VdmAttributesArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["dashboardAttributes"] = args ? args.dashboardAttributes : undefined;
            resourceInputs["guardianAttributes"] = args ? args.guardianAttributes : undefined;
            resourceInputs["vdmAttributesResourceId"] = undefined /*out*/;
        } else {
            resourceInputs["dashboardAttributes"] = undefined /*out*/;
            resourceInputs["guardianAttributes"] = undefined /*out*/;
            resourceInputs["vdmAttributesResourceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VdmAttributes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a VdmAttributes resource.
 */
export interface VdmAttributesArgs {
    dashboardAttributes?: pulumi.Input<inputs.ses.VdmAttributesDashboardAttributesArgs>;
    guardianAttributes?: pulumi.Input<inputs.ses.VdmAttributesGuardianAttributesArgs>;
}