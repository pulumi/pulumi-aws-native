// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AwsNative.ApiGatewayV2
{
    /// <summary>
    /// The routing mode API Gateway uses to route traffic to your APIs.
    /// </summary>
    [EnumType]
    public readonly struct DomainNameRoutingMode : IEquatable<DomainNameRoutingMode>
    {
        private readonly string _value;

        private DomainNameRoutingMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DomainNameRoutingMode ApiMappingOnly { get; } = new DomainNameRoutingMode("API_MAPPING_ONLY");
        public static DomainNameRoutingMode RoutingRuleThenApiMapping { get; } = new DomainNameRoutingMode("ROUTING_RULE_THEN_API_MAPPING");
        public static DomainNameRoutingMode RoutingRuleOnly { get; } = new DomainNameRoutingMode("ROUTING_RULE_ONLY");

        public static bool operator ==(DomainNameRoutingMode left, DomainNameRoutingMode right) => left.Equals(right);
        public static bool operator !=(DomainNameRoutingMode left, DomainNameRoutingMode right) => !left.Equals(right);

        public static explicit operator string(DomainNameRoutingMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DomainNameRoutingMode other && Equals(other);
        public bool Equals(DomainNameRoutingMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
