// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Creates an Amazon FSx for Lustre data repository association (DRA). A data repository association is a link between a directory on the file system and an Amazon S3 bucket or prefix. You can have a maximum of 8 data repository associations on a file system. Data repository associations are supported on all FSx for Lustre 2.12 and newer file systems, excluding ``scratch_1`` deployment type.
 *  Each data repository association must have a unique Amazon FSx file system directory and a unique S3 bucket or prefix associated with it. You can configure a data repository association for automatic import only, for automatic export only, or for both. To learn more about linking a data repository to your file system, see [Linking your file system to an S3 bucket](https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html).
 *
 * ## Example Usage
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const config = new pulumi.Config();
 * const fsId = config.require("fsId");
 * const draIdExportName = config.require("draIdExportName");
 * const fileSystemPath = config.require("fileSystemPath");
 * const importedFileChunkSize = config.require("importedFileChunkSize");
 * const testDRA = new aws_native.fsx.DataRepositoryAssociation("testDRA", {
 *     fileSystemId: fsId,
 *     fileSystemPath: fileSystemPath,
 *     dataRepositoryPath: "s3://example-bucket",
 *     batchImportMetaDataOnCreate: true,
 *     importedFileChunkSize: importedFileChunkSize,
 *     s3: {
 *         autoImportPolicy: {
 *             events: [
 *                 aws_native.fsx.DataRepositoryAssociationEventType.New,
 *                 aws_native.fsx.DataRepositoryAssociationEventType.Changed,
 *                 aws_native.fsx.DataRepositoryAssociationEventType.Deleted,
 *             ],
 *         },
 *         autoExportPolicy: {
 *             events: [
 *                 aws_native.fsx.DataRepositoryAssociationEventType.New,
 *                 aws_native.fsx.DataRepositoryAssociationEventType.Changed,
 *                 aws_native.fsx.DataRepositoryAssociationEventType.Deleted,
 *             ],
 *         },
 *     },
 *     tags: [{
 *         key: "Location",
 *         value: "Boston",
 *     }],
 * });
 * export const draId = testDRA.id;
 *
 * ```
 */
export class DataRepositoryAssociation extends pulumi.CustomResource {
    /**
     * Get an existing DataRepositoryAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataRepositoryAssociation {
        return new DataRepositoryAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:fsx:DataRepositoryAssociation';

    /**
     * Returns true if the given object is an instance of DataRepositoryAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataRepositoryAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataRepositoryAssociation.__pulumiType;
    }

    /**
     * Returns the data repository association's system generated Association ID.
     *
     * Example: `dra-abcdef0123456789d`
     */
    public /*out*/ readonly associationId!: pulumi.Output<string>;
    /**
     * A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to ``true``.
     */
    public readonly batchImportMetaDataOnCreate!: pulumi.Output<boolean | undefined>;
    /**
     * The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format ``s3://myBucket/myPrefix/``. This path specifies where in the S3 data repository files will be imported from or exported to.
     */
    public readonly dataRepositoryPath!: pulumi.Output<string>;
    /**
     * The ID of the file system on which the data repository association is configured.
     */
    public readonly fileSystemId!: pulumi.Output<string>;
    /**
     * A path on the Amazon FSx for Lustre file system that points to a high-level directory (such as ``/ns1/``) or subdirectory (such as ``/ns1/subdir/``) that will be mapped 1-1 with ``DataRepositoryPath``. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path ``/ns1/``, then you cannot link another data repository with file system path ``/ns1/ns2``.
     *  This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
     *   If you specify only a forward slash (``/``) as the file system path, you can link only one data repository to the file system. You can only specify "/" as the file system path for the first data repository associated with a file system.
     */
    public readonly fileSystemPath!: pulumi.Output<string>;
    /**
     * For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system or cache.
     *  The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.
     */
    public readonly importedFileChunkSize!: pulumi.Output<number | undefined>;
    /**
     * Returns the data repository association's Amazon Resource Name (ARN).
     *
     * Example: `arn:aws:fsx:us-east-1:111122223333:association/fs-abc012345def6789a/dra-abcdef0123456789b`
     */
    public /*out*/ readonly resourceArn!: pulumi.Output<string>;
    /**
     * The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
     */
    public readonly s3!: pulumi.Output<outputs.fsx.DataRepositoryAssociationS3 | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     *  For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a DataRepositoryAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataRepositoryAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dataRepositoryPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataRepositoryPath'");
            }
            if ((!args || args.fileSystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileSystemId'");
            }
            if ((!args || args.fileSystemPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileSystemPath'");
            }
            resourceInputs["batchImportMetaDataOnCreate"] = args ? args.batchImportMetaDataOnCreate : undefined;
            resourceInputs["dataRepositoryPath"] = args ? args.dataRepositoryPath : undefined;
            resourceInputs["fileSystemId"] = args ? args.fileSystemId : undefined;
            resourceInputs["fileSystemPath"] = args ? args.fileSystemPath : undefined;
            resourceInputs["importedFileChunkSize"] = args ? args.importedFileChunkSize : undefined;
            resourceInputs["s3"] = args ? args.s3 : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["associationId"] = undefined /*out*/;
            resourceInputs["resourceArn"] = undefined /*out*/;
        } else {
            resourceInputs["associationId"] = undefined /*out*/;
            resourceInputs["batchImportMetaDataOnCreate"] = undefined /*out*/;
            resourceInputs["dataRepositoryPath"] = undefined /*out*/;
            resourceInputs["fileSystemId"] = undefined /*out*/;
            resourceInputs["fileSystemPath"] = undefined /*out*/;
            resourceInputs["importedFileChunkSize"] = undefined /*out*/;
            resourceInputs["resourceArn"] = undefined /*out*/;
            resourceInputs["s3"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["batchImportMetaDataOnCreate", "dataRepositoryPath", "fileSystemId", "fileSystemPath"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DataRepositoryAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataRepositoryAssociation resource.
 */
export interface DataRepositoryAssociationArgs {
    /**
     * A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to ``true``.
     */
    batchImportMetaDataOnCreate?: pulumi.Input<boolean>;
    /**
     * The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format ``s3://myBucket/myPrefix/``. This path specifies where in the S3 data repository files will be imported from or exported to.
     */
    dataRepositoryPath: pulumi.Input<string>;
    /**
     * The ID of the file system on which the data repository association is configured.
     */
    fileSystemId: pulumi.Input<string>;
    /**
     * A path on the Amazon FSx for Lustre file system that points to a high-level directory (such as ``/ns1/``) or subdirectory (such as ``/ns1/subdir/``) that will be mapped 1-1 with ``DataRepositoryPath``. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path ``/ns1/``, then you cannot link another data repository with file system path ``/ns1/ns2``.
     *  This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
     *   If you specify only a forward slash (``/``) as the file system path, you can link only one data repository to the file system. You can only specify "/" as the file system path for the first data repository associated with a file system.
     */
    fileSystemPath: pulumi.Input<string>;
    /**
     * For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system or cache.
     *  The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.
     */
    importedFileChunkSize?: pulumi.Input<number>;
    /**
     * The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
     */
    s3?: pulumi.Input<inputs.fsx.DataRepositoryAssociationS3Args>;
    /**
     * An array of key-value pairs to apply to this resource.
     *  For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
