// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Description
 */
export function getResourcePolicy(args: GetResourcePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcePolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:vpclattice:getResourcePolicy", {
        "resourceArn": args.resourceArn,
    }, opts);
}

export interface GetResourcePolicyArgs {
    resourceArn: string;
}

export interface GetResourcePolicyResult {
    readonly policy?: any;
}
/**
 * Description
 */
export function getResourcePolicyOutput(args: GetResourcePolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourcePolicyResult> {
    return pulumi.output(args).apply((a: any) => getResourcePolicy(a, opts))
}

export interface GetResourcePolicyOutputArgs {
    resourceArn: pulumi.Input<string>;
}