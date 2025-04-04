// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of the AWS::QuickSight::Topic Resource Type.
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:quicksight:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the topic.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ID of the AWS account that you want to create a topic in.
     */
    public readonly awsAccountId!: pulumi.Output<string | undefined>;
    /**
     * Configuration options for a `Topic` .
     */
    public readonly configOptions!: pulumi.Output<outputs.quicksight.TopicConfigOptions | undefined>;
    /**
     * The data sets that the topic is associated with.
     */
    public readonly dataSets!: pulumi.Output<outputs.quicksight.TopicDatasetMetadata[] | undefined>;
    /**
     * The description of the topic.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly folderArns!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the topic.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The ID for the topic. This ID is unique per AWS Region for each AWS account.
     */
    public readonly topicId!: pulumi.Output<string | undefined>;
    /**
     * The user experience version of the topic.
     */
    public readonly userExperienceVersion!: pulumi.Output<enums.quicksight.TopicUserExperienceVersion | undefined>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TopicArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["awsAccountId"] = args ? args.awsAccountId : undefined;
            resourceInputs["configOptions"] = args ? args.configOptions : undefined;
            resourceInputs["dataSets"] = args ? args.dataSets : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderArns"] = args ? args.folderArns : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["topicId"] = args ? args.topicId : undefined;
            resourceInputs["userExperienceVersion"] = args ? args.userExperienceVersion : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsAccountId"] = undefined /*out*/;
            resourceInputs["configOptions"] = undefined /*out*/;
            resourceInputs["dataSets"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["folderArns"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["topicId"] = undefined /*out*/;
            resourceInputs["userExperienceVersion"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["awsAccountId", "folderArns[*]", "topicId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Topic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * The ID of the AWS account that you want to create a topic in.
     */
    awsAccountId?: pulumi.Input<string>;
    /**
     * Configuration options for a `Topic` .
     */
    configOptions?: pulumi.Input<inputs.quicksight.TopicConfigOptionsArgs>;
    /**
     * The data sets that the topic is associated with.
     */
    dataSets?: pulumi.Input<pulumi.Input<inputs.quicksight.TopicDatasetMetadataArgs>[]>;
    /**
     * The description of the topic.
     */
    description?: pulumi.Input<string>;
    folderArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the topic.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID for the topic. This ID is unique per AWS Region for each AWS account.
     */
    topicId?: pulumi.Input<string>;
    /**
     * The user experience version of the topic.
     */
    userExperienceVersion?: pulumi.Input<enums.quicksight.TopicUserExperienceVersion>;
}
