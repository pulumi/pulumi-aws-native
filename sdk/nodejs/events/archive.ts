// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Events::Archive
 */
export class Archive extends pulumi.CustomResource {
    /**
     * Get an existing Archive resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Archive {
        return new Archive(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:events:Archive';

    /**
     * Returns true if the given object is an instance of Archive.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Archive {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Archive.__pulumiType;
    }

    public readonly archiveName!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly eventPattern!: pulumi.Output<any | undefined>;
    public readonly retentionDays!: pulumi.Output<number | undefined>;
    public readonly sourceArn!: pulumi.Output<string>;

    /**
     * Create a Archive resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ArchiveArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.sourceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceArn'");
            }
            resourceInputs["archiveName"] = args ? args.archiveName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["eventPattern"] = args ? args.eventPattern : undefined;
            resourceInputs["retentionDays"] = args ? args.retentionDays : undefined;
            resourceInputs["sourceArn"] = args ? args.sourceArn : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["archiveName"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["eventPattern"] = undefined /*out*/;
            resourceInputs["retentionDays"] = undefined /*out*/;
            resourceInputs["sourceArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Archive.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Archive resource.
 */
export interface ArchiveArgs {
    archiveName?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    eventPattern?: any;
    retentionDays?: pulumi.Input<number>;
    sourceArn: pulumi.Input<string>;
}