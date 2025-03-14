// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::HealthImaging::Datastore Resource Type
 */
export function getDatastore(args: GetDatastoreArgs, opts?: pulumi.InvokeOptions): Promise<GetDatastoreResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:healthimaging:getDatastore", {
        "datastoreId": args.datastoreId,
    }, opts);
}

export interface GetDatastoreArgs {
    /**
     * The data store identifier.
     */
    datastoreId: string;
}

export interface GetDatastoreResult {
    /**
     * The timestamp when the data store was created.
     */
    readonly createdAt?: string;
    /**
     * The Amazon Resource Name (ARN) for the data store.
     */
    readonly datastoreArn?: string;
    /**
     * The data store identifier.
     */
    readonly datastoreId?: string;
    /**
     * The data store status.
     */
    readonly datastoreStatus?: enums.healthimaging.DatastoreStatus;
    /**
     * The timestamp when the data store was last updated.
     */
    readonly updatedAt?: string;
}
/**
 * Definition of AWS::HealthImaging::Datastore Resource Type
 */
export function getDatastoreOutput(args: GetDatastoreOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDatastoreResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:healthimaging:getDatastore", {
        "datastoreId": args.datastoreId,
    }, opts);
}

export interface GetDatastoreOutputArgs {
    /**
     * The data store identifier.
     */
    datastoreId: pulumi.Input<string>;
}
