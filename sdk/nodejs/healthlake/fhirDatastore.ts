// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * HealthLake FHIR Datastore
 */
export class FhirDatastore extends pulumi.CustomResource {
    /**
     * Get an existing FhirDatastore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FhirDatastore {
        return new FhirDatastore(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:healthlake:FhirDatastore';

    /**
     * Returns true if the given object is an instance of FhirDatastore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FhirDatastore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FhirDatastore.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<outputs.healthlake.FhirDatastoreCreatedAt>;
    /**
     * The Data Store ARN is generated during the creation of the Data Store and can be found in the output from the initial Data Store creation request.
     */
    public /*out*/ readonly datastoreArn!: pulumi.Output<string>;
    /**
     * The endpoint for the created Data Store.
     */
    public /*out*/ readonly datastoreEndpoint!: pulumi.Output<string>;
    /**
     * The Amazon generated Data Store id. This id is in the output from the initial Data Store creation call.
     */
    public /*out*/ readonly datastoreId!: pulumi.Output<string>;
    /**
     * The data store name (user-generated).
     */
    public readonly datastoreName!: pulumi.Output<string | undefined>;
    /**
     * The status of the FHIR Data Store. Possible statuses are ‘CREATING’, ‘ACTIVE’, ‘DELETING’, ‘DELETED’.
     */
    public /*out*/ readonly datastoreStatus!: pulumi.Output<enums.healthlake.FhirDatastoreDatastoreStatus>;
    /**
     * The FHIR release version supported by the data store. Current support is for version `R4` .
     */
    public readonly datastoreTypeVersion!: pulumi.Output<enums.healthlake.FhirDatastoreDatastoreTypeVersion>;
    /**
     * The identity provider configuration selected when the data store was created.
     */
    public readonly identityProviderConfiguration!: pulumi.Output<outputs.healthlake.FhirDatastoreIdentityProviderConfiguration | undefined>;
    /**
     * The preloaded Synthea data configuration for the data store.
     */
    public readonly preloadDataConfig!: pulumi.Output<outputs.healthlake.FhirDatastorePreloadDataConfig | undefined>;
    /**
     * The server-side encryption key configuration for a customer-provided encryption key specified for creating a data store.
     */
    public readonly sseConfiguration!: pulumi.Output<outputs.healthlake.FhirDatastoreSseConfiguration | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a FhirDatastore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FhirDatastoreArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.datastoreTypeVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datastoreTypeVersion'");
            }
            resourceInputs["datastoreName"] = args ? args.datastoreName : undefined;
            resourceInputs["datastoreTypeVersion"] = args ? args.datastoreTypeVersion : undefined;
            resourceInputs["identityProviderConfiguration"] = args ? args.identityProviderConfiguration : undefined;
            resourceInputs["preloadDataConfig"] = args ? args.preloadDataConfig : undefined;
            resourceInputs["sseConfiguration"] = args ? args.sseConfiguration : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["datastoreArn"] = undefined /*out*/;
            resourceInputs["datastoreEndpoint"] = undefined /*out*/;
            resourceInputs["datastoreId"] = undefined /*out*/;
            resourceInputs["datastoreStatus"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["datastoreArn"] = undefined /*out*/;
            resourceInputs["datastoreEndpoint"] = undefined /*out*/;
            resourceInputs["datastoreId"] = undefined /*out*/;
            resourceInputs["datastoreName"] = undefined /*out*/;
            resourceInputs["datastoreStatus"] = undefined /*out*/;
            resourceInputs["datastoreTypeVersion"] = undefined /*out*/;
            resourceInputs["identityProviderConfiguration"] = undefined /*out*/;
            resourceInputs["preloadDataConfig"] = undefined /*out*/;
            resourceInputs["sseConfiguration"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["datastoreName", "datastoreTypeVersion", "identityProviderConfiguration", "preloadDataConfig", "sseConfiguration"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(FhirDatastore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FhirDatastore resource.
 */
export interface FhirDatastoreArgs {
    /**
     * The data store name (user-generated).
     */
    datastoreName?: pulumi.Input<string>;
    /**
     * The FHIR release version supported by the data store. Current support is for version `R4` .
     */
    datastoreTypeVersion: pulumi.Input<enums.healthlake.FhirDatastoreDatastoreTypeVersion>;
    /**
     * The identity provider configuration selected when the data store was created.
     */
    identityProviderConfiguration?: pulumi.Input<inputs.healthlake.FhirDatastoreIdentityProviderConfigurationArgs>;
    /**
     * The preloaded Synthea data configuration for the data store.
     */
    preloadDataConfig?: pulumi.Input<inputs.healthlake.FhirDatastorePreloadDataConfigArgs>;
    /**
     * The server-side encryption key configuration for a customer-provided encryption key specified for creating a data store.
     */
    sseConfiguration?: pulumi.Input<inputs.healthlake.FhirDatastoreSseConfigurationArgs>;
    /**
     * An array of key-value pairs to apply to this resource.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
