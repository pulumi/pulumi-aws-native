// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Creates a custom DB engine version (CEV).
 */
export class CustomDbEngineVersion extends pulumi.CustomResource {
    /**
     * Get an existing CustomDbEngineVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomDbEngineVersion {
        return new CustomDbEngineVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:rds:CustomDbEngineVersion';

    /**
     * Returns true if the given object is an instance of CustomDbEngineVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomDbEngineVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomDbEngineVersion.__pulumiType;
    }

    /**
     * The name of an Amazon S3 bucket that contains database installation files for your CEV. For example, a valid bucket name is ``my-custom-installation-files``.
     */
    public readonly databaseInstallationFilesS3BucketName!: pulumi.Output<string | undefined>;
    /**
     * The Amazon S3 directory that contains the database installation files for your CEV. For example, a valid bucket name is ``123456789012/cev1``. If this setting isn't specified, no prefix is assumed.
     */
    public readonly databaseInstallationFilesS3Prefix!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the custom engine version.
     */
    public /*out*/ readonly dbEngineVersionArn!: pulumi.Output<string>;
    /**
     * An optional description of your CEV.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The database engine to use for your custom engine version (CEV).
     *  Valid values:
     *   +   ``custom-oracle-ee`` 
     *   +   ``custom-oracle-ee-cdb``
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * The name of your CEV. The name format is ``major version.customized_string``. For example, a valid CEV name is ``19.my_cev1``. This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of ``Engine`` and ``EngineVersion`` is unique per customer per Region.
     *   *Constraints:* Minimum length is 1. Maximum length is 60.
     *   *Pattern:* ``^[a-z0-9_.-]{1,60$``}
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * A value that indicates the ID of the AMI.
     */
    public readonly imageId!: pulumi.Output<string | undefined>;
    /**
     * The AWS KMS key identifier for an encrypted CEV. A symmetric encryption KMS key is required for RDS Custom, but optional for Amazon RDS.
     *  If you have an existing symmetric encryption KMS key in your account, you can use it with RDS Custom. No further action is necessary. If you don't already have a symmetric encryption KMS key in your account, follow the instructions in [Creating a symmetric encryption KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html#create-symmetric-cmk) in the *Key Management Service Developer Guide*.
     *  You can choose the same symmetric encryption key when you create a CEV and a DB instance, or choose different keys.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3. Specify the name/value pairs in a file or a quoted string. RDS Custom applies the patches in the order in which they are listed.
     *  The following JSON fields are valid:
     *   + MediaImportTemplateVersion Version of the CEV manifest. The date is in the format YYYY-MM-DD. + databaseInstallationFileNames Ordered list of installation files for the CEV. + opatchFileNames Ordered list of OPatch installers used for the Oracle DB engine. + psuRuPatchFileNames The PSU and RU patches for this CEV. + OtherPatchFileNames The patches that are not in the list of PSU and RU patches. Amazon RDS applies these patches after applying the PSU and RU patches. 
     *  For more information, see [Creating the CEV manifest](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.preparing.manifest) in the *Amazon RDS User Guide*.
     */
    public readonly manifest!: pulumi.Output<string | undefined>;
    /**
     * The ARN of a CEV to use as a source for creating a new CEV. You can specify a different Amazon Machine Imagine (AMI) by using either ``Source`` or ``UseAwsProvidedLatestImage``. You can't specify a different JSON manifest when you specify ``SourceCustomDbEngineVersionIdentifier``.
     */
    public readonly sourceCustomDbEngineVersionIdentifier!: pulumi.Output<string | undefined>;
    /**
     * A value that indicates the status of a custom engine version (CEV).
     */
    public readonly status!: pulumi.Output<enums.rds.CustomDbEngineVersionStatus | undefined>;
    /**
     * A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * Specifies whether to use the latest service-provided Amazon Machine Image (AMI) for the CEV. If you specify ``UseAwsProvidedLatestImage``, you can't also specify ``ImageId``.
     */
    public readonly useAwsProvidedLatestImage!: pulumi.Output<boolean | undefined>;

