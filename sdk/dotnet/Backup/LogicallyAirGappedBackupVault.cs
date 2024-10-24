// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Backup
{
    /// <summary>
    /// Resource Type definition for AWS::Backup::LogicallyAirGappedBackupVault
    /// </summary>
    [AwsNativeResourceType("aws-native:backup:LogicallyAirGappedBackupVault")]
    public partial class LogicallyAirGappedBackupVault : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Backup::LogicallyAirGappedBackupVault` for more information about the expected schema for this property.
        /// </summary>
        [Output("accessPolicy")]
        public Output<object?> AccessPolicy { get; private set; } = null!;

        [Output("backupVaultArn")]
        public Output<string> BackupVaultArn { get; private set; } = null!;

        [Output("backupVaultName")]
        public Output<string> BackupVaultName { get; private set; } = null!;

        [Output("backupVaultTags")]
        public Output<ImmutableDictionary<string, string>?> BackupVaultTags { get; private set; } = null!;

        [Output("encryptionKeyArn")]
        public Output<string> EncryptionKeyArn { get; private set; } = null!;

        [Output("maxRetentionDays")]
        public Output<int> MaxRetentionDays { get; private set; } = null!;

        [Output("minRetentionDays")]
        public Output<int> MinRetentionDays { get; private set; } = null!;

        [Output("notifications")]
        public Output<Outputs.LogicallyAirGappedBackupVaultNotificationObjectType?> Notifications { get; private set; } = null!;

        [Output("vaultState")]
        public Output<string?> VaultState { get; private set; } = null!;

        [Output("vaultType")]
        public Output<string?> VaultType { get; private set; } = null!;


        /// <summary>
        /// Create a LogicallyAirGappedBackupVault resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogicallyAirGappedBackupVault(string name, LogicallyAirGappedBackupVaultArgs args, CustomResourceOptions? options = null)
            : base("aws-native:backup:LogicallyAirGappedBackupVault", name, args ?? new LogicallyAirGappedBackupVaultArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogicallyAirGappedBackupVault(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:backup:LogicallyAirGappedBackupVault", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "backupVaultName",
                    "maxRetentionDays",
                    "minRetentionDays",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LogicallyAirGappedBackupVault resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogicallyAirGappedBackupVault Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LogicallyAirGappedBackupVault(name, id, options);
        }
    }

    public sealed class LogicallyAirGappedBackupVaultArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Backup::LogicallyAirGappedBackupVault` for more information about the expected schema for this property.
        /// </summary>
        [Input("accessPolicy")]
        public Input<object>? AccessPolicy { get; set; }

        [Input("backupVaultName")]
        public Input<string>? BackupVaultName { get; set; }

        [Input("backupVaultTags")]
        private InputMap<string>? _backupVaultTags;
        public InputMap<string> BackupVaultTags
        {
            get => _backupVaultTags ?? (_backupVaultTags = new InputMap<string>());
            set => _backupVaultTags = value;
        }

        [Input("maxRetentionDays", required: true)]
        public Input<int> MaxRetentionDays { get; set; } = null!;

        [Input("minRetentionDays", required: true)]
        public Input<int> MinRetentionDays { get; set; } = null!;

        [Input("notifications")]
        public Input<Inputs.LogicallyAirGappedBackupVaultNotificationObjectTypeArgs>? Notifications { get; set; }

        [Input("vaultState")]
        public Input<string>? VaultState { get; set; }

        [Input("vaultType")]
        public Input<string>? VaultType { get; set; }

        public LogicallyAirGappedBackupVaultArgs()
        {
        }
        public static new LogicallyAirGappedBackupVaultArgs Empty => new LogicallyAirGappedBackupVaultArgs();
    }
}
