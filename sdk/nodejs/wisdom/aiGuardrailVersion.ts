// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Wisdom::AIGuardrailVersion Resource Type
 */
export class AiGuardrailVersion extends pulumi.CustomResource {
    /**
     * Get an existing AiGuardrailVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AiGuardrailVersion {
        return new AiGuardrailVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wisdom:AiGuardrailVersion';

    /**
     * Returns true if the given object is an instance of AiGuardrailVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AiGuardrailVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AiGuardrailVersion.__pulumiType;
    }

    /**
     * The ARN of the AI guardrail version.
     */
    public /*out*/ readonly aiGuardrailArn!: pulumi.Output<string>;
    /**
     * The ID of the AI guardrail version.
     */
    public readonly aiGuardrailId!: pulumi.Output<string>;
    /**
     * The ID of the AI guardrail version.
     */
    public /*out*/ readonly aiGuardrailVersionId!: pulumi.Output<string>;
    /**
     * The ARN of the AI guardrail version assistant.
     */
    public /*out*/ readonly assistantArn!: pulumi.Output<string>;
    /**
     * The ID of the AI guardrail version assistant.
     */
    public readonly assistantId!: pulumi.Output<string>;
    /**
     * The modified time of the AI guardrail version in seconds.
     */
    public readonly modifiedTimeSeconds!: pulumi.Output<number | undefined>;
    /**
     * The version number for this AI Guardrail version.
     */
    public /*out*/ readonly versionNumber!: pulumi.Output<number>;

    /**
     * Create a AiGuardrailVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AiGuardrailVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.aiGuardrailId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'aiGuardrailId'");
            }
            if ((!args || args.assistantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'assistantId'");
            }
            resourceInputs["aiGuardrailId"] = args ? args.aiGuardrailId : undefined;
            resourceInputs["assistantId"] = args ? args.assistantId : undefined;
            resourceInputs["modifiedTimeSeconds"] = args ? args.modifiedTimeSeconds : undefined;
            resourceInputs["aiGuardrailArn"] = undefined /*out*/;
            resourceInputs["aiGuardrailVersionId"] = undefined /*out*/;
            resourceInputs["assistantArn"] = undefined /*out*/;
            resourceInputs["versionNumber"] = undefined /*out*/;
        } else {
            resourceInputs["aiGuardrailArn"] = undefined /*out*/;
            resourceInputs["aiGuardrailId"] = undefined /*out*/;
            resourceInputs["aiGuardrailVersionId"] = undefined /*out*/;
            resourceInputs["assistantArn"] = undefined /*out*/;
            resourceInputs["assistantId"] = undefined /*out*/;
            resourceInputs["modifiedTimeSeconds"] = undefined /*out*/;
            resourceInputs["versionNumber"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["aiGuardrailId", "assistantId", "modifiedTimeSeconds"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(AiGuardrailVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AiGuardrailVersion resource.
 */
export interface AiGuardrailVersionArgs {
    /**
     * The ID of the AI guardrail version.
     */
    aiGuardrailId: pulumi.Input<string>;
    /**
     * The ID of the AI guardrail version assistant.
     */
    assistantId: pulumi.Input<string>;
    /**
     * The modified time of the AI guardrail version in seconds.
     */
    modifiedTimeSeconds?: pulumi.Input<number>;
}
