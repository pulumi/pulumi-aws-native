// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A resource schema representing a Lake Formation Data Cells Filter.
 */
export class DataCellsFilter extends pulumi.CustomResource {
    /**
     * Get an existing DataCellsFilter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataCellsFilter {
        return new DataCellsFilter(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lakeformation:DataCellsFilter';

    /**
     * Returns true if the given object is an instance of DataCellsFilter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataCellsFilter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataCellsFilter.__pulumiType;
    }

    /**
     * A list of columns to be included in this Data Cells Filter.
     */
    public readonly columnNames!: pulumi.Output<string[] | undefined>;
    /**
     * An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required
     */
    public readonly columnWildcard!: pulumi.Output<outputs.lakeformation.DataCellsFilterColumnWildcard | undefined>;
    /**
     * The name of the Database that the Table resides in.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * The desired name of the Data Cells Filter.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required
     */
    public readonly rowFilter!: pulumi.Output<outputs.lakeformation.DataCellsFilterRowFilter | undefined>;
    /**
     * The Catalog Id of the Table on which to create a Data Cells Filter.
     */
    public readonly tableCatalogId!: pulumi.Output<string>;
    /**
     * The name of the Table to create a Data Cells Filter for.
     */
    public readonly tableName!: pulumi.Output<string>;

    /**
     * Create a DataCellsFilter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataCellsFilterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.tableCatalogId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableCatalogId'");
            }
            if ((!args || args.tableName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableName'");
            }
            resourceInputs["columnNames"] = args ? args.columnNames : undefined;
            resourceInputs["columnWildcard"] = args ? args.columnWildcard : undefined;
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rowFilter"] = args ? args.rowFilter : undefined;
            resourceInputs["tableCatalogId"] = args ? args.tableCatalogId : undefined;
            resourceInputs["tableName"] = args ? args.tableName : undefined;
        } else {
            resourceInputs["columnNames"] = undefined /*out*/;
            resourceInputs["columnWildcard"] = undefined /*out*/;
            resourceInputs["databaseName"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["rowFilter"] = undefined /*out*/;
            resourceInputs["tableCatalogId"] = undefined /*out*/;
            resourceInputs["tableName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["columnNames[*]", "columnWildcard", "databaseName", "name", "rowFilter", "tableCatalogId", "tableName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DataCellsFilter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataCellsFilter resource.
 */
export interface DataCellsFilterArgs {
    /**
     * A list of columns to be included in this Data Cells Filter.
     */
    columnNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required
     */
    columnWildcard?: pulumi.Input<inputs.lakeformation.DataCellsFilterColumnWildcardArgs>;
    /**
     * The name of the Database that the Table resides in.
     */
    databaseName: pulumi.Input<string>;
    /**
     * The desired name of the Data Cells Filter.
     */
    name?: pulumi.Input<string>;
    /**
     * An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required
     */
    rowFilter?: pulumi.Input<inputs.lakeformation.DataCellsFilterRowFilterArgs>;
    /**
     * The Catalog Id of the Table on which to create a Data Cells Filter.
     */
    tableCatalogId: pulumi.Input<string>;
    /**
     * The name of the Table to create a Data Cells Filter for.
     */
    tableName: pulumi.Input<string>;
}
