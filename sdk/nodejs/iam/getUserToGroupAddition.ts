// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::UserToGroupAddition
 */
export function getUserToGroupAddition(args: GetUserToGroupAdditionArgs, opts?: pulumi.InvokeOptions): Promise<GetUserToGroupAdditionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iam:getUserToGroupAddition", {
        "id": args.id,
    }, opts);
}

export interface GetUserToGroupAdditionArgs {
    id: string;
}

export interface GetUserToGroupAdditionResult {
    readonly groupName?: string;
    readonly id?: string;
    readonly users?: string[];
}

export function getUserToGroupAdditionOutput(args: GetUserToGroupAdditionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserToGroupAdditionResult> {
    return pulumi.output(args).apply(a => getUserToGroupAddition(a, opts))
}

export interface GetUserToGroupAdditionOutputArgs {
    id: pulumi.Input<string>;
}