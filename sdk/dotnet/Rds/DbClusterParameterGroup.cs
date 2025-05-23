// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rds
{
    /// <summary>
    /// The ``AWS::RDS::DBClusterParameterGroup`` resource creates a new Amazon RDS DB cluster parameter group.
    ///  For information about configuring parameters for Amazon Aurora DB clusters, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.
    ///   If you apply a parameter group to a DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.
    ///  If you apply a change to parameter group associated with a stopped DB cluster, then the updated stack waits until the DB cluster is started.
    /// </summary>
    [AwsNativeResourceType("aws-native:rds:DbClusterParameterGroup")]
    public partial class DbClusterParameterGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the DB cluster parameter group.
        ///  Constraints:
        ///   +  Must not match the name of an existing DB cluster parameter group.
        ///   
        ///   This value is stored as a lowercase string.
        /// </summary>
        [Output("dbClusterParameterGroupName")]
        public Output<string?> DbClusterParameterGroupName { get; private set; } = null!;

        /// <summary>
        /// The description for the DB cluster parameter group.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.
        ///   *Aurora MySQL* 
        ///  Example: ``aurora-mysql5.7``, ``aurora-mysql8.0``
        ///   *Aurora PostgreSQL* 
        ///  Example: ``aurora-postgresql14``
        ///   *RDS for MySQL* 
        ///  Example: ``mysql8.0``
        ///   *RDS for PostgreSQL* 
        ///  Example: ``postgres13``
        ///  To list all of the available parameter group families for a DB engine, use the following command:
        ///   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine &lt;engine&gt;`` 
        ///  For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:
        ///   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine aurora-postgresql`` 
        ///   The output contains duplicates.
        ///   The following are the valid DB engine values:
        ///   +   ``aurora-mysql`` 
        ///   +   ``aurora-postgresql`` 
        ///   +   ``mysql`` 
        ///   +   ``postgres``
        /// </summary>
        [Output("family")]
        public Output<string> Family { get; private set; } = null!;

        /// <summary>
        /// Provides a list of parameters for the DB cluster parameter group.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::RDS::DBClusterParameterGroup` for more information about the expected schema for this property.
        /// </summary>
        [Output("parameters")]
        public Output<object> Parameters { get; private set; } = null!;

        /// <summary>
        /// Tags to assign to the DB cluster parameter group.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DbClusterParameterGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DbClusterParameterGroup(string name, DbClusterParameterGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:rds:DbClusterParameterGroup", name, args ?? new DbClusterParameterGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DbClusterParameterGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:rds:DbClusterParameterGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "dbClusterParameterGroupName",
                    "description",
                    "family",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DbClusterParameterGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DbClusterParameterGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DbClusterParameterGroup(name, id, options);
        }
    }

    public sealed class DbClusterParameterGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the DB cluster parameter group.
        ///  Constraints:
        ///   +  Must not match the name of an existing DB cluster parameter group.
        ///   
        ///   This value is stored as a lowercase string.
        /// </summary>
        [Input("dbClusterParameterGroupName")]
        public Input<string>? DbClusterParameterGroupName { get; set; }

        /// <summary>
        /// The description for the DB cluster parameter group.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.
        ///   *Aurora MySQL* 
        ///  Example: ``aurora-mysql5.7``, ``aurora-mysql8.0``
        ///   *Aurora PostgreSQL* 
        ///  Example: ``aurora-postgresql14``
        ///   *RDS for MySQL* 
        ///  Example: ``mysql8.0``
        ///   *RDS for PostgreSQL* 
        ///  Example: ``postgres13``
        ///  To list all of the available parameter group families for a DB engine, use the following command:
        ///   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine &lt;engine&gt;`` 
        ///  For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:
        ///   ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily" --engine aurora-postgresql`` 
        ///   The output contains duplicates.
        ///   The following are the valid DB engine values:
        ///   +   ``aurora-mysql`` 
        ///   +   ``aurora-postgresql`` 
        ///   +   ``mysql`` 
        ///   +   ``postgres``
        /// </summary>
        [Input("family", required: true)]
        public Input<string> Family { get; set; } = null!;

        /// <summary>
        /// Provides a list of parameters for the DB cluster parameter group.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::RDS::DBClusterParameterGroup` for more information about the expected schema for this property.
        /// </summary>
        [Input("parameters", required: true)]
        public Input<object> Parameters { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Tags to assign to the DB cluster parameter group.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public DbClusterParameterGroupArgs()
        {
        }
        public static new DbClusterParameterGroupArgs Empty => new DbClusterParameterGroupArgs();
    }
}
