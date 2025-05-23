// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::QBusiness::Retriever Resource Type
 */
export function getRetriever(args: GetRetrieverArgs, opts?: pulumi.InvokeOptions): Promise<GetRetrieverResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:qbusiness:getRetriever", {
        "applicationId": args.applicationId,
        "retrieverId": args.retrieverId,
    }, opts);
}

export interface GetRetrieverArgs {
    /**
     * The identifier of the Amazon Q Business application using the retriever.
     */
    applicationId: string;
    /**
     * The identifier of the retriever used by your Amazon Q Business application.
     */
    retrieverId: string;
}

export interface GetRetrieverResult {
    /**
     * Provides information on how the retriever used for your Amazon Q Business application is configured.
     */
    readonly configuration?: outputs.qbusiness.RetrieverConfiguration0Properties | outputs.qbusiness.RetrieverConfiguration1Properties;
    /**
     * The Unix timestamp when the retriever was created.
     */
    readonly createdAt?: string;
    /**
     * The name of your retriever.
     */
    readonly displayName?: string;
    /**
     * The Amazon Resource Name (ARN) of the IAM role associated with the retriever.
     */
    readonly retrieverArn?: string;
    /**
     * The identifier of the retriever used by your Amazon Q Business application.
     */
    readonly retrieverId?: string;
    /**
     * The ARN of an IAM role used by Amazon Q Business to access the basic authentication credentials stored in a Secrets Manager secret.
     */
    readonly roleArn?: string;
    /**
     * The status of your retriever.
     */
    readonly status?: enums.qbusiness.RetrieverStatus;
    /**
     * A list of key-value pairs that identify or categorize the retriever. You can also use tags to help control access to the retriever. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @.
     */
    readonly tags?: outputs.Tag[];
    /**
     * The Unix timestamp when the retriever was last updated.
     */
    readonly updatedAt?: string;
}
/**
 * Definition of AWS::QBusiness::Retriever Resource Type
 */
export function getRetrieverOutput(args: GetRetrieverOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRetrieverResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:qbusiness:getRetriever", {
        "applicationId": args.applicationId,
        "retrieverId": args.retrieverId,
    }, opts);
}

export interface GetRetrieverOutputArgs {
    /**
     * The identifier of the Amazon Q Business application using the retriever.
     */
    applicationId: pulumi.Input<string>;
    /**
     * The identifier of the retriever used by your Amazon Q Business application.
     */
    retrieverId: pulumi.Input<string>;
}