    /**
     * Create a CustomDbEngineVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomDbEngineVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.engineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineVersion'");
            }
            resourceInputs["databaseInstallationFilesS3BucketName"] = args ? args.databaseInstallationFilesS3BucketName : undefined;
            resourceInputs["databaseInstallationFilesS3Prefix"] = args ? args.databaseInstallationFilesS3Prefix : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["manifest"] = args ? args.manifest : undefined;
            resourceInputs["sourceCustomDbEngineVersionIdentifier"] = args ? args.sourceCustomDbEngineVersionIdentifier : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["useAwsProvidedLatestImage"] = args ? args.useAwsProvidedLatestImage : undefined;
            resourceInputs["dbEngineVersionArn"] = undefined /*out*/;
        } else {
            resourceInputs["databaseInstallationFilesS3BucketName"] = undefined /*out*/;
            resourceInputs["databaseInstallationFilesS3Prefix"] = undefined /*out*/;
            resourceInputs["dbEngineVersionArn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["engine"] = undefined /*out*/;
            resourceInputs["engineVersion"] = undefined /*out*/;
            resourceInputs["imageId"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["manifest"] = undefined /*out*/;
            resourceInputs["sourceCustomDbEngineVersionIdentifier"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["useAwsProvidedLatestImage"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["databaseInstallationFilesS3BucketName", "databaseInstallationFilesS3Prefix", "engine", "engineVersion", "imageId", "kmsKeyId", "manifest", "sourceCustomDbEngineVersionIdentifier", "useAwsProvidedLatestImage"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(CustomDbEngineVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomDbEngineVersion resource.
 */
export interface CustomDbEngineVersionArgs {
    /**
     * The name of an Amazon S3 bucket that contains database installation files for your CEV. For example, a valid bucket name is ``my-custom-installation-files``.
     */
    databaseInstallationFilesS3BucketName?: pulumi.Input<string>;
    /**
     * The Amazon S3 directory that contains the database installation files for your CEV. For example, a valid bucket name is ``123456789012/cev1``. If this setting isn't specified, no prefix is assumed.
     */
    databaseInstallationFilesS3Prefix?: pulumi.Input<string>;
    /**
     * An optional description of your CEV.
     */
    description?: pulumi.Input<string>;
    /**
     * The database engine to use for your custom engine version (CEV).
     *  Valid values:
     *   +   ``custom-oracle-ee`` 
     *   +   ``custom-oracle-ee-cdb``
     */
    engine: pulumi.Input<string>;
    /**
     * The name of your CEV. The name format is ``major version.customized_string``. For example, a valid CEV name is ``19.my_cev1``. This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of ``Engine`` and ``EngineVersion`` is unique per customer per Region.
     *   *Constraints:* Minimum length is 1. Maximum length is 60.
     *   *Pattern:* ``^[a-z0-9_.-]{1,60$``}
     */
    engineVersion: pulumi.Input<string>;
    /**
     * A value that indicates the ID of the AMI.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The AWS KMS key identifier for an encrypted CEV. A symmetric encryption KMS key is required for RDS Custom, but optional for Amazon RDS.
     *  If you have an existing symmetric encryption KMS key in your account, you can use it with RDS Custom. No further action is necessary. If you don't already have a symmetric encryption KMS key in your account, follow the instructions in [Creating a symmetric encryption KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html#create-symmetric-cmk) in the *Key Management Service Developer Guide*.
     *  You can choose the same symmetric encryption key when you create a CEV and a DB instance, or choose different keys.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3. Specify the name/value pairs in a file or a quoted string. RDS Custom applies the patches in the order in which they are listed.
     *  The following JSON fields are valid:
     *   + MediaImportTemplateVersion Version of the CEV manifest. The date is in the format YYYY-MM-DD. + databaseInstallationFileNames Ordered list of installation files for the CEV. + opatchFileNames Ordered list of OPatch installers used for the Oracle DB engine. + psuRuPatchFileNames The PSU and RU patches for this CEV. + OtherPatchFileNames The patches that are not in the list of PSU and RU patches. Amazon RDS applies these patches after applying the PSU and RU patches. 
     *  For more information, see [Creating the CEV manifest](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.preparing.manifest) in the *Amazon RDS User Guide*.
     */
    manifest?: pulumi.Input<string>;
    /**
     * The ARN of a CEV to use as a source for creating a new CEV. You can specify a different Amazon Machine Imagine (AMI) by using either ``Source`` or ``UseAwsProvidedLatestImage``. You can't specify a different JSON manifest when you specify ``SourceCustomDbEngineVersionIdentifier``.
     */
    sourceCustomDbEngineVersionIdentifier?: pulumi.Input<string>;
    /**
     * A value that indicates the status of a custom engine version (CEV).
     */
    status?: pulumi.Input<enums.rds.CustomDbEngineVersionStatus>;
    /**
     * A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * Specifies whether to use the latest service-provided Amazon Machine Image (AMI) for the CEV. If you specify ``UseAwsProvidedLatestImage``, you can't also specify ``ImageId``.
     */
    useAwsProvidedLatestImage?: pulumi.Input<boolean>;
}
