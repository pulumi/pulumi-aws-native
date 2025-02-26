// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::CodePipeline::CustomActionType resource creates a custom action for activities that aren't included in the CodePipeline default actions, such as running an internally developed build process or a test suite. You can use these custom actions in the stage of a pipeline.
 */
export function getCustomActionType(args: GetCustomActionTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomActionTypeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:codepipeline:getCustomActionType", {
        "category": args.category,
        "provider": args.provider,
        "version": args.version,
    }, opts);
}

export interface GetCustomActionTypeArgs {
    /**
     * The category of the custom action, such as a build action or a test action.
     */
    category: string;
    /**
     * The provider of the service used in the custom action, such as AWS CodeDeploy.
     */
    provider: string;
    /**
     * The version identifier of the custom action.
     */
    version: string;
}

export interface GetCustomActionTypeResult {
    readonly id?: string;
    /**
     * Any tags assigned to the custom action.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * The AWS::CodePipeline::CustomActionType resource creates a custom action for activities that aren't included in the CodePipeline default actions, such as running an internally developed build process or a test suite. You can use these custom actions in the stage of a pipeline.
 */
export function getCustomActionTypeOutput(args: GetCustomActionTypeOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCustomActionTypeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:codepipeline:getCustomActionType", {
        "category": args.category,
        "provider": args.provider,
        "version": args.version,
    }, opts);
}

export interface GetCustomActionTypeOutputArgs {
    /**
     * The category of the custom action, such as a build action or a test action.
     */
    category: pulumi.Input<string>;
    /**
     * The provider of the service used in the custom action, such as AWS CodeDeploy.
     */
    provider: pulumi.Input<string>;
    /**
     * The version identifier of the custom action.
     */
    version: pulumi.Input<string>;
}
