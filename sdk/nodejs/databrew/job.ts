// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataBrew::Job.
 *
 * ## Example Usage
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const myDataBrewProfileJob = new aws_native.databrew.Job("myDataBrewProfileJob", {
 *     type: aws_native.databrew.JobType.Profile,
 *     name: "job-test",
 *     datasetName: "dataset-test",
 *     roleArn: "arn:aws:iam::1234567891011:role/PassRoleAdmin",
 *     jobSample: {
 *         mode: aws_native.databrew.JobSampleMode.FullDataset,
 *     },
 *     outputLocation: {
 *         bucket: "test-output",
 *         key: "job-output.json",
 *     },
 *     tags: [{
 *         key: "key00AtCreate",
 *         value: "value001AtCreate",
 *     }],
 * });
 *
 * ```
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:databrew:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * One or more artifacts that represent the AWS Glue Data Catalog output from running the job.
     */
    public readonly dataCatalogOutputs!: pulumi.Output<outputs.databrew.JobDataCatalogOutput[] | undefined>;
    /**
     * Represents a list of JDBC database output objects which defines the output destination for a DataBrew recipe job to write into.
     */
    public readonly databaseOutputs!: pulumi.Output<outputs.databrew.JobDatabaseOutput[] | undefined>;
    /**
     * Dataset name
     */
    public readonly datasetName!: pulumi.Output<string | undefined>;
    /**
     * Encryption Key Arn
     */
    public readonly encryptionKeyArn!: pulumi.Output<string | undefined>;
    /**
     * Encryption mode
     */
    public readonly encryptionMode!: pulumi.Output<enums.databrew.JobEncryptionMode | undefined>;
    /**
     * Job Sample
     */
    public readonly jobSample!: pulumi.Output<outputs.databrew.JobSample | undefined>;
    /**
     * Log subscription
     */
    public readonly logSubscription!: pulumi.Output<enums.databrew.JobLogSubscription | undefined>;
    /**
     * Max capacity
     */
    public readonly maxCapacity!: pulumi.Output<number | undefined>;
    /**
     * Max retries
     */
    public readonly maxRetries!: pulumi.Output<number | undefined>;
    /**
     * Job name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Output location
     */
    public readonly outputLocation!: pulumi.Output<outputs.databrew.JobOutputLocation | undefined>;
    /**
     * One or more artifacts that represent output from running the job.
     */
    public readonly outputs!: pulumi.Output<outputs.databrew.JobOutput[] | undefined>;
    /**
     * Profile Job configuration
     */
    public readonly profileConfiguration!: pulumi.Output<outputs.databrew.JobProfileConfiguration | undefined>;
    /**
     * Project name
     */
    public readonly projectName!: pulumi.Output<string | undefined>;
    /**
     * A series of data transformation steps that the job runs.
     */
    public readonly recipe!: pulumi.Output<outputs.databrew.JobRecipe | undefined>;
    /**
     * Role arn
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * Metadata tags that have been applied to the job.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * Timeout
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * Job type
     */
    public readonly type!: pulumi.Output<enums.databrew.JobType>;
    /**
     * Data quality rules configuration
     */
    public readonly validationConfigurations!: pulumi.Output<outputs.databrew.JobValidationConfiguration[] | undefined>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["dataCatalogOutputs"] = args ? args.dataCatalogOutputs : undefined;
            resourceInputs["databaseOutputs"] = args ? args.databaseOutputs : undefined;
            resourceInputs["datasetName"] = args ? args.datasetName : undefined;
            resourceInputs["encryptionKeyArn"] = args ? args.encryptionKeyArn : undefined;
            resourceInputs["encryptionMode"] = args ? args.encryptionMode : undefined;
            resourceInputs["jobSample"] = args ? args.jobSample : undefined;
            resourceInputs["logSubscription"] = args ? args.logSubscription : undefined;
            resourceInputs["maxCapacity"] = args ? args.maxCapacity : undefined;
            resourceInputs["maxRetries"] = args ? args.maxRetries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outputLocation"] = args ? args.outputLocation : undefined;
            resourceInputs["outputs"] = args ? args.outputs : undefined;
            resourceInputs["profileConfiguration"] = args ? args.profileConfiguration : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["recipe"] = args ? args.recipe : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["validationConfigurations"] = args ? args.validationConfigurations : undefined;
        } else {
            resourceInputs["dataCatalogOutputs"] = undefined /*out*/;
            resourceInputs["databaseOutputs"] = undefined /*out*/;
            resourceInputs["datasetName"] = undefined /*out*/;
            resourceInputs["encryptionKeyArn"] = undefined /*out*/;
            resourceInputs["encryptionMode"] = undefined /*out*/;
            resourceInputs["jobSample"] = undefined /*out*/;
            resourceInputs["logSubscription"] = undefined /*out*/;
            resourceInputs["maxCapacity"] = undefined /*out*/;
            resourceInputs["maxRetries"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["outputLocation"] = undefined /*out*/;
            resourceInputs["outputs"] = undefined /*out*/;
            resourceInputs["profileConfiguration"] = undefined /*out*/;
            resourceInputs["projectName"] = undefined /*out*/;
            resourceInputs["recipe"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["timeout"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["validationConfigurations"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name", "type"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Job.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * One or more artifacts that represent the AWS Glue Data Catalog output from running the job.
     */
    dataCatalogOutputs?: pulumi.Input<pulumi.Input<inputs.databrew.JobDataCatalogOutputArgs>[]>;
    /**
     * Represents a list of JDBC database output objects which defines the output destination for a DataBrew recipe job to write into.
     */
    databaseOutputs?: pulumi.Input<pulumi.Input<inputs.databrew.JobDatabaseOutputArgs>[]>;
    /**
     * Dataset name
     */
    datasetName?: pulumi.Input<string>;
    /**
     * Encryption Key Arn
     */
    encryptionKeyArn?: pulumi.Input<string>;
    /**
     * Encryption mode
     */
    encryptionMode?: pulumi.Input<enums.databrew.JobEncryptionMode>;
    /**
     * Job Sample
     */
    jobSample?: pulumi.Input<inputs.databrew.JobSampleArgs>;
    /**
     * Log subscription
     */
    logSubscription?: pulumi.Input<enums.databrew.JobLogSubscription>;
    /**
     * Max capacity
     */
    maxCapacity?: pulumi.Input<number>;
    /**
     * Max retries
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * Job name
     */
    name?: pulumi.Input<string>;
    /**
     * Output location
     */
    outputLocation?: pulumi.Input<inputs.databrew.JobOutputLocationArgs>;
    /**
     * One or more artifacts that represent output from running the job.
     */
    outputs?: pulumi.Input<pulumi.Input<inputs.databrew.JobOutputArgs>[]>;
    /**
     * Profile Job configuration
     */
    profileConfiguration?: pulumi.Input<inputs.databrew.JobProfileConfigurationArgs>;
    /**
     * Project name
     */
    projectName?: pulumi.Input<string>;
    /**
     * A series of data transformation steps that the job runs.
     */
    recipe?: pulumi.Input<inputs.databrew.JobRecipeArgs>;
    /**
     * Role arn
     */
    roleArn: pulumi.Input<string>;
    /**
     * Metadata tags that have been applied to the job.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * Timeout
     */
    timeout?: pulumi.Input<number>;
    /**
     * Job type
     */
    type: pulumi.Input<enums.databrew.JobType>;
    /**
     * Data quality rules configuration
     */
    validationConfigurations?: pulumi.Input<pulumi.Input<inputs.databrew.JobValidationConfigurationArgs>[]>;
}
