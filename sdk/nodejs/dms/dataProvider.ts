// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DMS::DataProvider
 */
export class DataProvider extends pulumi.CustomResource {
    /**
     * Get an existing DataProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataProvider {
        return new DataProvider(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:dms:DataProvider';

    /**
     * Returns true if the given object is an instance of DataProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataProvider.__pulumiType;
    }

    /**
     * The data provider ARN.
     */
    public /*out*/ readonly dataProviderArn!: pulumi.Output<string>;
    /**
     * The data provider creation time.
     */
    public /*out*/ readonly dataProviderCreationTime!: pulumi.Output<string>;
    /**
     * The property describes an identifier for the data provider. It is used for describing/deleting/modifying can be name/arn
     */
    public readonly dataProviderIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The property describes a name to identify the data provider.
     */
    public readonly dataProviderName!: pulumi.Output<string | undefined>;
    /**
     * The optional description of the data provider.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The property describes a data engine for the data provider.
     */
    public readonly engine!: pulumi.Output<enums.dms.DataProviderEngine>;
    /**
     * The property describes the exact settings which can be modified
     */
    public readonly exactSettings!: pulumi.Output<boolean | undefined>;
    /**
     * The property identifies the exact type of settings for the data provider.
     */
    public readonly settings!: pulumi.Output<outputs.dms.Settings0Properties | outputs.dms.Settings1Properties | outputs.dms.Settings2Properties | outputs.dms.Settings3Properties | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.dms.DataProviderTag[] | undefined>;

    /**
     * Create a DataProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataProviderArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            resourceInputs["dataProviderIdentifier"] = args ? args.dataProviderIdentifier : undefined;
            resourceInputs["dataProviderName"] = args ? args.dataProviderName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["exactSettings"] = args ? args.exactSettings : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["dataProviderArn"] = undefined /*out*/;
            resourceInputs["dataProviderCreationTime"] = undefined /*out*/;
        } else {
            resourceInputs["dataProviderArn"] = undefined /*out*/;
            resourceInputs["dataProviderCreationTime"] = undefined /*out*/;
            resourceInputs["dataProviderIdentifier"] = undefined /*out*/;
            resourceInputs["dataProviderName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["engine"] = undefined /*out*/;
            resourceInputs["exactSettings"] = undefined /*out*/;
            resourceInputs["settings"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DataProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataProvider resource.
 */
export interface DataProviderArgs {
    /**
     * The property describes an identifier for the data provider. It is used for describing/deleting/modifying can be name/arn
     */
    dataProviderIdentifier?: pulumi.Input<string>;
    /**
     * The property describes a name to identify the data provider.
     */
    dataProviderName?: pulumi.Input<string>;
    /**
     * The optional description of the data provider.
     */
    description?: pulumi.Input<string>;
    /**
     * The property describes a data engine for the data provider.
     */
    engine: pulumi.Input<enums.dms.DataProviderEngine>;
    /**
     * The property describes the exact settings which can be modified
     */
    exactSettings?: pulumi.Input<boolean>;
    /**
     * The property identifies the exact type of settings for the data provider.
     */
    settings?: pulumi.Input<inputs.dms.Settings0PropertiesArgs | inputs.dms.Settings1PropertiesArgs | inputs.dms.Settings2PropertiesArgs | inputs.dms.Settings3PropertiesArgs>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.dms.DataProviderTagArgs>[]>;
}