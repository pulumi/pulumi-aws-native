// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DMS::DataMigration.
 */
export class DataMigration extends pulumi.CustomResource {
    /**
     * Get an existing DataMigration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataMigration {
        return new DataMigration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:dms:DataMigration';

    /**
     * Returns true if the given object is an instance of DataMigration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataMigration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataMigration.__pulumiType;
    }

    /**
     * The property describes an ARN of the data migration.
     */
    public /*out*/ readonly dataMigrationArn!: pulumi.Output<string>;
    /**
     * The property describes the create time of the data migration.
     */
    public /*out*/ readonly dataMigrationCreateTime!: pulumi.Output<string>;
    /**
     * The property describes an ARN of the data migration.
     */
    public readonly dataMigrationIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The property describes a name to identify the data migration.
     */
    public readonly dataMigrationName!: pulumi.Output<string | undefined>;
    /**
     * The property describes the settings for the data migration.
     */
    public readonly dataMigrationSettings!: pulumi.Output<outputs.dms.DataMigrationSettings | undefined>;
    /**
     * The property describes the type of migration.
     */
    public readonly dataMigrationType!: pulumi.Output<enums.dms.DataMigrationType>;
    /**
     * The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn
     */
    public readonly migrationProjectIdentifier!: pulumi.Output<string>;
    /**
     * The property describes Amazon Resource Name (ARN) of the service access role.
     */
    public readonly serviceAccessRoleArn!: pulumi.Output<string>;
    /**
     * The property describes the settings for the data migration.
     */
    public readonly sourceDataSettings!: pulumi.Output<outputs.dms.DataMigrationSourceDataSettings[] | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a DataMigration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataMigrationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dataMigrationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataMigrationType'");
            }
            if ((!args || args.migrationProjectIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'migrationProjectIdentifier'");
            }
            if ((!args || args.serviceAccessRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccessRoleArn'");
            }
            resourceInputs["dataMigrationIdentifier"] = args ? args.dataMigrationIdentifier : undefined;
            resourceInputs["dataMigrationName"] = args ? args.dataMigrationName : undefined;
            resourceInputs["dataMigrationSettings"] = args ? args.dataMigrationSettings : undefined;
            resourceInputs["dataMigrationType"] = args ? args.dataMigrationType : undefined;
            resourceInputs["migrationProjectIdentifier"] = args ? args.migrationProjectIdentifier : undefined;
            resourceInputs["serviceAccessRoleArn"] = args ? args.serviceAccessRoleArn : undefined;
            resourceInputs["sourceDataSettings"] = args ? args.sourceDataSettings : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["dataMigrationArn"] = undefined /*out*/;
            resourceInputs["dataMigrationCreateTime"] = undefined /*out*/;
        } else {
            resourceInputs["dataMigrationArn"] = undefined /*out*/;
            resourceInputs["dataMigrationCreateTime"] = undefined /*out*/;
            resourceInputs["dataMigrationIdentifier"] = undefined /*out*/;
            resourceInputs["dataMigrationName"] = undefined /*out*/;
            resourceInputs["dataMigrationSettings"] = undefined /*out*/;
            resourceInputs["dataMigrationType"] = undefined /*out*/;
            resourceInputs["migrationProjectIdentifier"] = undefined /*out*/;
            resourceInputs["serviceAccessRoleArn"] = undefined /*out*/;
            resourceInputs["sourceDataSettings"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DataMigration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataMigration resource.
 */
export interface DataMigrationArgs {
    /**
     * The property describes an ARN of the data migration.
     */
    dataMigrationIdentifier?: pulumi.Input<string>;
    /**
     * The property describes a name to identify the data migration.
     */
    dataMigrationName?: pulumi.Input<string>;
    /**
     * The property describes the settings for the data migration.
     */
    dataMigrationSettings?: pulumi.Input<inputs.dms.DataMigrationSettingsArgs>;
    /**
     * The property describes the type of migration.
     */
    dataMigrationType: pulumi.Input<enums.dms.DataMigrationType>;
    /**
     * The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn
     */
    migrationProjectIdentifier: pulumi.Input<string>;
    /**
     * The property describes Amazon Resource Name (ARN) of the service access role.
     */
    serviceAccessRoleArn: pulumi.Input<string>;
    /**
     * The property describes the settings for the data migration.
     */
    sourceDataSettings?: pulumi.Input<pulumi.Input<inputs.dms.DataMigrationSourceDataSettingsArgs>[]>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
