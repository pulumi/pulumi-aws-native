// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::NotebookInstance
 *
 * @deprecated NotebookInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class NotebookInstance extends pulumi.CustomResource {
    /**
     * Get an existing NotebookInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NotebookInstance {
        pulumi.log.warn("NotebookInstance is deprecated: NotebookInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new NotebookInstance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:NotebookInstance';

    /**
     * Returns true if the given object is an instance of NotebookInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotebookInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotebookInstance.__pulumiType;
    }

    public readonly acceleratorTypes!: pulumi.Output<string[] | undefined>;
    public readonly additionalCodeRepositories!: pulumi.Output<string[] | undefined>;
    public readonly defaultCodeRepository!: pulumi.Output<string | undefined>;
    public readonly directInternetAccess!: pulumi.Output<string | undefined>;
    public readonly instanceMetadataServiceConfiguration!: pulumi.Output<outputs.sagemaker.NotebookInstanceInstanceMetadataServiceConfiguration | undefined>;
    public readonly instanceType!: pulumi.Output<string>;
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    public readonly lifecycleConfigName!: pulumi.Output<string | undefined>;
    public readonly notebookInstanceName!: pulumi.Output<string | undefined>;
    public readonly platformIdentifier!: pulumi.Output<string | undefined>;
    public readonly roleArn!: pulumi.Output<string>;
    public readonly rootAccess!: pulumi.Output<string | undefined>;
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    public readonly subnetId!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.sagemaker.NotebookInstanceTag[] | undefined>;
    public readonly volumeSizeInGB!: pulumi.Output<number | undefined>;

    /**
     * Create a NotebookInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated NotebookInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: NotebookInstanceArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("NotebookInstance is deprecated: NotebookInstance is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["acceleratorTypes"] = args ? args.acceleratorTypes : undefined;
            resourceInputs["additionalCodeRepositories"] = args ? args.additionalCodeRepositories : undefined;
            resourceInputs["defaultCodeRepository"] = args ? args.defaultCodeRepository : undefined;
            resourceInputs["directInternetAccess"] = args ? args.directInternetAccess : undefined;
            resourceInputs["instanceMetadataServiceConfiguration"] = args ? args.instanceMetadataServiceConfiguration : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["lifecycleConfigName"] = args ? args.lifecycleConfigName : undefined;
            resourceInputs["notebookInstanceName"] = args ? args.notebookInstanceName : undefined;
            resourceInputs["platformIdentifier"] = args ? args.platformIdentifier : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["rootAccess"] = args ? args.rootAccess : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["volumeSizeInGB"] = args ? args.volumeSizeInGB : undefined;
        } else {
            resourceInputs["acceleratorTypes"] = undefined /*out*/;
            resourceInputs["additionalCodeRepositories"] = undefined /*out*/;
            resourceInputs["defaultCodeRepository"] = undefined /*out*/;
            resourceInputs["directInternetAccess"] = undefined /*out*/;
            resourceInputs["instanceMetadataServiceConfiguration"] = undefined /*out*/;
            resourceInputs["instanceType"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["lifecycleConfigName"] = undefined /*out*/;
            resourceInputs["notebookInstanceName"] = undefined /*out*/;
            resourceInputs["platformIdentifier"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["rootAccess"] = undefined /*out*/;
            resourceInputs["securityGroupIds"] = undefined /*out*/;
            resourceInputs["subnetId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["volumeSizeInGB"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NotebookInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a NotebookInstance resource.
 */
export interface NotebookInstanceArgs {
    acceleratorTypes?: pulumi.Input<pulumi.Input<string>[]>;
    additionalCodeRepositories?: pulumi.Input<pulumi.Input<string>[]>;
    defaultCodeRepository?: pulumi.Input<string>;
    directInternetAccess?: pulumi.Input<string>;
    instanceMetadataServiceConfiguration?: pulumi.Input<inputs.sagemaker.NotebookInstanceInstanceMetadataServiceConfigurationArgs>;
    instanceType: pulumi.Input<string>;
    kmsKeyId?: pulumi.Input<string>;
    lifecycleConfigName?: pulumi.Input<string>;
    notebookInstanceName?: pulumi.Input<string>;
    platformIdentifier?: pulumi.Input<string>;
    roleArn: pulumi.Input<string>;
    rootAccess?: pulumi.Input<string>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    subnetId?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.sagemaker.NotebookInstanceTagArgs>[]>;
    volumeSizeInGB?: pulumi.Input<number>;
}