// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Macie FindingsFilter resource schema.
 */
export class FindingsFilter extends pulumi.CustomResource {
    /**
     * Get an existing FindingsFilter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FindingsFilter {
        return new FindingsFilter(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:macie:FindingsFilter';

    /**
     * Returns true if the given object is an instance of FindingsFilter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FindingsFilter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FindingsFilter.__pulumiType;
    }

    /**
     * Findings filter action.
     */
    public readonly action!: pulumi.Output<enums.macie.FindingsFilterFindingFilterAction | undefined>;
    /**
     * Findings filter ARN.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Findings filter description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Findings filter criteria.
     */
    public readonly findingCriteria!: pulumi.Output<outputs.macie.FindingsFilterFindingCriteria>;
    /**
     * Findings filters list.
     */
    public /*out*/ readonly findingsFilterListItems!: pulumi.Output<outputs.macie.FindingsFilterListItem[]>;
    /**
     * Findings filter name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Findings filter position.
     */
    public readonly position!: pulumi.Output<number | undefined>;

    /**
     * Create a FindingsFilter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FindingsFilterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.findingCriteria === undefined) && !opts.urn) {
                throw new Error("Missing required property 'findingCriteria'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["findingCriteria"] = args ? args.findingCriteria : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["position"] = args ? args.position : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["findingsFilterListItems"] = undefined /*out*/;
        } else {
            resourceInputs["action"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["findingCriteria"] = undefined /*out*/;
            resourceInputs["findingsFilterListItems"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["position"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FindingsFilter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FindingsFilter resource.
 */
export interface FindingsFilterArgs {
    /**
     * Findings filter action.
     */
    action?: pulumi.Input<enums.macie.FindingsFilterFindingFilterAction>;
    /**
     * Findings filter description
     */
    description?: pulumi.Input<string>;
    /**
     * Findings filter criteria.
     */
    findingCriteria: pulumi.Input<inputs.macie.FindingsFilterFindingCriteriaArgs>;
    /**
     * Findings filter name
     */
    name?: pulumi.Input<string>;
    /**
     * Findings filter position.
     */
    position?: pulumi.Input<number>;
}