// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.S3Tables
{
    /// <summary>
    /// Server-side encryption algorithm
    /// </summary>
    [EnumType]
    public readonly struct TableBucketEncryptionConfigurationSseAlgorithm : IEquatable<TableBucketEncryptionConfigurationSseAlgorithm>
    {
        private readonly string _value;

        private TableBucketEncryptionConfigurationSseAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TableBucketEncryptionConfigurationSseAlgorithm Aes256 { get; } = new TableBucketEncryptionConfigurationSseAlgorithm("AES256");
        public static TableBucketEncryptionConfigurationSseAlgorithm Awskms { get; } = new TableBucketEncryptionConfigurationSseAlgorithm("aws:kms");

        public static bool operator ==(TableBucketEncryptionConfigurationSseAlgorithm left, TableBucketEncryptionConfigurationSseAlgorithm right) => left.Equals(right);
        public static bool operator !=(TableBucketEncryptionConfigurationSseAlgorithm left, TableBucketEncryptionConfigurationSseAlgorithm right) => !left.Equals(right);

        public static explicit operator string(TableBucketEncryptionConfigurationSseAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TableBucketEncryptionConfigurationSseAlgorithm other && Equals(other);
        public bool Equals(TableBucketEncryptionConfigurationSseAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Indicates whether the Unreferenced File Removal maintenance action is enabled.
    /// </summary>
    [EnumType]
    public readonly struct TableBucketUnreferencedFileRemovalStatus : IEquatable<TableBucketUnreferencedFileRemovalStatus>
    {
        private readonly string _value;

        private TableBucketUnreferencedFileRemovalStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TableBucketUnreferencedFileRemovalStatus Enabled { get; } = new TableBucketUnreferencedFileRemovalStatus("Enabled");
        public static TableBucketUnreferencedFileRemovalStatus Disabled { get; } = new TableBucketUnreferencedFileRemovalStatus("Disabled");

        public static bool operator ==(TableBucketUnreferencedFileRemovalStatus left, TableBucketUnreferencedFileRemovalStatus right) => left.Equals(right);
        public static bool operator !=(TableBucketUnreferencedFileRemovalStatus left, TableBucketUnreferencedFileRemovalStatus right) => !left.Equals(right);

        public static explicit operator string(TableBucketUnreferencedFileRemovalStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TableBucketUnreferencedFileRemovalStatus other && Equals(other);
        public bool Equals(TableBucketUnreferencedFileRemovalStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
