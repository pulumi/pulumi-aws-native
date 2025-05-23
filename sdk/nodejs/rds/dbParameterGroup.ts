// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ``AWS::RDS::DBParameterGroup`` resource creates a custom parameter group for an RDS database family.
 *  This type can be declared in a template and referenced in the ``DBParameterGroupName`` property of an ``AWS::RDS::DBInstance`` resource.
 *  For information about configuring parameters for Amazon RDS DB instances, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide*.
 *  For information about configuring parameters for Amazon Aurora DB instances, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.
 *   Applying a parameter group to a DB instance may require the DB instance to reboot, resulting in a database outage for the duration of the reboot.
 */
export class DbParameterGroup extends pulumi.CustomResource {
    /**
     * Get an existing DbParameterGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DbParameterGroup {
        return new DbParameterGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:rds:DbParameterGroup';

    /**
     * Returns true if the given object is an instance of DbParameterGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DbParameterGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DbParameterGroup.__pulumiType;
    }

    /**
     * The name of the DB parameter group.
     *  Constraints:
     *   +  Must be 1 to 255 letters, numbers, or hyphens.
     *   +  First character must be a letter
     *   +  Can't end with a hyphen or contain two consecutive hyphens
     *   
     *  If you don't specify a value for ``DBParameterGroupName`` property, a name is automatically created for the DB parameter group.
     *   This value is stored as a lowercase string.
     */
    public readonly dbParameterGroupName!: pulumi.Output<string | undefined>;
    /**
     * Provides the customer-specified description for this DB parameter group.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The DB parameter group family name. A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a database engine and engine version compatible with that DB parameter group family.
     *  To list all of the available parameter group families for a DB engine, use the following command:
     *   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine <engine>`` 
     *  For example, to list all of the available parameter group families for the MySQL DB engine, use the following command:
     *   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine mysql`` 
     *   The output contains duplicates.
     *   The following are the valid DB engine values:
     *   +   ``aurora-mysql`` 
     *   +   ``aurora-postgresql`` 
     *   +   ``db2-ae`` 
     *   +   ``db2-se`` 
     *   +   ``mysql`` 
     *   +   ``oracle-ee`` 
     *   +   ``oracle-ee-cdb`` 
     *   +   ``oracle-se2`` 
     *   +   ``oracle-se2-cdb`` 
     *   +   ``postgres`` 
     *   +   ``sqlserver-ee`` 
     *   +   ``sqlserver-se`` 
     *   +   ``sqlserver-ex`` 
     *   +   ``sqlserver-web``
     */
    public readonly family!: pulumi.Output<string>;
    /**
     * A mapping of parameter names and values for the parameter update. You must specify at least one parameter name and value.
     *  For more information about parameter groups, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide*, or [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.
     *   AWS CloudFormation doesn't support specifying an apply method for each individual parameter. The default apply method for each parameter is used.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::RDS::DBParameterGroup` for more information about the expected schema for this property.
     */
    public readonly parameters!: pulumi.Output<any | undefined>;
    /**
     * Tags to assign to the DB parameter group.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a DbParameterGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DbParameterGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.family === undefined) && !opts.urn) {
                throw new Error("Missing required property 'family'");
            }
            resourceInputs["dbParameterGroupName"] = args ? args.dbParameterGroupName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["family"] = args ? args.family : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        } else {
            resourceInputs["dbParameterGroupName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["family"] = undefined /*out*/;
            resourceInputs["parameters"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["dbParameterGroupName", "description", "family"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DbParameterGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DbParameterGroup resource.
 */
export interface DbParameterGroupArgs {
    /**
     * The name of the DB parameter group.
     *  Constraints:
     *   +  Must be 1 to 255 letters, numbers, or hyphens.
     *   +  First character must be a letter
     *   +  Can't end with a hyphen or contain two consecutive hyphens
     *   
     *  If you don't specify a value for ``DBParameterGroupName`` property, a name is automatically created for the DB parameter group.
     *   This value is stored as a lowercase string.
     */
    dbParameterGroupName?: pulumi.Input<string>;
    /**
     * Provides the customer-specified description for this DB parameter group.
     */
    description: pulumi.Input<string>;
    /**
     * The DB parameter group family name. A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a database engine and engine version compatible with that DB parameter group family.
     *  To list all of the available parameter group families for a DB engine, use the following command:
     *   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine <engine>`` 
     *  For example, to list all of the available parameter group families for the MySQL DB engine, use the following command:
     *   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine mysql`` 
     *   The output contains duplicates.
     *   The following are the valid DB engine values:
     *   +   ``aurora-mysql`` 
     *   +   ``aurora-postgresql`` 
     *   +   ``db2-ae`` 
     *   +   ``db2-se`` 
     *   +   ``mysql`` 
     *   +   ``oracle-ee`` 
     *   +   ``oracle-ee-cdb`` 
     *   +   ``oracle-se2`` 
     *   +   ``oracle-se2-cdb`` 
     *   +   ``postgres`` 
     *   +   ``sqlserver-ee`` 
     *   +   ``sqlserver-se`` 
     *   +   ``sqlserver-ex`` 
     *   +   ``sqlserver-web``
     */
    family: pulumi.Input<string>;
    /**
     * A mapping of parameter names and values for the parameter update. You must specify at least one parameter name and value.
     *  For more information about parameter groups, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html) in the *Amazon RDS User Guide*, or [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.
     *   AWS CloudFormation doesn't support specifying an apply method for each individual parameter. The default apply method for each parameter is used.
     *
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::RDS::DBParameterGroup` for more information about the expected schema for this property.
     */
    parameters?: any;
    /**
     * Tags to assign to the DB parameter group.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
