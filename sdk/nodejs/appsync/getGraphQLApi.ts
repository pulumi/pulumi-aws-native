// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppSync::GraphQLApi
 */
export function getGraphQLApi(args: GetGraphQLApiArgs, opts?: pulumi.InvokeOptions): Promise<GetGraphQLApiResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appsync:getGraphQLApi", {
        "id": args.id,
    }, opts);
}

export interface GetGraphQLApiArgs {
    id: string;
}

export interface GetGraphQLApiResult {
    readonly additionalAuthenticationProviders?: outputs.appsync.GraphQLApiAdditionalAuthenticationProvider[];
    readonly apiId?: string;
    readonly arn?: string;
    readonly authenticationType?: string;
    readonly graphQLUrl?: string;
    readonly id?: string;
    readonly lambdaAuthorizerConfig?: outputs.appsync.GraphQLApiLambdaAuthorizerConfig;
    readonly logConfig?: outputs.appsync.GraphQLApiLogConfig;
    readonly name?: string;
    readonly openIDConnectConfig?: outputs.appsync.GraphQLApiOpenIDConnectConfig;
    readonly tags?: outputs.appsync.GraphQLApiTag[];
    readonly userPoolConfig?: outputs.appsync.GraphQLApiUserPoolConfig;
    readonly xrayEnabled?: boolean;
}

export function getGraphQLApiOutput(args: GetGraphQLApiOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGraphQLApiResult> {
    return pulumi.output(args).apply(a => getGraphQLApi(a, opts))
}

export interface GetGraphQLApiOutputArgs {
    id: pulumi.Input<string>;
}