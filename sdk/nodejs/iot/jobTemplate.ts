// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IoT::JobTemplate. Job templates enable you to preconfigure jobs so that you can deploy them to multiple sets of target devices.
 */
export class JobTemplate extends pulumi.CustomResource {
    /**
     * Get an existing JobTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): JobTemplate {
        return new JobTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iot:JobTemplate';

    /**
     * Returns true if the given object is an instance of JobTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobTemplate.__pulumiType;
    }

    /**
     * The criteria that determine when and how a job abort takes place.
     */
    public readonly abortConfig!: pulumi.Output<outputs.iot.AbortConfigProperties | undefined>;
    /**
     * The ARN of the job to use as the basis for the job template.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A description of the Job Template.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The package version Amazon Resource Names (ARNs) that are installed on the device’s reserved named shadow ( `$package` ) when the job successfully completes.
     *
     * *Note:* Up to 25 package version ARNS are allowed.
     */
    public readonly destinationPackageVersions!: pulumi.Output<string[] | undefined>;
    /**
     * The job document. Required if you don't specify a value for documentSource.
     */
    public readonly document!: pulumi.Output<string | undefined>;
    /**
     * An S3 link to the job document to use in the template. Required if you don't specify a value for document.
     */
    public readonly documentSource!: pulumi.Output<string | undefined>;
    /**
     * Optional for copying a JobTemplate from a pre-existing Job configuration.
     */
    public readonly jobArn!: pulumi.Output<string | undefined>;
    /**
     * Allows you to create the criteria to retry a job.
     */
    public readonly jobExecutionsRetryConfig!: pulumi.Output<outputs.iot.JobExecutionsRetryConfigProperties | undefined>;
    /**
     * Allows you to create a staged rollout of a job.
     */
    public readonly jobExecutionsRolloutConfig!: pulumi.Output<outputs.iot.JobExecutionsRolloutConfigProperties | undefined>;
    /**
     * A unique identifier for the job template. We recommend using a UUID. Alpha-numeric characters, "-", and "_" are valid for use here.
     */
    public readonly jobTemplateId!: pulumi.Output<string>;
    /**
     * An optional configuration within the SchedulingConfig to setup a recurring maintenance window with a predetermined start time and duration for the rollout of a job document to all devices in a target group for a job.
     */
    public readonly maintenanceWindows!: pulumi.Output<outputs.iot.JobTemplateMaintenanceWindow[] | undefined>;
    /**
     * Configuration for pre-signed S3 URLs.
     */
    public readonly presignedUrlConfig!: pulumi.Output<outputs.iot.PresignedUrlConfigProperties | undefined>;
    /**
     * Metadata that can be used to manage the JobTemplate.
     */
    public readonly tags!: pulumi.Output<outputs.CreateOnlyTag[] | undefined>;
    /**
     * Specifies the amount of time each device has to finish its execution of the job.
     */
    public readonly timeoutConfig!: pulumi.Output<outputs.iot.TimeoutConfigProperties | undefined>;

    /**
     * Create a JobTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.jobTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobTemplateId'");
            }
            resourceInputs["abortConfig"] = args ? args.abortConfig : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationPackageVersions"] = args ? args.destinationPackageVersions : undefined;
            resourceInputs["document"] = args ? args.document : undefined;
            resourceInputs["documentSource"] = args ? args.documentSource : undefined;
            resourceInputs["jobArn"] = args ? args.jobArn : undefined;
            resourceInputs["jobExecutionsRetryConfig"] = args ? args.jobExecutionsRetryConfig : undefined;
            resourceInputs["jobExecutionsRolloutConfig"] = args ? args.jobExecutionsRolloutConfig : undefined;
            resourceInputs["jobTemplateId"] = args ? args.jobTemplateId : undefined;
            resourceInputs["maintenanceWindows"] = args ? args.maintenanceWindows : undefined;
            resourceInputs["presignedUrlConfig"] = args ? args.presignedUrlConfig : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timeoutConfig"] = args ? args.timeoutConfig : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["abortConfig"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["destinationPackageVersions"] = undefined /*out*/;
            resourceInputs["document"] = undefined /*out*/;
            resourceInputs["documentSource"] = undefined /*out*/;
            resourceInputs["jobArn"] = undefined /*out*/;
            resourceInputs["jobExecutionsRetryConfig"] = undefined /*out*/;
            resourceInputs["jobExecutionsRolloutConfig"] = undefined /*out*/;
            resourceInputs["jobTemplateId"] = undefined /*out*/;
            resourceInputs["maintenanceWindows"] = undefined /*out*/;
            resourceInputs["presignedUrlConfig"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["timeoutConfig"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["abortConfig", "description", "destinationPackageVersions[*]", "document", "documentSource", "jobArn", "jobExecutionsRetryConfig", "jobExecutionsRolloutConfig", "jobTemplateId", "maintenanceWindows[*]", "presignedUrlConfig", "tags[*]", "timeoutConfig"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(JobTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a JobTemplate resource.
 */
export interface JobTemplateArgs {
    /**
     * The criteria that determine when and how a job abort takes place.
     */
    abortConfig?: pulumi.Input<inputs.iot.AbortConfigPropertiesArgs>;
    /**
     * A description of the Job Template.
     */
    description: pulumi.Input<string>;
    /**
     * The package version Amazon Resource Names (ARNs) that are installed on the device’s reserved named shadow ( `$package` ) when the job successfully completes.
     *
     * *Note:* Up to 25 package version ARNS are allowed.
     */
    destinationPackageVersions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The job document. Required if you don't specify a value for documentSource.
     */
    document?: pulumi.Input<string>;
    /**
     * An S3 link to the job document to use in the template. Required if you don't specify a value for document.
     */
    documentSource?: pulumi.Input<string>;
    /**
     * Optional for copying a JobTemplate from a pre-existing Job configuration.
     */
    jobArn?: pulumi.Input<string>;
    /**
     * Allows you to create the criteria to retry a job.
     */
    jobExecutionsRetryConfig?: pulumi.Input<inputs.iot.JobExecutionsRetryConfigPropertiesArgs>;
    /**
     * Allows you to create a staged rollout of a job.
     */
    jobExecutionsRolloutConfig?: pulumi.Input<inputs.iot.JobExecutionsRolloutConfigPropertiesArgs>;
    /**
     * A unique identifier for the job template. We recommend using a UUID. Alpha-numeric characters, "-", and "_" are valid for use here.
     */
    jobTemplateId: pulumi.Input<string>;
    /**
     * An optional configuration within the SchedulingConfig to setup a recurring maintenance window with a predetermined start time and duration for the rollout of a job document to all devices in a target group for a job.
     */
    maintenanceWindows?: pulumi.Input<pulumi.Input<inputs.iot.JobTemplateMaintenanceWindowArgs>[]>;
    /**
     * Configuration for pre-signed S3 URLs.
     */
    presignedUrlConfig?: pulumi.Input<inputs.iot.PresignedUrlConfigPropertiesArgs>;
    /**
     * Metadata that can be used to manage the JobTemplate.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.CreateOnlyTagArgs>[]>;
    /**
     * Specifies the amount of time each device has to finish its execution of the job.
     */
    timeoutConfig?: pulumi.Input<inputs.iot.TimeoutConfigPropertiesArgs>;
}
