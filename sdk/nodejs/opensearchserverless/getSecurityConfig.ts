// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Amazon OpenSearchServerless security config resource
 */
export function getSecurityConfig(args: GetSecurityConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:opensearchserverless:getSecurityConfig", {
        "id": args.id,
    }, opts);
}

export interface GetSecurityConfigArgs {
    /**
     * The identifier of the security config
     */
    id: string;
}

export interface GetSecurityConfigResult {
    /**
     * Security config description
     */
    readonly description?: string;
    /**
     * The identifier of the security config
     */
    readonly id?: string;
    readonly samlOptions?: outputs.opensearchserverless.SecurityConfigSamlConfigOptions;
}

export function getSecurityConfigOutput(args: GetSecurityConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityConfigResult> {
    return pulumi.output(args).apply(a => getSecurityConfig(a, opts))
}

export interface GetSecurityConfigOutputArgs {
    /**
     * The identifier of the security config
     */
    id: pulumi.Input<string>;
}