// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Location
{
    /// <summary>
    /// Definition of AWS::Location::GeofenceCollection Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:location:GeofenceCollection")]
    public partial class GeofenceCollection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the geofence collection resource. Used when you need to specify a resource across all AWS .
        /// 
        /// - Format example: `arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollection`
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Synonym for `Arn` . The Amazon Resource Name (ARN) for the geofence collection resource. Used when you need to specify a resource across all AWS .
        /// 
        /// - Format example: `arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollection`
        /// </summary>
        [Output("collectionArn")]
        public Output<string> CollectionArn { get; private set; } = null!;

        /// <summary>
        /// A custom name for the geofence collection.
        /// 
        /// Requirements:
        /// 
        /// - Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
        /// - Must be a unique geofence collection name.
        /// - No spaces allowed. For example, `ExampleGeofenceCollection` .
        /// </summary>
        [Output("collectionName")]
        public Output<string> CollectionName { get; private set; } = null!;

        /// <summary>
        /// The timestamp for when the geofence collection resource was created in [ISO 8601](https://docs.aws.amazon.com/https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ` .
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// An optional description for the geofence collection.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A key identifier for an [AWS KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html) . Enter a key ID, key ARN, alias name, or alias ARN.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        [Output("pricingPlan")]
        public Output<Pulumi.AwsNative.Location.GeofenceCollectionPricingPlan?> PricingPlan { get; private set; } = null!;

        /// <summary>
        /// This shape is deprecated since 2022-02-01: Deprecated. No longer allowed.
        /// </summary>
        [Output("pricingPlanDataSource")]
        public Output<string?> PricingPlanDataSource { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The timestamp for when the geofence collection resource was last updated in [ISO 8601](https://docs.aws.amazon.com/https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ` .
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a GeofenceCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GeofenceCollection(string name, GeofenceCollectionArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:location:GeofenceCollection", name, args ?? new GeofenceCollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GeofenceCollection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:location:GeofenceCollection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "collectionName",
                    "kmsKeyId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GeofenceCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GeofenceCollection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GeofenceCollection(name, id, options);
        }
    }

    public sealed class GeofenceCollectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A custom name for the geofence collection.
        /// 
        /// Requirements:
        /// 
        /// - Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
        /// - Must be a unique geofence collection name.
        /// - No spaces allowed. For example, `ExampleGeofenceCollection` .
        /// </summary>
        [Input("collectionName")]
        public Input<string>? CollectionName { get; set; }

        /// <summary>
        /// An optional description for the geofence collection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A key identifier for an [AWS KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html) . Enter a key ID, key ARN, alias name, or alias ARN.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("pricingPlan")]
        public Input<Pulumi.AwsNative.Location.GeofenceCollectionPricingPlan>? PricingPlan { get; set; }

        /// <summary>
        /// This shape is deprecated since 2022-02-01: Deprecated. No longer allowed.
        /// </summary>
        [Input("pricingPlanDataSource")]
        public Input<string>? PricingPlanDataSource { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public GeofenceCollectionArgs()
        {
        }
        public static new GeofenceCollectionArgs Empty => new GeofenceCollectionArgs();
    }
}
