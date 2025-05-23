// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cassandra
{
    /// <summary>
    /// Resource schema for AWS::Cassandra::Keyspace
    /// 
    /// ## Example Usage
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myNewKeyspace = new AwsNative.Cassandra.Keyspace("myNewKeyspace", new()
    ///     {
    ///         KeyspaceName = "MyNewKeyspace",
    ///         Tags = new[]
    ///         {
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "tag1",
    ///                 Value = "val1",
    ///             },
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "tag2",
    ///                 Value = "val2",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var multiRegionKeyspace = new AwsNative.Cassandra.Keyspace("multiRegionKeyspace", new()
    ///     {
    ///         KeyspaceName = "MultiRegionKeyspace",
    ///         ReplicationSpecification = new AwsNative.Cassandra.Inputs.KeyspaceReplicationSpecificationArgs
    ///         {
    ///             ReplicationStrategy = AwsNative.Cassandra.KeyspaceReplicationSpecificationReplicationStrategy.MultiRegion,
    ///             RegionList = new[]
    ///             {
    ///                 AwsNative.Cassandra.KeyspaceRegionListItem.UsEast1,
    ///                 AwsNative.Cassandra.KeyspaceRegionListItem.UsWest2,
    ///                 AwsNative.Cassandra.KeyspaceRegionListItem.EuWest1,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var multiRegionKeyspace = new AwsNative.Cassandra.Keyspace("multiRegionKeyspace", new()
    ///     {
    ///         KeyspaceName = "MultiRegionKeyspace",
    ///         ReplicationSpecification = new AwsNative.Cassandra.Inputs.KeyspaceReplicationSpecificationArgs
    ///         {
    ///             ReplicationStrategy = AwsNative.Cassandra.KeyspaceReplicationSpecificationReplicationStrategy.MultiRegion,
    ///             RegionList = new[]
    ///             {
    ///                 AwsNative.Cassandra.KeyspaceRegionListItem.UsEast1,
    ///                 AwsNative.Cassandra.KeyspaceRegionListItem.UsWest2,
    ///                 AwsNative.Cassandra.KeyspaceRegionListItem.EuWest1,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:cassandra:Keyspace")]
    public partial class Keyspace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether client-side timestamps are enabled (true) or disabled (false) for all tables in the keyspace. To add a Region to a single-Region keyspace with at least one table, the value must be set to true. After you enabled client-side timestamps for a table, you can’t disable it again.
        /// </summary>
        [Output("clientSideTimestampsEnabled")]
        public Output<bool?> ClientSideTimestampsEnabled { get; private set; } = null!;

        /// <summary>
        /// Name for Cassandra keyspace
        /// </summary>
        [Output("keyspaceName")]
        public Output<string?> KeyspaceName { get; private set; } = null!;

        /// <summary>
        /// Specifies the `ReplicationStrategy` of a keyspace. The options are:
        /// 
        /// - `SINGLE_REGION` for a single Region keyspace (optional) or
        /// - `MULTI_REGION` for a multi-Region keyspace
        /// 
        /// If no `ReplicationStrategy` is provided, the default is `SINGLE_REGION` . If you choose `MULTI_REGION` , you must also provide a `RegionList` with the AWS Regions that the keyspace is replicated in.
        /// </summary>
        [Output("replicationSpecification")]
        public Output<Outputs.KeyspaceReplicationSpecification?> ReplicationSpecification { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// 
        /// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Keyspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Keyspace(string name, KeyspaceArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:cassandra:Keyspace", name, args ?? new KeyspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Keyspace(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cassandra:Keyspace", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "keyspaceName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Keyspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Keyspace Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Keyspace(name, id, options);
        }
    }

    public sealed class KeyspaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether client-side timestamps are enabled (true) or disabled (false) for all tables in the keyspace. To add a Region to a single-Region keyspace with at least one table, the value must be set to true. After you enabled client-side timestamps for a table, you can’t disable it again.
        /// </summary>
        [Input("clientSideTimestampsEnabled")]
        public Input<bool>? ClientSideTimestampsEnabled { get; set; }

        /// <summary>
        /// Name for Cassandra keyspace
        /// </summary>
        [Input("keyspaceName")]
        public Input<string>? KeyspaceName { get; set; }

        /// <summary>
        /// Specifies the `ReplicationStrategy` of a keyspace. The options are:
        /// 
        /// - `SINGLE_REGION` for a single Region keyspace (optional) or
        /// - `MULTI_REGION` for a multi-Region keyspace
        /// 
        /// If no `ReplicationStrategy` is provided, the default is `SINGLE_REGION` . If you choose `MULTI_REGION` , you must also provide a `RegionList` with the AWS Regions that the keyspace is replicated in.
        /// </summary>
        [Input("replicationSpecification")]
        public Input<Inputs.KeyspaceReplicationSpecificationArgs>? ReplicationSpecification { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// 
        /// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public KeyspaceArgs()
        {
        }
        public static new KeyspaceArgs Empty => new KeyspaceArgs();
    }
}
