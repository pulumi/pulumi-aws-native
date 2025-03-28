// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Set a version as default version for a hook in CloudFormation Registry.
 */
export function getHookDefaultVersion(args: GetHookDefaultVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetHookDefaultVersionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cloudformation:getHookDefaultVersion", {
        "arn": args.arn,
    }, opts);
}

export interface GetHookDefaultVersionArgs {
    /**
     * The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion
     */
    arn: string;
}

export interface GetHookDefaultVersionResult {
    /**
     * The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion
     */
    readonly arn?: string;
    /**
     * The name of the type being registered.
     *
     * We recommend that type names adhere to the following pattern: company_or_organization::service::type.
     */
    readonly typeName?: string;
    /**
     * The Amazon Resource Name (ARN) of the type version.
     */
    readonly typeVersionArn?: string;
    /**
     * The ID of an existing version of the hook to set as the default.
     */
    readonly versionId?: string;
}
/**
 * Set a version as default version for a hook in CloudFormation Registry.
 */
export function getHookDefaultVersionOutput(args: GetHookDefaultVersionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetHookDefaultVersionResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:cloudformation:getHookDefaultVersion", {
        "arn": args.arn,
    }, opts);
}

export interface GetHookDefaultVersionOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion
     */
    arn: pulumi.Input<string>;
}
