// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Deadline::QueueEnvironment Resource Type
 */
export function getQueueEnvironment(args: GetQueueEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetQueueEnvironmentResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:deadline:getQueueEnvironment", {
        "farmId": args.farmId,
        "queueEnvironmentId": args.queueEnvironmentId,
        "queueId": args.queueId,
    }, opts);
}

export interface GetQueueEnvironmentArgs {
    farmId: string;
    queueEnvironmentId: string;
    queueId: string;
}

export interface GetQueueEnvironmentResult {
    readonly name?: string;
    readonly priority?: number;
    readonly queueEnvironmentId?: string;
    readonly template?: string;
    readonly templateType?: enums.deadline.QueueEnvironmentEnvironmentTemplateType;
}
/**
 * Definition of AWS::Deadline::QueueEnvironment Resource Type
 */
export function getQueueEnvironmentOutput(args: GetQueueEnvironmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQueueEnvironmentResult> {
    return pulumi.output(args).apply((a: any) => getQueueEnvironment(a, opts))
}

export interface GetQueueEnvironmentOutputArgs {
    farmId: pulumi.Input<string>;
    queueEnvironmentId: pulumi.Input<string>;
    queueId: pulumi.Input<string>;
}