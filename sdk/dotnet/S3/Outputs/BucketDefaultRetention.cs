// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// The container element for optionally specifying the default Object Lock retention settings for new objects placed in the specified bucket.
    ///    +  The ``DefaultRetention`` settings require both a mode and a period.
    ///   +  The ``DefaultRetention`` period can be either ``Days`` or ``Years`` but you must select one. You cannot specify ``Days`` and ``Years`` at the same time.
    /// </summary>
    [OutputType]
    public sealed class BucketDefaultRetention
    {
        /// <summary>
        /// The number of days that you want to specify for the default retention period. If Object Lock is turned on, you must specify ``Mode`` and specify either ``Days`` or ``Years``.
        /// </summary>
        public readonly int? Days;
        /// <summary>
        /// The default Object Lock retention mode you want to apply to new objects placed in the specified bucket. If Object Lock is turned on, you must specify ``Mode`` and specify either ``Days`` or ``Years``.
        /// </summary>
        public readonly Pulumi.AwsNative.S3.BucketDefaultRetentionMode? Mode;
        /// <summary>
        /// The number of years that you want to specify for the default retention period. If Object Lock is turned on, you must specify ``Mode`` and specify either ``Days`` or ``Years``.
        /// </summary>
        public readonly int? Years;

        [OutputConstructor]
        private BucketDefaultRetention(
            int? days,

            Pulumi.AwsNative.S3.BucketDefaultRetentionMode? mode,

            int? years)
        {
            Days = days;
            Mode = mode;
            Years = years;
        }
    }
}
