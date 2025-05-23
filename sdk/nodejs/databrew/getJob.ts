// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataBrew::Job.
 */
export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:databrew:getJob", {
        "name": args.name,
    }, opts);
}

export interface GetJobArgs {
    /**
     * Job name
     */
    name: string;
}

export interface GetJobResult {
    /**
     * One or more artifacts that represent the AWS Glue Data Catalog output from running the job.
     */
    readonly dataCatalogOutputs?: outputs.databrew.JobDataCatalogOutput[];
    /**
     * Represents a list of JDBC database output objects which defines the output destination for a DataBrew recipe job to write into.
     */
    readonly databaseOutputs?: outputs.databrew.JobDatabaseOutput[];
    /**
     * Dataset name
     */
    readonly datasetName?: string;
    /**
     * Encryption Key Arn
     */
    readonly encryptionKeyArn?: string;
    /**
     * Encryption mode
     */
    readonly encryptionMode?: enums.databrew.JobEncryptionMode;
    /**
     * Job Sample
     */
    readonly jobSample?: outputs.databrew.JobSample;
    /**
     * Log subscription
     */
    readonly logSubscription?: enums.databrew.JobLogSubscription;
    /**
     * Max capacity
     */
    readonly maxCapacity?: number;
    /**
     * Max retries
     */
    readonly maxRetries?: number;
    /**
     * Output location
     */
    readonly outputLocation?: outputs.databrew.JobOutputLocation;
    /**
     * One or more artifacts that represent output from running the job.
     */
    readonly outputs?: outputs.databrew.JobOutput[];
    /**
     * Profile Job configuration
     */
    readonly profileConfiguration?: outputs.databrew.JobProfileConfiguration;
    /**
     * Project name
     */
    readonly projectName?: string;
    /**
     * A series of data transformation steps that the job runs.
     */
    readonly recipe?: outputs.databrew.JobRecipe;
    /**
     * Role arn
     */
    readonly roleArn?: string;
    /**
     * Metadata tags that have been applied to the job.
     */
    readonly tags?: outputs.Tag[];
    /**
     * Timeout
     */
    readonly timeout?: number;
    /**
     * Data quality rules configuration
     */
    readonly validationConfigurations?: outputs.databrew.JobValidationConfiguration[];
}
/**
 * Resource schema for AWS::DataBrew::Job.
 */
export function getJobOutput(args: GetJobOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetJobResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("aws-native:databrew:getJob", {
        "name": args.name,
    }, opts);
}

export interface GetJobOutputArgs {
    /**
     * Job name
     */
    name: pulumi.Input<string>;
}
