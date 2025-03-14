// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Definition of AWS::MediaLive::CloudWatchAlarmTemplateGroup Resource Type
 */
export class CloudWatchAlarmTemplateGroup extends pulumi.CustomResource {
    /**
     * Get an existing CloudWatchAlarmTemplateGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CloudWatchAlarmTemplateGroup {
        return new CloudWatchAlarmTemplateGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:medialive:CloudWatchAlarmTemplateGroup';

    /**
     * Returns true if the given object is an instance of CloudWatchAlarmTemplateGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudWatchAlarmTemplateGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudWatchAlarmTemplateGroup.__pulumiType;
    }

    /**
     * A cloudwatch alarm template group's ARN (Amazon Resource Name)
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A cloudwatch alarm template group's id. AWS provided template groups have ids that start with `aws-`
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The date and time of resource creation.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A resource's optional description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly identifier!: pulumi.Output<string>;
    /**
     * The date and time of latest resource modification.
     */
    public /*out*/ readonly modifiedAt!: pulumi.Output<string>;
    /**
     * A resource's name. Names must be unique within the scope of a resource type in a specific region.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a CloudWatchAlarmTemplateGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CloudWatchAlarmTemplateGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["identifier"] = undefined /*out*/;
            resourceInputs["modifiedAt"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["identifier"] = undefined /*out*/;
            resourceInputs["modifiedAt"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name", "tags.*"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(CloudWatchAlarmTemplateGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CloudWatchAlarmTemplateGroup resource.
 */
export interface CloudWatchAlarmTemplateGroupArgs {
    /**
     * A resource's optional description.
     */
    description?: pulumi.Input<string>;
    /**
     * A resource's name. Names must be unique within the scope of a resource type in a specific region.
     */
    name?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
