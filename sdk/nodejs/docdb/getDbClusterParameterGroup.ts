// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::DocDB::DBClusterParameterGroup
 */
export function getDbClusterParameterGroup(args: GetDbClusterParameterGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetDbClusterParameterGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:docdb:getDbClusterParameterGroup", {
        "id": args.id,
    }, opts);
}

export interface GetDbClusterParameterGroupArgs {
    id: string;
}

export interface GetDbClusterParameterGroupResult {
    readonly id?: string;
    readonly parameters?: any;
    readonly tags?: outputs.docdb.DbClusterParameterGroupTag[];
}
/**
 * Resource Type definition for AWS::DocDB::DBClusterParameterGroup
 */
export function getDbClusterParameterGroupOutput(args: GetDbClusterParameterGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDbClusterParameterGroupResult> {
    return pulumi.output(args).apply((a: any) => getDbClusterParameterGroup(a, opts))
}

export interface GetDbClusterParameterGroupOutputArgs {
    id: pulumi.Input<string>;
}