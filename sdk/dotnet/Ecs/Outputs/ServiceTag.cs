// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Outputs
{

    /// <summary>
    /// The metadata that you apply to a resource to help you categorize and organize them. Each tag consists of a key and an optional value. You define them.
    ///  The following basic restrictions apply to tags:
    ///   +  Maximum number of tags per resource - 50
    ///   +  For each resource, each tag key must be unique, and each tag key can have only one value.
    ///   +  Maximum key length - 128 Unicode characters in UTF-8
    ///   +  Maximum value length - 256 Unicode characters in UTF-8
    ///   +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
    ///   +  Tag keys and values are case-sensitive.
    ///   +  Do not use ``aws:``, ``AWS:``, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
    /// </summary>
    [OutputType]
    public sealed class ServiceTag
    {
        /// <summary>
        /// One part of a key-value pair that make up a tag. A ``key`` is a general label that acts like a category for more specific tag values.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The optional part of a key-value pair that make up a tag. A ``value`` acts as a descriptor within a tag category (key).
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ServiceTag(
            string? key,

            string? value)
        {
            Key = key;
            Value = value;
        }
    }
}
