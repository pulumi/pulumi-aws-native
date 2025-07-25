// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The ``AWS::ApiGateway::BasePathMapping`` resource creates a base path that clients who call your API must use in the invocation URL. Supported only for public custom domain names.
 */
export function getBasePathMapping(args: GetBasePathMappingArgs, opts?: pulumi.InvokeOptions): Promise<GetBasePathMappingResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:apigateway:getBasePathMapping", {
        "basePath": args.basePath,
        "domainName": args.domainName,
    }, opts);
}

export interface GetBasePathMappingArgs {
    /**
     * The base path name that callers of the API must provide as part of the URL after the domain name.
     */
    basePath: string;
    /**
     * The domain name of the BasePathMapping resource to be described.
     */
    domainName: string;
}

export interface GetBasePathMappingResult {
    /**
     * The string identifier of the associated RestApi.
     */
    readonly restApiId?: string;
    /**
     * The name of the associated stage.
     */
    readonly stage?: string;
}
/**
 * The ``AWS::ApiGateway::BasePathMapping`` resource creates a base path that clients who call your API must use in the invocation URL. Supported only for public custom domain names.
 */
export function getBasePathMappingOutput(args: GetBasePathMappingOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBasePathMappingResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:apigateway:getBasePathMapping", {
        "basePath": args.basePath,
        "domainName": args.domainName,
    }, opts);
}

export interface GetBasePathMappingOutputArgs {
    /**
     * The base path name that callers of the API must provide as part of the URL after the domain name.
     */
    basePath: pulumi.Input<string>;
    /**
     * The domain name of the BasePathMapping resource to be described.
     */
    domainName: pulumi.Input<string>;
}
