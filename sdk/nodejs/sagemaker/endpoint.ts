// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::Endpoint
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    /**
     * Specifies deployment configuration for updating the SageMaker endpoint. Includes rollback and update policies.
     */
    public readonly deploymentConfig!: pulumi.Output<outputs.sagemaker.EndpointDeploymentConfig | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the endpoint.
     */
    public /*out*/ readonly endpointArn!: pulumi.Output<string>;
    /**
     * The name of the endpoint configuration for the SageMaker endpoint. This is a required property.
     */
    public readonly endpointConfigName!: pulumi.Output<string>;
    /**
     * The name of the SageMaker endpoint. This name must be unique within an AWS Region.
     */
    public readonly endpointName!: pulumi.Output<string>;
    /**
     * Specifies a list of variant properties that you want to exclude when updating an endpoint.
     */
    public readonly excludeRetainedVariantProperties!: pulumi.Output<outputs.sagemaker.EndpointVariantProperty[] | undefined>;
    /**
     * When set to true, retains all variant properties for an endpoint when it is updated.
     */
    public readonly retainAllVariantProperties!: pulumi.Output<boolean | undefined>;
    /**
     * When set to true, retains the deployment configuration during endpoint updates.
     */
    public readonly retainDeploymentConfig!: pulumi.Output<boolean | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.endpointConfigName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointConfigName'");
            }
            resourceInputs["deploymentConfig"] = args ? args.deploymentConfig : undefined;
            resourceInputs["endpointConfigName"] = args ? args.endpointConfigName : undefined;
            resourceInputs["endpointName"] = args ? args.endpointName : undefined;
            resourceInputs["excludeRetainedVariantProperties"] = args ? args.excludeRetainedVariantProperties : undefined;
            resourceInputs["retainAllVariantProperties"] = args ? args.retainAllVariantProperties : undefined;
            resourceInputs["retainDeploymentConfig"] = args ? args.retainDeploymentConfig : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["endpointArn"] = undefined /*out*/;
        } else {
            resourceInputs["deploymentConfig"] = undefined /*out*/;
            resourceInputs["endpointArn"] = undefined /*out*/;
            resourceInputs["endpointConfigName"] = undefined /*out*/;
            resourceInputs["endpointName"] = undefined /*out*/;
            resourceInputs["excludeRetainedVariantProperties"] = undefined /*out*/;
            resourceInputs["retainAllVariantProperties"] = undefined /*out*/;
            resourceInputs["retainDeploymentConfig"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["endpointName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Endpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * Specifies deployment configuration for updating the SageMaker endpoint. Includes rollback and update policies.
     */
    deploymentConfig?: pulumi.Input<inputs.sagemaker.EndpointDeploymentConfigArgs>;
    /**
     * The name of the endpoint configuration for the SageMaker endpoint. This is a required property.
     */
    endpointConfigName: pulumi.Input<string>;
    /**
     * The name of the SageMaker endpoint. This name must be unique within an AWS Region.
     */
    endpointName?: pulumi.Input<string>;
    /**
     * Specifies a list of variant properties that you want to exclude when updating an endpoint.
     */
    excludeRetainedVariantProperties?: pulumi.Input<pulumi.Input<inputs.sagemaker.EndpointVariantPropertyArgs>[]>;
    /**
     * When set to true, retains all variant properties for an endpoint when it is updated.
     */
    retainAllVariantProperties?: pulumi.Input<boolean>;
    /**
     * When set to true, retains the deployment configuration during endpoint updates.
     */
    retainDeploymentConfig?: pulumi.Input<boolean>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
