// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    [OutputType]
    public sealed class UsageProfileConfigurationObject
    {
        public readonly ImmutableArray<string> AllowedValues;
        public readonly string? DefaultValue;
        public readonly string? MaxValue;
        public readonly string? MinValue;

        [OutputConstructor]
        private UsageProfileConfigurationObject(
            ImmutableArray<string> allowedValues,

            string? defaultValue,

            string? maxValue,

            string? minValue)
        {
            AllowedValues = allowedValues;
            DefaultValue = defaultValue;
            MaxValue = maxValue;
            MinValue = minValue;
        }
    }
}