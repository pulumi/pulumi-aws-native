// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Deadline::Queue Resource Type
 */
export class Queue extends pulumi.CustomResource {
    /**
     * Get an existing Queue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Queue {
        return new Queue(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:deadline:Queue';

    /**
     * Returns true if the given object is an instance of Queue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Queue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Queue.__pulumiType;
    }

    /**
     * The identifiers of the storage profiles that this queue can use to share assets between workers using different operating systems.
     */
    public readonly allowedStorageProfileIds!: pulumi.Output<string[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the queue.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The default action taken on a queue summary if a budget wasn't configured.
     */
    public readonly defaultBudgetAction!: pulumi.Output<enums.deadline.QueueDefaultQueueBudgetAction | undefined>;
    /**
     * A description of the queue that helps identify what the queue is used for.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name of the queue summary to update.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The farm ID.
     */
    public readonly farmId!: pulumi.Output<string | undefined>;
    /**
     * The job attachment settings. These are the Amazon S3 bucket name and the Amazon S3 prefix.
     */
    public readonly jobAttachmentSettings!: pulumi.Output<outputs.deadline.QueueJobAttachmentSettings | undefined>;
    /**
     * Identifies the user for a job.
     */
    public readonly jobRunAsUser!: pulumi.Output<outputs.deadline.QueueJobRunAsUser | undefined>;
    /**
     * The queue ID.
     */
    public /*out*/ readonly queueId!: pulumi.Output<string>;
    /**
     * The file system location that the queue uses.
     */
    public readonly requiredFileSystemLocationNames!: pulumi.Output<string[] | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that workers use when running jobs in this queue.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Queue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueueArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["allowedStorageProfileIds"] = args ? args.allowedStorageProfileIds : undefined;
            resourceInputs["defaultBudgetAction"] = args ? args.defaultBudgetAction : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["farmId"] = args ? args.farmId : undefined;
            resourceInputs["jobAttachmentSettings"] = args ? args.jobAttachmentSettings : undefined;
            resourceInputs["jobRunAsUser"] = args ? args.jobRunAsUser : undefined;
            resourceInputs["requiredFileSystemLocationNames"] = args ? args.requiredFileSystemLocationNames : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["queueId"] = undefined /*out*/;
        } else {
            resourceInputs["allowedStorageProfileIds"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["defaultBudgetAction"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["farmId"] = undefined /*out*/;
            resourceInputs["jobAttachmentSettings"] = undefined /*out*/;
            resourceInputs["jobRunAsUser"] = undefined /*out*/;
            resourceInputs["queueId"] = undefined /*out*/;
            resourceInputs["requiredFileSystemLocationNames"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["farmId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Queue.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Queue resource.
 */
export interface QueueArgs {
    /**
     * The identifiers of the storage profiles that this queue can use to share assets between workers using different operating systems.
     */
    allowedStorageProfileIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The default action taken on a queue summary if a budget wasn't configured.
     */
    defaultBudgetAction?: pulumi.Input<enums.deadline.QueueDefaultQueueBudgetAction>;
    /**
     * A description of the queue that helps identify what the queue is used for.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the queue summary to update.
     */
    displayName: pulumi.Input<string>;
    /**
     * The farm ID.
     */
    farmId?: pulumi.Input<string>;
    /**
     * The job attachment settings. These are the Amazon S3 bucket name and the Amazon S3 prefix.
     */
    jobAttachmentSettings?: pulumi.Input<inputs.deadline.QueueJobAttachmentSettingsArgs>;
    /**
     * Identifies the user for a job.
     */
    jobRunAsUser?: pulumi.Input<inputs.deadline.QueueJobRunAsUserArgs>;
    /**
     * The file system location that the queue uses.
     */
    requiredFileSystemLocationNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that workers use when running jobs in this queue.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}