// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::LocationAzureBlob.
 */
export class LocationAzureBlob extends pulumi.CustomResource {
    /**
     * Get an existing LocationAzureBlob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LocationAzureBlob {
        return new LocationAzureBlob(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:datasync:LocationAzureBlob';

    /**
     * Returns true if the given object is an instance of LocationAzureBlob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocationAzureBlob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocationAzureBlob.__pulumiType;
    }

    /**
     * The Amazon Resource Names (ARNs) of agents to use for an Azure Blob Location.
     */
    public readonly agentArns!: pulumi.Output<string[]>;
    /**
     * Specifies an access tier for the objects you're transferring into your Azure Blob Storage container.
     */
    public readonly azureAccessTier!: pulumi.Output<enums.datasync.LocationAzureBlobAzureAccessTier | undefined>;
    /**
     * The specific authentication type that you want DataSync to use to access your Azure Blob Container.
     */
    public readonly azureBlobAuthenticationType!: pulumi.Output<enums.datasync.LocationAzureBlobAzureBlobAuthenticationType>;
    /**
     * The URL of the Azure Blob container that was described.
     */
    public readonly azureBlobContainerUrl!: pulumi.Output<string | undefined>;
    public readonly azureBlobSasConfiguration!: pulumi.Output<outputs.datasync.LocationAzureBlobAzureBlobSasConfiguration | undefined>;
    /**
     * Specifies a blob type for the objects you're transferring into your Azure Blob Storage container.
     */
    public readonly azureBlobType!: pulumi.Output<enums.datasync.LocationAzureBlobAzureBlobType | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the Azure Blob Location that is created.
     */
    public /*out*/ readonly locationArn!: pulumi.Output<string>;
    /**
     * The URL of the Azure Blob Location that was described.
     */
    public /*out*/ readonly locationUri!: pulumi.Output<string>;
    /**
     * The subdirectory in the Azure Blob Container that is used to read data from the Azure Blob Source Location.
     */
    public readonly subdirectory!: pulumi.Output<string | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.datasync.LocationAzureBlobTag[] | undefined>;

    /**
     * Create a LocationAzureBlob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocationAzureBlobArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentArns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentArns'");
            }
            if ((!args || args.azureBlobAuthenticationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'azureBlobAuthenticationType'");
            }
            resourceInputs["agentArns"] = args ? args.agentArns : undefined;
            resourceInputs["azureAccessTier"] = args ? args.azureAccessTier : undefined;
            resourceInputs["azureBlobAuthenticationType"] = args ? args.azureBlobAuthenticationType : undefined;
            resourceInputs["azureBlobContainerUrl"] = args ? args.azureBlobContainerUrl : undefined;
            resourceInputs["azureBlobSasConfiguration"] = args ? args.azureBlobSasConfiguration : undefined;
            resourceInputs["azureBlobType"] = args ? args.azureBlobType : undefined;
            resourceInputs["subdirectory"] = args ? args.subdirectory : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["locationArn"] = undefined /*out*/;
            resourceInputs["locationUri"] = undefined /*out*/;
        } else {
            resourceInputs["agentArns"] = undefined /*out*/;
            resourceInputs["azureAccessTier"] = undefined /*out*/;
            resourceInputs["azureBlobAuthenticationType"] = undefined /*out*/;
            resourceInputs["azureBlobContainerUrl"] = undefined /*out*/;
            resourceInputs["azureBlobSasConfiguration"] = undefined /*out*/;
            resourceInputs["azureBlobType"] = undefined /*out*/;
            resourceInputs["locationArn"] = undefined /*out*/;
            resourceInputs["locationUri"] = undefined /*out*/;
            resourceInputs["subdirectory"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LocationAzureBlob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a LocationAzureBlob resource.
 */
export interface LocationAzureBlobArgs {
    /**
     * The Amazon Resource Names (ARNs) of agents to use for an Azure Blob Location.
     */
    agentArns: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies an access tier for the objects you're transferring into your Azure Blob Storage container.
     */
    azureAccessTier?: pulumi.Input<enums.datasync.LocationAzureBlobAzureAccessTier>;
    /**
     * The specific authentication type that you want DataSync to use to access your Azure Blob Container.
     */
    azureBlobAuthenticationType: pulumi.Input<enums.datasync.LocationAzureBlobAzureBlobAuthenticationType>;
    /**
     * The URL of the Azure Blob container that was described.
     */
    azureBlobContainerUrl?: pulumi.Input<string>;
    azureBlobSasConfiguration?: pulumi.Input<inputs.datasync.LocationAzureBlobAzureBlobSasConfigurationArgs>;
    /**
     * Specifies a blob type for the objects you're transferring into your Azure Blob Storage container.
     */
    azureBlobType?: pulumi.Input<enums.datasync.LocationAzureBlobAzureBlobType>;
    /**
     * The subdirectory in the Azure Blob Container that is used to read data from the Azure Blob Source Location.
     */
    subdirectory?: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.datasync.LocationAzureBlobTagArgs>[]>;
}