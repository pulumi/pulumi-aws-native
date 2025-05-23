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
    /// A zero-ETL integration with Amazon Redshift.
    /// </summary>
    [AwsNativeResourceType("aws-native:rds:Integration")]
    public partial class Integration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An optional set of non-secret key–value pairs that contains additional contextual information about the data. For more information, see [Encryption context](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context) in the *Key Management Service Developer Guide*.
        ///  You can only include this parameter if you specify the ``KMSKeyId`` parameter.
        /// </summary>
        [Output("additionalEncryptionContext")]
        public Output<ImmutableDictionary<string, string>?> AdditionalEncryptionContext { get; private set; } = null!;

        /// <summary>
        /// The time when the integration was created, in Universal Coordinated Time (UTC).
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Data filters for the integration. These filters determine which tables from the source database are sent to the target Amazon Redshift data warehouse.
        /// </summary>
        [Output("dataFilter")]
        public Output<string?> DataFilter { get; private set; } = null!;

        /// <summary>
        /// A description of the integration.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ARN of the integration.
        /// </summary>
        [Output("integrationArn")]
        public Output<string> IntegrationArn { get; private set; } = null!;

        /// <summary>
        /// The name of the integration.
        /// </summary>
        [Output("integrationName")]
        public Output<string?> IntegrationName { get; private set; } = null!;

        /// <summary>
        /// The AWS Key Management System (AWS KMS) key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, RDS uses a default AWS owned key.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the database to use as the source for replication.
        /// </summary>
        [Output("sourceArn")]
        public Output<string> SourceArn { get; private set; } = null!;

        /// <summary>
        /// A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The ARN of the Redshift data warehouse to use as the target for replication.
        /// </summary>
        [Output("targetArn")]
        public Output<string> TargetArn { get; private set; } = null!;


        /// <summary>
        /// Create a Integration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integration(string name, IntegrationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:rds:Integration", name, args ?? new IntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Integration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:rds:Integration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "additionalEncryptionContext.*",
                    "kmsKeyId",
                    "sourceArn",
                    "targetArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Integration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Integration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Integration(name, id, options);
        }
    }

    public sealed class IntegrationArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalEncryptionContext")]
        private InputMap<string>? _additionalEncryptionContext;

        /// <summary>
        /// An optional set of non-secret key–value pairs that contains additional contextual information about the data. For more information, see [Encryption context](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context) in the *Key Management Service Developer Guide*.
        ///  You can only include this parameter if you specify the ``KMSKeyId`` parameter.
        /// </summary>
        public InputMap<string> AdditionalEncryptionContext
        {
            get => _additionalEncryptionContext ?? (_additionalEncryptionContext = new InputMap<string>());
            set => _additionalEncryptionContext = value;
        }

        /// <summary>
        /// Data filters for the integration. These filters determine which tables from the source database are sent to the target Amazon Redshift data warehouse.
        /// </summary>
        [Input("dataFilter")]
        public Input<string>? DataFilter { get; set; }

        /// <summary>
        /// A description of the integration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the integration.
        /// </summary>
        [Input("integrationName")]
        public Input<string>? IntegrationName { get; set; }

        /// <summary>
        /// The AWS Key Management System (AWS KMS) key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, RDS uses a default AWS owned key.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the database to use as the source for replication.
        /// </summary>
        [Input("sourceArn", required: true)]
        public Input<string> SourceArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ARN of the Redshift data warehouse to use as the target for replication.
        /// </summary>
        [Input("targetArn", required: true)]
        public Input<string> TargetArn { get; set; } = null!;

        public IntegrationArgs()
        {
        }
        public static new IntegrationArgs Empty => new IntegrationArgs();
    }
}
