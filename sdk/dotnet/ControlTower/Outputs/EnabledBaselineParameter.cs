// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ControlTower.Outputs
{

    [OutputType]
    public sealed class EnabledBaselineParameter
    {
        /// <summary>
        /// A string denoting the parameter key.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// A low-level `Document` object of any type (for example, a Java Object).
        /// </summary>
        public readonly object? Value;

        [OutputConstructor]
        private EnabledBaselineParameter(
            string? key,

            object? value)
        {
            Key = key;
            Value = value;
        }
    }
}