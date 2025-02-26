// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::DataZone::Environment Resource Type
 */
export function getEnvironment(args: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:datazone:getEnvironment", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetEnvironmentArgs {
    /**
     * The identifier of the Amazon DataZone domain in which the environment is created.
     */
    domainId: string;
    /**
     * The ID of the Amazon DataZone environment.
     */
    id: string;
}

export interface GetEnvironmentResult {
    /**
     * The AWS account in which the Amazon DataZone environment is created.
     */
    readonly awsAccountId?: string;
    /**
     * The AWS region in which the Amazon DataZone environment is created.
     */
    readonly awsAccountRegion?: string;
    /**
     * The timestamp of when the environment was created.
     */
    readonly createdAt?: string;
    /**
     * The Amazon DataZone user who created the environment.
     */
    readonly createdBy?: string;
    /**
     * The description of the Amazon DataZone environment.
     */
    readonly description?: string;
    /**
     * The identifier of the Amazon DataZone domain in which the environment is created.
     */
    readonly domainId?: string;
    /**
     * The ID of the blueprint with which the Amazon DataZone environment was created.
     */
    readonly environmentBlueprintId?: string;
    /**
     * The ID of the environment profile with which the Amazon DataZone environment was created.
     */
    readonly environmentProfileId?: string;
    /**
     * The glossary terms that can be used in the Amazon DataZone environment.
     */
    readonly glossaryTerms?: string[];
    /**
     * The ID of the Amazon DataZone environment.
     */
    readonly id?: string;
    /**
     * The name of the environment.
     */
    readonly name?: string;
    /**
     * The ID of the Amazon DataZone project in which the environment is created.
     */
    readonly projectId?: string;
    /**
     * The provider of the Amazon DataZone environment.
     */
    readonly provider?: string;
    /**
     * The status of the Amazon DataZone environment.
     */
    readonly status?: enums.datazone.EnvironmentStatus;
    /**
     * The timestamp of when the environment was updated.
     */
    readonly updatedAt?: string;
}
/**
 * Definition of AWS::DataZone::Environment Resource Type
 */
export function getEnvironmentOutput(args: GetEnvironmentOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEnvironmentResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:datazone:getEnvironment", {
        "domainId": args.domainId,
        "id": args.id,
    }, opts);
}

export interface GetEnvironmentOutputArgs {
    /**
     * The identifier of the Amazon DataZone domain in which the environment is created.
     */
    domainId: pulumi.Input<string>;
    /**
     * The ID of the Amazon DataZone environment.
     */
    id: pulumi.Input<string>;
}
