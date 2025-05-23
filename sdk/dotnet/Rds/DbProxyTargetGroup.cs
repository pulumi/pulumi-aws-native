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
    /// Resource schema for AWS::RDS::DBProxyTargetGroup
    /// </summary>
    [AwsNativeResourceType("aws-native:rds:DbProxyTargetGroup")]
    public partial class DbProxyTargetGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Displays the settings that control the size and behavior of the connection pool associated with a `DBProxyTarget` .
        /// </summary>
        [Output("connectionPoolConfigurationInfo")]
        public Output<Outputs.DbProxyTargetGroupConnectionPoolConfigurationInfoFormat?> ConnectionPoolConfigurationInfo { get; private set; } = null!;

        /// <summary>
        /// One or more DB cluster identifiers.
        /// </summary>
        [Output("dbClusterIdentifiers")]
        public Output<ImmutableArray<string>> DbClusterIdentifiers { get; private set; } = null!;

        /// <summary>
        /// One or more DB instance identifiers.
        /// </summary>
        [Output("dbInstanceIdentifiers")]
        public Output<ImmutableArray<string>> DbInstanceIdentifiers { get; private set; } = null!;

        /// <summary>
        /// The identifier for the proxy.
        /// </summary>
        [Output("dbProxyName")]
        public Output<string> DbProxyName { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) representing the target group.
        /// </summary>
        [Output("targetGroupArn")]
        public Output<string> TargetGroupArn { get; private set; } = null!;

        /// <summary>
        /// The identifier for the DBProxyTargetGroup
        /// </summary>
        [Output("targetGroupName")]
        public Output<Pulumi.AwsNative.Rds.DbProxyTargetGroupTargetGroupName> TargetGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a DbProxyTargetGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DbProxyTargetGroup(string name, DbProxyTargetGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:rds:DbProxyTargetGroup", name, args ?? new DbProxyTargetGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DbProxyTargetGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:rds:DbProxyTargetGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "dbProxyName",
                    "targetGroupName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DbProxyTargetGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DbProxyTargetGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DbProxyTargetGroup(name, id, options);
        }
    }

    public sealed class DbProxyTargetGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Displays the settings that control the size and behavior of the connection pool associated with a `DBProxyTarget` .
        /// </summary>
        [Input("connectionPoolConfigurationInfo")]
        public Input<Inputs.DbProxyTargetGroupConnectionPoolConfigurationInfoFormatArgs>? ConnectionPoolConfigurationInfo { get; set; }

        [Input("dbClusterIdentifiers")]
        private InputList<string>? _dbClusterIdentifiers;

        /// <summary>
        /// One or more DB cluster identifiers.
        /// </summary>
        public InputList<string> DbClusterIdentifiers
        {
            get => _dbClusterIdentifiers ?? (_dbClusterIdentifiers = new InputList<string>());
            set => _dbClusterIdentifiers = value;
        }

        [Input("dbInstanceIdentifiers")]
        private InputList<string>? _dbInstanceIdentifiers;

        /// <summary>
        /// One or more DB instance identifiers.
        /// </summary>
        public InputList<string> DbInstanceIdentifiers
        {
            get => _dbInstanceIdentifiers ?? (_dbInstanceIdentifiers = new InputList<string>());
            set => _dbInstanceIdentifiers = value;
        }

        /// <summary>
        /// The identifier for the proxy.
        /// </summary>
        [Input("dbProxyName", required: true)]
        public Input<string> DbProxyName { get; set; } = null!;

        /// <summary>
        /// The identifier for the DBProxyTargetGroup
        /// </summary>
        [Input("targetGroupName", required: true)]
        public Input<Pulumi.AwsNative.Rds.DbProxyTargetGroupTargetGroupName> TargetGroupName { get; set; } = null!;

        public DbProxyTargetGroupArgs()
        {
        }
        public static new DbProxyTargetGroupArgs Empty => new DbProxyTargetGroupArgs();
    }
}
