// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SSMContacts::Contact
 */
export function getContact(args: GetContactArgs, opts?: pulumi.InvokeOptions): Promise<GetContactResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ssmcontacts:getContact", {
        "arn": args.arn,
    }, opts);
}

export interface GetContactArgs {
    /**
     * The Amazon Resource Name (ARN) of the contact.
     */
    arn: string;
}

export interface GetContactResult {
    /**
     * The Amazon Resource Name (ARN) of the contact.
     */
    readonly arn?: string;
    /**
     * Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.
     */
    readonly displayName?: string;
}

export function getContactOutput(args: GetContactOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetContactResult> {
    return pulumi.output(args).apply(a => getContact(a, opts))
}

export interface GetContactOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the contact.
     */
    arn: pulumi.Input<string>;
}