// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::Macie::Session resource specifies a new Amazon Macie session. A session is an object that represents the Amazon Macie service. A session is required for Amazon Macie to become operational.
 */
export function getSession(args: GetSessionArgs, opts?: pulumi.InvokeOptions): Promise<GetSessionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:macie:getSession", {
        "awsAccountId": args.awsAccountId,
    }, opts);
}

export interface GetSessionArgs {
    /**
     * AWS account ID of customer
     */
    awsAccountId: string;
}

export interface GetSessionResult {
    /**
     * AWS account ID of customer
     */
    readonly awsAccountId?: string;
    /**
     * A enumeration value that specifies how frequently finding updates are published.
     */
    readonly findingPublishingFrequency?: enums.macie.SessionFindingPublishingFrequency;
    /**
     * Service role used by Macie
     */
    readonly serviceRole?: string;
    /**
     * A enumeration value that specifies the status of the Macie Session.
     */
    readonly status?: enums.macie.SessionStatus;
}

export function getSessionOutput(args: GetSessionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSessionResult> {
    return pulumi.output(args).apply(a => getSession(a, opts))
}

export interface GetSessionOutputArgs {
    /**
     * AWS account ID of customer
     */
    awsAccountId: pulumi.Input<string>;
}