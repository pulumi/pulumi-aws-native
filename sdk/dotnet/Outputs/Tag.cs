// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Outputs
{

    /// <summary>
    /// A set of tags to apply to the resource.
    /// </summary>
    [OutputType]
    public sealed class Tag
    {
        /// <summary>
        /// The key name of the tag
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The value of the tag
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private Tag(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}