// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::ARCZonalShift::ZonalAutoshiftConfiguration Resource Type
 */
export function getZonalAutoshiftConfiguration(args: GetZonalAutoshiftConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetZonalAutoshiftConfigurationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:arczonalshift:getZonalAutoshiftConfiguration", {
        "resourceIdentifier": args.resourceIdentifier,
    }, opts);
}

export interface GetZonalAutoshiftConfigurationArgs {
    resourceIdentifier: string;
}

export interface GetZonalAutoshiftConfigurationResult {
    readonly practiceRunConfiguration?: outputs.arczonalshift.ZonalAutoshiftConfigurationPracticeRunConfiguration;
    readonly zonalAutoshiftStatus?: enums.arczonalshift.ZonalAutoshiftConfigurationZonalAutoshiftStatus;
}
/**
 * Definition of AWS::ARCZonalShift::ZonalAutoshiftConfiguration Resource Type
 */
export function getZonalAutoshiftConfigurationOutput(args: GetZonalAutoshiftConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZonalAutoshiftConfigurationResult> {
    return pulumi.output(args).apply((a: any) => getZonalAutoshiftConfiguration(a, opts))
}

export interface GetZonalAutoshiftConfigurationOutputArgs {
    resourceIdentifier: pulumi.Input<string>;
}