// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::Url
 */
export function getUrl(args: GetUrlArgs, opts?: pulumi.InvokeOptions): Promise<GetUrlResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:lambda:getUrl", {
        "functionArn": args.functionArn,
    }, opts);
}

export interface GetUrlArgs {
    /**
     * The full Amazon Resource Name (ARN) of the function associated with the Function URL.
     */
    functionArn: string;
}

export interface GetUrlResult {
    /**
     * Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
     */
    readonly authType?: enums.lambda.UrlAuthType;
    /**
     * The [Cross-Origin Resource Sharing (CORS)](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for your function URL.
     */
    readonly cors?: outputs.lambda.UrlCors;
    /**
     * The full Amazon Resource Name (ARN) of the function associated with the Function URL.
     */
    readonly functionArn?: string;
    /**
     * The generated url for this resource.
     */
    readonly functionUrl?: string;
    /**
     * The invocation mode for the function's URL. Set to BUFFERED if you want to buffer responses before returning them to the client. Set to RESPONSE_STREAM if you want to stream responses, allowing faster time to first byte and larger response payload sizes. If not set, defaults to BUFFERED.
     */
    readonly invokeMode?: enums.lambda.UrlInvokeMode;
}
/**
 * Resource Type definition for AWS::Lambda::Url
 */
export function getUrlOutput(args: GetUrlOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUrlResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:lambda:getUrl", {
        "functionArn": args.functionArn,
    }, opts);
}

export interface GetUrlOutputArgs {
    /**
     * The full Amazon Resource Name (ARN) of the function associated with the Function URL.
     */
    functionArn: pulumi.Input<string>;
}
