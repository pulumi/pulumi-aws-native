// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Bedrock::DataSource Resource Type
 */
export class DataSource extends pulumi.CustomResource {
    /**
     * Get an existing DataSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataSource {
        return new DataSource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:bedrock:DataSource';

    /**
     * Returns true if the given object is an instance of DataSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSource.__pulumiType;
    }

    /**
     * The time at which the data source was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The data deletion policy for the data source.
     */
    public readonly dataDeletionPolicy!: pulumi.Output<enums.bedrock.DataSourceDataDeletionPolicy | undefined>;
    /**
     * The connection configuration for the data source.
     */
    public readonly dataSourceConfiguration!: pulumi.Output<outputs.bedrock.DataSourceConfiguration>;
    /**
     * Identifier for a resource.
     */
    public /*out*/ readonly dataSourceId!: pulumi.Output<string>;
    /**
     * The status of the data source. The following statuses are possible:
     *
     * - Available – The data source has been created and is ready for ingestion into the knowledge base.
     * - Deleting – The data source is being deleted.
     */
    public /*out*/ readonly dataSourceStatus!: pulumi.Output<enums.bedrock.DataSourceStatus>;
    /**
     * Description of the Resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The details of the failure reasons related to the data source.
     */
    public /*out*/ readonly failureReasons!: pulumi.Output<string[]>;
    /**
     * The unique identifier of the knowledge base to which to add the data source.
     */
    public readonly knowledgeBaseId!: pulumi.Output<string>;
    /**
     * The name of the data source.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Contains details about the configuration of the server-side encryption.
     */
    public readonly serverSideEncryptionConfiguration!: pulumi.Output<outputs.bedrock.DataSourceServerSideEncryptionConfiguration | undefined>;
    /**
     * The time at which the knowledge base was last updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Contains details about how to ingest the documents in the data source.
     */
    public readonly vectorIngestionConfiguration!: pulumi.Output<outputs.bedrock.DataSourceVectorIngestionConfiguration | undefined>;

    /**
     * Create a DataSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataSourceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dataSourceConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataSourceConfiguration'");
            }
            if ((!args || args.knowledgeBaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'knowledgeBaseId'");
            }
            resourceInputs["dataDeletionPolicy"] = args ? args.dataDeletionPolicy : undefined;
            resourceInputs["dataSourceConfiguration"] = args ? args.dataSourceConfiguration : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["knowledgeBaseId"] = args ? args.knowledgeBaseId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serverSideEncryptionConfiguration"] = args ? args.serverSideEncryptionConfiguration : undefined;
            resourceInputs["vectorIngestionConfiguration"] = args ? args.vectorIngestionConfiguration : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["dataSourceId"] = undefined /*out*/;
            resourceInputs["dataSourceStatus"] = undefined /*out*/;
            resourceInputs["failureReasons"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["dataDeletionPolicy"] = undefined /*out*/;
            resourceInputs["dataSourceConfiguration"] = undefined /*out*/;
            resourceInputs["dataSourceId"] = undefined /*out*/;
            resourceInputs["dataSourceStatus"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["failureReasons"] = undefined /*out*/;
            resourceInputs["knowledgeBaseId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serverSideEncryptionConfiguration"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["vectorIngestionConfiguration"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["dataSourceConfiguration.type", "knowledgeBaseId", "vectorIngestionConfiguration.chunkingConfiguration", "vectorIngestionConfiguration.parsingConfiguration"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DataSource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataSource resource.
 */
export interface DataSourceArgs {
    /**
     * The data deletion policy for the data source.
     */
    dataDeletionPolicy?: pulumi.Input<enums.bedrock.DataSourceDataDeletionPolicy>;
    /**
     * The connection configuration for the data source.
     */
    dataSourceConfiguration: pulumi.Input<inputs.bedrock.DataSourceConfigurationArgs>;
    /**
     * Description of the Resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The unique identifier of the knowledge base to which to add the data source.
     */
    knowledgeBaseId: pulumi.Input<string>;
    /**
     * The name of the data source.
     */
    name?: pulumi.Input<string>;
    /**
     * Contains details about the configuration of the server-side encryption.
     */
    serverSideEncryptionConfiguration?: pulumi.Input<inputs.bedrock.DataSourceServerSideEncryptionConfigurationArgs>;
    /**
     * Contains details about how to ingest the documents in the data source.
     */
    vectorIngestionConfiguration?: pulumi.Input<inputs.bedrock.DataSourceVectorIngestionConfigurationArgs>;
}
