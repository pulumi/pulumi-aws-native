// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource type definition for AWS::NetworkFirewall::TLSInspectionConfiguration
 */
export function getTlsInspectionConfiguration(args: GetTlsInspectionConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetTlsInspectionConfigurationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:networkfirewall:getTlsInspectionConfiguration", {
        "tlsInspectionConfigurationArn": args.tlsInspectionConfigurationArn,
    }, opts);
}

export interface GetTlsInspectionConfigurationArgs {
    tlsInspectionConfigurationArn: string;
}

export interface GetTlsInspectionConfigurationResult {
    readonly description?: string;
    readonly tags?: outputs.networkfirewall.TlsInspectionConfigurationTag[];
    readonly tlsInspectionConfiguration?: outputs.networkfirewall.TlsInspectionConfigurationTlsInspectionConfiguration;
    readonly tlsInspectionConfigurationArn?: string;
    readonly tlsInspectionConfigurationId?: string;
}
/**
 * Resource type definition for AWS::NetworkFirewall::TLSInspectionConfiguration
 */
export function getTlsInspectionConfigurationOutput(args: GetTlsInspectionConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTlsInspectionConfigurationResult> {
    return pulumi.output(args).apply((a: any) => getTlsInspectionConfiguration(a, opts))
}

export interface GetTlsInspectionConfigurationOutputArgs {
    tlsInspectionConfigurationArn: pulumi.Input<string>;
}