// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Wisdom::Assistant Resource Type
 */
export class Assistant extends pulumi.CustomResource {
    /**
     * Get an existing Assistant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Assistant {
        return new Assistant(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wisdom:Assistant';

    /**
     * Returns true if the given object is an instance of Assistant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Assistant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Assistant.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the assistant.
     */
    public /*out*/ readonly assistantArn!: pulumi.Output<string>;
    /**
     * The ID of the Wisdom assistant.
     */
    public /*out*/ readonly assistantId!: pulumi.Output<string>;
    /**
     * The description of the assistant.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the assistant.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The configuration information for the customer managed key used for encryption. The customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom. To use Wisdom with chat, the key policy must also allow `kms:Decrypt` , `kms:GenerateDataKey*` , and `kms:DescribeKey` permissions to the `connect.amazonaws.com` service principal. For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) .
     */
    public readonly serverSideEncryptionConfiguration!: pulumi.Output<outputs.wisdom.AssistantServerSideEncryptionConfiguration | undefined>;
    /**
     * The tags used to organize, track, or control access for this resource.
     */
    public readonly tags!: pulumi.Output<outputs.CreateOnlyTag[] | undefined>;
    /**
     * The type of assistant.
     */
    public readonly type!: pulumi.Output<enums.wisdom.AssistantType>;

    /**
     * Create a Assistant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssistantArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serverSideEncryptionConfiguration"] = args ? args.serverSideEncryptionConfiguration : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["assistantArn"] = undefined /*out*/;
            resourceInputs["assistantId"] = undefined /*out*/;
        } else {
            resourceInputs["assistantArn"] = undefined /*out*/;
            resourceInputs["assistantId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serverSideEncryptionConfiguration"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["description", "name", "serverSideEncryptionConfiguration", "tags[*]", "type"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Assistant.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Assistant resource.
 */
export interface AssistantArgs {
    /**
     * The description of the assistant.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the assistant.
     */
    name?: pulumi.Input<string>;
    /**
     * The configuration information for the customer managed key used for encryption. The customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom. To use Wisdom with chat, the key policy must also allow `kms:Decrypt` , `kms:GenerateDataKey*` , and `kms:DescribeKey` permissions to the `connect.amazonaws.com` service principal. For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) .
     */
    serverSideEncryptionConfiguration?: pulumi.Input<inputs.wisdom.AssistantServerSideEncryptionConfigurationArgs>;
    /**
     * The tags used to organize, track, or control access for this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.CreateOnlyTagArgs>[]>;
    /**
     * The type of assistant.
     */
    type: pulumi.Input<enums.wisdom.AssistantType>;
}
