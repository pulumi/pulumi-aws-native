// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:finspace:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * AWS account ID associated with the Environment
     */
    public /*out*/ readonly awsAccountId!: pulumi.Output<string>;
    /**
     * ARNs of FinSpace Data Bundles to install
     */
    public readonly dataBundles!: pulumi.Output<string[] | undefined>;
    /**
     * ID for FinSpace created account used to store Environment artifacts
     */
    public /*out*/ readonly dedicatedServiceAccountId!: pulumi.Output<string>;
    /**
     * Description of the Environment
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ARN of the Environment
     */
    public /*out*/ readonly environmentArn!: pulumi.Output<string>;
    /**
     * Unique identifier for representing FinSpace Environment
     */
    public /*out*/ readonly environmentId!: pulumi.Output<string>;
    /**
     * URL used to login to the Environment
     */
    public /*out*/ readonly environmentUrl!: pulumi.Output<string>;
    /**
     * Federation mode used with the Environment
     */
    public readonly federationMode!: pulumi.Output<enums.finspace.EnvironmentFederationMode | undefined>;
    public readonly federationParameters!: pulumi.Output<outputs.finspace.EnvironmentFederationParameters | undefined>;
    /**
     * KMS key used to encrypt customer data within FinSpace Environment infrastructure
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Name of the Environment
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * SageMaker Studio Domain URL associated with the Environment
     */
    public /*out*/ readonly sageMakerStudioDomainUrl!: pulumi.Output<string>;
    /**
     * State of the Environment
     */
    public /*out*/ readonly status!: pulumi.Output<enums.finspace.EnvironmentStatus>;
    public readonly superuserParameters!: pulumi.Output<outputs.finspace.EnvironmentSuperuserParameters | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.finspace.EnvironmentTag[] | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["dataBundles"] = args ? args.dataBundles : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["federationMode"] = args ? args.federationMode : undefined;
            resourceInputs["federationParameters"] = args ? args.federationParameters : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["superuserParameters"] = args ? args.superuserParameters : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["awsAccountId"] = undefined /*out*/;
            resourceInputs["dedicatedServiceAccountId"] = undefined /*out*/;
            resourceInputs["environmentArn"] = undefined /*out*/;
            resourceInputs["environmentId"] = undefined /*out*/;
            resourceInputs["environmentUrl"] = undefined /*out*/;
            resourceInputs["sageMakerStudioDomainUrl"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["awsAccountId"] = undefined /*out*/;
            resourceInputs["dataBundles"] = undefined /*out*/;
            resourceInputs["dedicatedServiceAccountId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["environmentArn"] = undefined /*out*/;
            resourceInputs["environmentId"] = undefined /*out*/;
            resourceInputs["environmentUrl"] = undefined /*out*/;
            resourceInputs["federationMode"] = undefined /*out*/;
            resourceInputs["federationParameters"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["sageMakerStudioDomainUrl"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["superuserParameters"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Environment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * ARNs of FinSpace Data Bundles to install
     */
    dataBundles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of the Environment
     */
    description?: pulumi.Input<string>;
    /**
     * Federation mode used with the Environment
     */
    federationMode?: pulumi.Input<enums.finspace.EnvironmentFederationMode>;
    federationParameters?: pulumi.Input<inputs.finspace.EnvironmentFederationParametersArgs>;
    /**
     * KMS key used to encrypt customer data within FinSpace Environment infrastructure
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Name of the Environment
     */
    name?: pulumi.Input<string>;
    superuserParameters?: pulumi.Input<inputs.finspace.EnvironmentSuperuserParametersArgs>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.finspace.EnvironmentTagArgs>[]>;
}