// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IoT::PolicyPrincipalAttachment
 */
export function getPolicyPrincipalAttachment(args: GetPolicyPrincipalAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyPrincipalAttachmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iot:getPolicyPrincipalAttachment", {
        "id": args.id,
    }, opts);
}

export interface GetPolicyPrincipalAttachmentArgs {
    id: string;
}

export interface GetPolicyPrincipalAttachmentResult {
    readonly id?: string;
}

export function getPolicyPrincipalAttachmentOutput(args: GetPolicyPrincipalAttachmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyPrincipalAttachmentResult> {
    return pulumi.output(args).apply(a => getPolicyPrincipalAttachment(a, opts))
}

export interface GetPolicyPrincipalAttachmentOutputArgs {
    id: pulumi.Input<string>;
}