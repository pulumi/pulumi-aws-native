// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * This resource represents a schema of Glue Schema Registry.
 */
export function getSchema(args: GetSchemaArgs, opts?: pulumi.InvokeOptions): Promise<GetSchemaResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:glue:getSchema", {
        "arn": args.arn,
    }, opts);
}

export interface GetSchemaArgs {
    /**
     * Amazon Resource Name for the Schema.
     */
    arn: string;
}

export interface GetSchemaResult {
    /**
     * Amazon Resource Name for the Schema.
     */
    readonly arn?: string;
    /**
     * Specify the `VersionNumber` or the `IsLatest` for setting the checkpoint for the schema. This is only required for updating a checkpoint.
     */
    readonly checkpointVersion?: outputs.glue.SchemaVersion;
    /**
     * Compatibility setting for the schema.
     */
    readonly compatibility?: enums.glue.SchemaCompatibility;
    /**
     * A description of the schema. If description is not provided, there will not be any default value for this.
     */
    readonly description?: string;
    /**
     * Represents the version ID associated with the initial schema version.
     */
    readonly initialSchemaVersionId?: string;
    /**
     * List of tags to tag the schema
     */
    readonly tags?: outputs.Tag[];
}
/**
 * This resource represents a schema of Glue Schema Registry.
 */
export function getSchemaOutput(args: GetSchemaOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSchemaResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:glue:getSchema", {
        "arn": args.arn,
    }, opts);
}

export interface GetSchemaOutputArgs {
    /**
     * Amazon Resource Name for the Schema.
     */
    arn: pulumi.Input<string>;
}
