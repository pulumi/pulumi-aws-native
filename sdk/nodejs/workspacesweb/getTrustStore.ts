// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::WorkSpacesWeb::TrustStore Resource Type
 */
export function getTrustStore(args: GetTrustStoreArgs, opts?: pulumi.InvokeOptions): Promise<GetTrustStoreResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:workspacesweb:getTrustStore", {
        "trustStoreArn": args.trustStoreArn,
    }, opts);
}

export interface GetTrustStoreArgs {
    /**
     * The ARN of the trust store.
     */
    trustStoreArn: string;
}

export interface GetTrustStoreResult {
    /**
     * A list of web portal ARNs that this trust store is associated with.
     */
    readonly associatedPortalArns?: string[];
    /**
     * A list of CA certificates to be added to the trust store.
     */
    readonly certificateList?: string[];
    /**
     * The tags to add to the trust store. A tag is a key-value pair.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The ARN of the trust store.
     */
    readonly trustStoreArn?: string;
}
/**
 * Definition of AWS::WorkSpacesWeb::TrustStore Resource Type
 */
export function getTrustStoreOutput(args: GetTrustStoreOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTrustStoreResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:workspacesweb:getTrustStore", {
        "trustStoreArn": args.trustStoreArn,
    }, opts);
}

export interface GetTrustStoreOutputArgs {
    /**
     * The ARN of the trust store.
     */
    trustStoreArn: pulumi.Input<string>;
}
