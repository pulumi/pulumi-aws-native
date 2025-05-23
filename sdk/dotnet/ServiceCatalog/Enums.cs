// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.ServiceCatalog
{
    /// <summary>
    /// The language code.
    /// 
    /// - `jp` - Japanese
    /// - `zh` - Chinese
    /// </summary>
    [EnumType]
    public readonly struct CloudFormationProvisionedProductAcceptLanguage : IEquatable<CloudFormationProvisionedProductAcceptLanguage>
    {
        private readonly string _value;

        private CloudFormationProvisionedProductAcceptLanguage(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CloudFormationProvisionedProductAcceptLanguage En { get; } = new CloudFormationProvisionedProductAcceptLanguage("en");
        public static CloudFormationProvisionedProductAcceptLanguage Jp { get; } = new CloudFormationProvisionedProductAcceptLanguage("jp");
        public static CloudFormationProvisionedProductAcceptLanguage Zh { get; } = new CloudFormationProvisionedProductAcceptLanguage("zh");

        public static bool operator ==(CloudFormationProvisionedProductAcceptLanguage left, CloudFormationProvisionedProductAcceptLanguage right) => left.Equals(right);
        public static bool operator !=(CloudFormationProvisionedProductAcceptLanguage left, CloudFormationProvisionedProductAcceptLanguage right) => !left.Equals(right);

        public static explicit operator string(CloudFormationProvisionedProductAcceptLanguage value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CloudFormationProvisionedProductAcceptLanguage other && Equals(other);
        public bool Equals(CloudFormationProvisionedProductAcceptLanguage other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Determines what action AWS Service Catalog performs to a stack set or a stack instance represented by the provisioned product. The default value is `UPDATE` if nothing is specified.
    /// 
    /// Applicable only to a `CFN_STACKSET` provisioned product type.
    /// 
    /// - **CREATE** - Creates a new stack instance in the stack set represented by the provisioned product. In this case, only new stack instances are created based on accounts and Regions; if new ProductId or ProvisioningArtifactID are passed, they will be ignored.
    /// - **UPDATE** - Updates the stack set represented by the provisioned product and also its stack instances.
    /// - **DELETE** - Deletes a stack instance in the stack set represented by the provisioned product.
    /// </summary>
    [EnumType]
    public readonly struct CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType : IEquatable<CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType>
    {
        private readonly string _value;

        private CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType Create { get; } = new CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType("CREATE");
        public static CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType Update { get; } = new CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType("UPDATE");
        public static CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType Delete { get; } = new CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType("DELETE");

        public static bool operator ==(CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType left, CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType right) => left.Equals(right);
        public static bool operator !=(CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType left, CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType right) => !left.Equals(right);

        public static explicit operator string(CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType other && Equals(other);
        public bool Equals(CloudFormationProvisionedProductProvisioningPreferencesStackSetOperationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The language code.
    /// 
    /// - `en` - English (default)
    /// - `jp` - Japanese
    /// - `zh` - Chinese
    /// </summary>
    [EnumType]
    public readonly struct ServiceActionAcceptLanguage : IEquatable<ServiceActionAcceptLanguage>
    {
        private readonly string _value;

        private ServiceActionAcceptLanguage(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceActionAcceptLanguage En { get; } = new ServiceActionAcceptLanguage("en");
        public static ServiceActionAcceptLanguage Jp { get; } = new ServiceActionAcceptLanguage("jp");
        public static ServiceActionAcceptLanguage Zh { get; } = new ServiceActionAcceptLanguage("zh");

        public static bool operator ==(ServiceActionAcceptLanguage left, ServiceActionAcceptLanguage right) => left.Equals(right);
        public static bool operator !=(ServiceActionAcceptLanguage left, ServiceActionAcceptLanguage right) => !left.Equals(right);

        public static explicit operator string(ServiceActionAcceptLanguage value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceActionAcceptLanguage other && Equals(other);
        public bool Equals(ServiceActionAcceptLanguage other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The self-service action definition type. For example, `SSM_AUTOMATION` .
    /// </summary>
    [EnumType]
    public readonly struct ServiceActionDefinitionType : IEquatable<ServiceActionDefinitionType>
    {
        private readonly string _value;

        private ServiceActionDefinitionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ServiceActionDefinitionType SsmAutomation { get; } = new ServiceActionDefinitionType("SSM_AUTOMATION");

        public static bool operator ==(ServiceActionDefinitionType left, ServiceActionDefinitionType right) => left.Equals(right);
        public static bool operator !=(ServiceActionDefinitionType left, ServiceActionDefinitionType right) => !left.Equals(right);

        public static explicit operator string(ServiceActionDefinitionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ServiceActionDefinitionType other && Equals(other);
        public bool Equals(ServiceActionDefinitionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
