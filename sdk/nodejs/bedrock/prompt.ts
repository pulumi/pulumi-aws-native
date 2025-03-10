// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Bedrock::Prompt Resource Type
 */
export class Prompt extends pulumi.CustomResource {
    /**
     * Get an existing Prompt resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Prompt {
        return new Prompt(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:bedrock:Prompt';

    /**
     * Returns true if the given object is an instance of Prompt.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Prompt {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Prompt.__pulumiType;
    }

    /**
     * ARN of a prompt resource possibly with a version
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Identifier for a Prompt
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * Time Stamp.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A KMS key ARN
     */
    public readonly customerEncryptionKeyArn!: pulumi.Output<string | undefined>;
    /**
     * Name for a variant.
     */
    public readonly defaultVariant!: pulumi.Output<string | undefined>;
    /**
     * Name for a prompt resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name for a prompt resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
     *
     * - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
     * - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Time Stamp.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * List of prompt variants
     */
    public readonly variants!: pulumi.Output<outputs.bedrock.PromptVariant[] | undefined>;
    /**
     * Draft Version.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a Prompt resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PromptArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["customerEncryptionKeyArn"] = args ? args.customerEncryptionKeyArn : undefined;
            resourceInputs["defaultVariant"] = args ? args.defaultVariant : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["variants"] = args ? args.variants : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["customerEncryptionKeyArn"] = undefined /*out*/;
            resourceInputs["defaultVariant"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["variants"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Prompt.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Prompt resource.
 */
export interface PromptArgs {
    /**
     * A KMS key ARN
     */
    customerEncryptionKeyArn?: pulumi.Input<string>;
    /**
     * Name for a variant.
     */
    defaultVariant?: pulumi.Input<string>;
    /**
     * Name for a prompt resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Name for a prompt resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
     *
     * - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
     * - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of prompt variants
     */
    variants?: pulumi.Input<pulumi.Input<inputs.bedrock.PromptVariantArgs>[]>;
}
