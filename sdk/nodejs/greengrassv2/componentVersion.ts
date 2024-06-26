// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource for Greengrass component version.
 */
export class ComponentVersion extends pulumi.CustomResource {
    /**
     * Get an existing ComponentVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ComponentVersion {
        return new ComponentVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:greengrassv2:ComponentVersion';

    /**
     * Returns true if the given object is an instance of ComponentVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComponentVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComponentVersion.__pulumiType;
    }

    /**
     * The ARN of the component version.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the component.
     */
    public /*out*/ readonly componentName!: pulumi.Output<string>;
    /**
     * The version of the component.
     */
    public /*out*/ readonly componentVersion!: pulumi.Output<string>;
    /**
     * The recipe to use to create the component. The recipe defines the component's metadata, parameters, dependencies, lifecycle, artifacts, and platform compatibility.
     *
     * You must specify either `InlineRecipe` or `LambdaFunction` .
     */
    public readonly inlineRecipe!: pulumi.Output<string | undefined>;
    /**
     * The parameters to create a component from a Lambda function.
     *
     * You must specify either `InlineRecipe` or `LambdaFunction` .
     */
    public readonly lambdaFunction!: pulumi.Output<outputs.greengrassv2.ComponentVersionLambdaFunctionRecipeSource | undefined>;
    /**
     * Application-specific metadata to attach to the component version. You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .
     *
     * This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
     *
     * ```json
     * "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
     * }
     * ```
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a ComponentVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ComponentVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["inlineRecipe"] = args ? args.inlineRecipe : undefined;
            resourceInputs["lambdaFunction"] = args ? args.lambdaFunction : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["componentName"] = undefined /*out*/;
            resourceInputs["componentVersion"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["componentName"] = undefined /*out*/;
            resourceInputs["componentVersion"] = undefined /*out*/;
            resourceInputs["inlineRecipe"] = undefined /*out*/;
            resourceInputs["lambdaFunction"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["inlineRecipe", "lambdaFunction"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ComponentVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ComponentVersion resource.
 */
export interface ComponentVersionArgs {
    /**
     * The recipe to use to create the component. The recipe defines the component's metadata, parameters, dependencies, lifecycle, artifacts, and platform compatibility.
     *
     * You must specify either `InlineRecipe` or `LambdaFunction` .
     */
    inlineRecipe?: pulumi.Input<string>;
    /**
     * The parameters to create a component from a Lambda function.
     *
     * You must specify either `InlineRecipe` or `LambdaFunction` .
     */
    lambdaFunction?: pulumi.Input<inputs.greengrassv2.ComponentVersionLambdaFunctionRecipeSourceArgs>;
    /**
     * Application-specific metadata to attach to the component version. You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .
     *
     * This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
     *
     * ```json
     * "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
     * }
     * ```
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
