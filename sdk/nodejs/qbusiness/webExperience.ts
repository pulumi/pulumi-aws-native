// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::QBusiness::WebExperience Resource Type
 */
export class WebExperience extends pulumi.CustomResource {
    /**
     * Get an existing WebExperience resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebExperience {
        return new WebExperience(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:qbusiness:WebExperience';

    /**
     * Returns true if the given object is an instance of WebExperience.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebExperience {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebExperience.__pulumiType;
    }

    /**
     * The identifier of the Amazon Q Business web experience.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * The Unix timestamp when the Amazon Q Business application was last updated.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The endpoint URLs for your Amazon Q Business web experience. The URLs are unique and fully hosted by AWS .
     */
    public /*out*/ readonly defaultEndpoint!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the service role attached to your web experience.
     *
     * > You must provide this value if you're using IAM Identity Center to manage end user access to your application. If you're using legacy identity management to manage user access, you don't need to provide this value.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * Determines whether sample prompts are enabled in the web experience for an end user.
     */
    public readonly samplePromptsControlMode!: pulumi.Output<enums.qbusiness.WebExperienceSamplePromptsControlMode | undefined>;
    /**
     * The status of your Amazon Q Business web experience.
     */
    public /*out*/ readonly status!: pulumi.Output<enums.qbusiness.WebExperienceStatus>;
    /**
     * A subtitle to personalize your Amazon Q Business web experience.
     */
    public readonly subtitle!: pulumi.Output<string | undefined>;
    /**
     * A list of key-value pairs that identify or categorize your Amazon Q Business web experience. You can also use tags to help control access to the web experience. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The title for your Amazon Q Business web experience.
     */
    public readonly title!: pulumi.Output<string | undefined>;
    /**
     * The Unix timestamp when your Amazon Q Business web experience was updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of an Amazon Q Business web experience.
     */
    public /*out*/ readonly webExperienceArn!: pulumi.Output<string>;
    /**
     * The identifier of your Amazon Q Business web experience.
     */
    public /*out*/ readonly webExperienceId!: pulumi.Output<string>;
    /**
     * A message in an Amazon Q Business web experience.
     */
    public readonly welcomeMessage!: pulumi.Output<string | undefined>;

    /**
     * Create a WebExperience resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebExperienceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["samplePromptsControlMode"] = args ? args.samplePromptsControlMode : undefined;
            resourceInputs["subtitle"] = args ? args.subtitle : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["welcomeMessage"] = args ? args.welcomeMessage : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["defaultEndpoint"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["webExperienceArn"] = undefined /*out*/;
            resourceInputs["webExperienceId"] = undefined /*out*/;
        } else {
            resourceInputs["applicationId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["defaultEndpoint"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["samplePromptsControlMode"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subtitle"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["webExperienceArn"] = undefined /*out*/;
            resourceInputs["webExperienceId"] = undefined /*out*/;
            resourceInputs["welcomeMessage"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["applicationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(WebExperience.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebExperience resource.
 */
export interface WebExperienceArgs {
    /**
     * The identifier of the Amazon Q Business web experience.
     */
    applicationId: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the service role attached to your web experience.
     *
     * > You must provide this value if you're using IAM Identity Center to manage end user access to your application. If you're using legacy identity management to manage user access, you don't need to provide this value.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * Determines whether sample prompts are enabled in the web experience for an end user.
     */
    samplePromptsControlMode?: pulumi.Input<enums.qbusiness.WebExperienceSamplePromptsControlMode>;
    /**
     * A subtitle to personalize your Amazon Q Business web experience.
     */
    subtitle?: pulumi.Input<string>;
    /**
     * A list of key-value pairs that identify or categorize your Amazon Q Business web experience. You can also use tags to help control access to the web experience. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The title for your Amazon Q Business web experience.
     */
    title?: pulumi.Input<string>;
    /**
     * A message in an Amazon Q Business web experience.
     */
    welcomeMessage?: pulumi.Input<string>;
}