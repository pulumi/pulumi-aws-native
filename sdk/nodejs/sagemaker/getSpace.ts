// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::Space
 */
export function getSpace(args: GetSpaceArgs, opts?: pulumi.InvokeOptions): Promise<GetSpaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sagemaker:getSpace", {
        "domainId": args.domainId,
        "spaceName": args.spaceName,
    }, opts);
}

export interface GetSpaceArgs {
    /**
     * The ID of the associated Domain.
     */
    domainId: string;
    /**
     * A name for the Space.
     */
    spaceName: string;
}

export interface GetSpaceResult {
    /**
     * The space Amazon Resource Name (ARN).
     */
    readonly spaceArn?: string;
}
/**
 * Resource Type definition for AWS::SageMaker::Space
 */
export function getSpaceOutput(args: GetSpaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSpaceResult> {
    return pulumi.output(args).apply((a: any) => getSpace(a, opts))
}

export interface GetSpaceOutputArgs {
    /**
     * The ID of the associated Domain.
     */
    domainId: pulumi.Input<string>;
    /**
     * A name for the Space.
     */
    spaceName: pulumi.Input<string>;
}