// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Bedrock::DataAutomationProject Resource Type
 */
export class DataAutomationProject extends pulumi.CustomResource {
    /**
     * Get an existing DataAutomationProject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataAutomationProject {
        return new DataAutomationProject(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:bedrock:DataAutomationProject';

    /**
     * Returns true if the given object is an instance of DataAutomationProject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataAutomationProject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataAutomationProject.__pulumiType;
    }

    /**
     * Time Stamp
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Blueprints to apply to objects processed by the project.
     */
    public readonly customOutputConfiguration!: pulumi.Output<outputs.bedrock.DataAutomationProjectCustomOutputConfiguration | undefined>;
    /**
     * KMS encryption context
     */
    public readonly kmsEncryptionContext!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * KMS key identifier
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Time Stamp
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * Additional settings for the project.
     */
    public readonly overrideConfiguration!: pulumi.Output<outputs.bedrock.DataAutomationProjectOverrideConfiguration | undefined>;
    /**
     * ARN of a DataAutomationProject
     */
    public /*out*/ readonly projectArn!: pulumi.Output<string>;
    /**
     * Description of the DataAutomationProject
     */
    public readonly projectDescription!: pulumi.Output<string | undefined>;
    /**
     * Name of the DataAutomationProject
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * The project's stage.
     */
    public /*out*/ readonly projectStage!: pulumi.Output<enums.bedrock.DataAutomationProjectStage>;
    /**
     * The project's standard output configuration.
     */
    public readonly standardOutputConfiguration!: pulumi.Output<outputs.bedrock.DataAutomationProjectStandardOutputConfiguration | undefined>;
    /**
     * The project's status.
     */
    public /*out*/ readonly status!: pulumi.Output<enums.bedrock.DataAutomationProjectStatus>;
    /**
     * List of Tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a DataAutomationProject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DataAutomationProjectArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["customOutputConfiguration"] = args ? args.customOutputConfiguration : undefined;
            resourceInputs["kmsEncryptionContext"] = args ? args.kmsEncryptionContext : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["overrideConfiguration"] = args ? args.overrideConfiguration : undefined;
            resourceInputs["projectDescription"] = args ? args.projectDescription : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["standardOutputConfiguration"] = args ? args.standardOutputConfiguration : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["projectArn"] = undefined /*out*/;
            resourceInputs["projectStage"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["customOutputConfiguration"] = undefined /*out*/;
            resourceInputs["kmsEncryptionContext"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["overrideConfiguration"] = undefined /*out*/;
            resourceInputs["projectArn"] = undefined /*out*/;
            resourceInputs["projectDescription"] = undefined /*out*/;
            resourceInputs["projectName"] = undefined /*out*/;
            resourceInputs["projectStage"] = undefined /*out*/;
            resourceInputs["standardOutputConfiguration"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["projectName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DataAutomationProject.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataAutomationProject resource.
 */
export interface DataAutomationProjectArgs {
    /**
     * Blueprints to apply to objects processed by the project.
     */
    customOutputConfiguration?: pulumi.Input<inputs.bedrock.DataAutomationProjectCustomOutputConfigurationArgs>;
    /**
     * KMS encryption context
     */
    kmsEncryptionContext?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * KMS key identifier
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Additional settings for the project.
     */
    overrideConfiguration?: pulumi.Input<inputs.bedrock.DataAutomationProjectOverrideConfigurationArgs>;
    /**
     * Description of the DataAutomationProject
     */
    projectDescription?: pulumi.Input<string>;
    /**
     * Name of the DataAutomationProject
     */
    projectName?: pulumi.Input<string>;
    /**
     * The project's standard output configuration.
     */
    standardOutputConfiguration?: pulumi.Input<inputs.bedrock.DataAutomationProjectStandardOutputConfigurationArgs>;
    /**
     * List of Tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
