// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::StudioLifecycleConfig
 */
export function getStudioLifecycleConfig(args: GetStudioLifecycleConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetStudioLifecycleConfigResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sagemaker:getStudioLifecycleConfig", {
        "studioLifecycleConfigName": args.studioLifecycleConfigName,
    }, opts);
}

export interface GetStudioLifecycleConfigArgs {
    /**
     * The name of the Amazon SageMaker Studio Lifecycle Configuration.
     */
    studioLifecycleConfigName: string;
}

export interface GetStudioLifecycleConfigResult {
    /**
     * The Amazon Resource Name (ARN) of the Lifecycle Configuration.
     */
    readonly studioLifecycleConfigArn?: string;
}
/**
 * Resource Type definition for AWS::SageMaker::StudioLifecycleConfig
 */
export function getStudioLifecycleConfigOutput(args: GetStudioLifecycleConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStudioLifecycleConfigResult> {
    return pulumi.output(args).apply((a: any) => getStudioLifecycleConfig(a, opts))
}

export interface GetStudioLifecycleConfigOutputArgs {
    /**
     * The name of the Amazon SageMaker Studio Lifecycle Configuration.
     */
    studioLifecycleConfigName: pulumi.Input<string>;
}