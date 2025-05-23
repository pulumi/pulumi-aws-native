// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IoTAnalytics::Datastore
 *
 * ## Example Usage
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const datastore = new aws_native.iotanalytics.Datastore("datastore", {datastoreName: "SimpleDatastore"});
 *
 * ```
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const datastore = new aws_native.iotanalytics.Datastore("datastore", {datastoreName: "SimpleDatastore"});
 *
 * ```
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const datastore = new aws_native.iotanalytics.Datastore("datastore", {
 *     datastoreName: "ComplexDatastore",
 *     retentionPeriod: {
 *         unlimited: false,
 *         numberOfDays: 10,
 *     },
 *     tags: [
 *         {
 *             key: "keyname1",
 *             value: "value1",
 *         },
 *         {
 *             key: "keyname2",
 *             value: "value2",
 *         },
 *     ],
 * });
 *
 * ```
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const datastore = new aws_native.iotanalytics.Datastore("datastore", {
 *     datastoreName: "ComplexDatastore",
 *     retentionPeriod: {
 *         unlimited: false,
 *         numberOfDays: 10,
 *     },
 *     tags: [
 *         {
 *             key: "keyname1",
 *             value: "value1",
 *         },
 *         {
 *             key: "keyname2",
 *             value: "value2",
 *         },
 *     ],
 * });
 *
 * ```
 */
export class Datastore extends pulumi.CustomResource {
    /**
     * Get an existing Datastore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Datastore {
        return new Datastore(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iotanalytics:Datastore';

    /**
     * Returns true if the given object is an instance of Datastore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Datastore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Datastore.__pulumiType;
    }

    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The name of the data store.
     */
    public readonly datastoreName!: pulumi.Output<string | undefined>;
    /**
     * Information about the partition dimensions in a data store.
     */
    public readonly datastorePartitions!: pulumi.Output<outputs.iotanalytics.DatastorePartitions | undefined>;
    /**
     * Where data store data is stored.
     */
    public readonly datastoreStorage!: pulumi.Output<outputs.iotanalytics.DatastoreStorage | undefined>;
    /**
     * Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
     *
     * The default file format is JSON. You can specify only one format.
     *
     * You can't change the file format after you create the data store.
     */
    public readonly fileFormatConfiguration!: pulumi.Output<outputs.iotanalytics.DatastoreFileFormatConfiguration | undefined>;
    /**
     * How long, in days, message data is kept for the data store. When `customerManagedS3` storage is selected, this parameter is ignored.
     */
    public readonly retentionPeriod!: pulumi.Output<outputs.iotanalytics.DatastoreRetentionPeriod | undefined>;
    /**
     * Metadata which can be used to manage the data store.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a Datastore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatastoreArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["datastoreName"] = args ? args.datastoreName : undefined;
            resourceInputs["datastorePartitions"] = args ? args.datastorePartitions : undefined;
            resourceInputs["datastoreStorage"] = args ? args.datastoreStorage : undefined;
            resourceInputs["fileFormatConfiguration"] = args ? args.fileFormatConfiguration : undefined;
            resourceInputs["retentionPeriod"] = args ? args.retentionPeriod : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["datastoreName"] = undefined /*out*/;
            resourceInputs["datastorePartitions"] = undefined /*out*/;
            resourceInputs["datastoreStorage"] = undefined /*out*/;
            resourceInputs["fileFormatConfiguration"] = undefined /*out*/;
            resourceInputs["retentionPeriod"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["datastoreName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Datastore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Datastore resource.
 */
export interface DatastoreArgs {
    /**
     * The name of the data store.
     */
    datastoreName?: pulumi.Input<string>;
    /**
     * Information about the partition dimensions in a data store.
     */
    datastorePartitions?: pulumi.Input<inputs.iotanalytics.DatastorePartitionsArgs>;
    /**
     * Where data store data is stored.
     */
    datastoreStorage?: pulumi.Input<inputs.iotanalytics.DatastoreStorageArgs>;
    /**
     * Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://docs.aws.amazon.com/https://parquet.apache.org/) .
     *
     * The default file format is JSON. You can specify only one format.
     *
     * You can't change the file format after you create the data store.
     */
    fileFormatConfiguration?: pulumi.Input<inputs.iotanalytics.DatastoreFileFormatConfigurationArgs>;
    /**
     * How long, in days, message data is kept for the data store. When `customerManagedS3` storage is selected, this parameter is ignored.
     */
    retentionPeriod?: pulumi.Input<inputs.iotanalytics.DatastoreRetentionPeriodArgs>;
    /**
     * Metadata which can be used to manage the data store.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
