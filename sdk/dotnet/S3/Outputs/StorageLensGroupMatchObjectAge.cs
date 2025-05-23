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
    /// Filter to match all of the specified values for the minimum and maximum object age.
    /// </summary>
    [OutputType]
    public sealed class StorageLensGroupMatchObjectAge
    {
        /// <summary>
        /// Minimum object age to which the rule applies.
        /// </summary>
        public readonly int? DaysGreaterThan;
        /// <summary>
        /// Maximum object age to which the rule applies.
        /// </summary>
        public readonly int? DaysLessThan;

        [OutputConstructor]
        private StorageLensGroupMatchObjectAge(
            int? daysGreaterThan,

            int? daysLessThan)
        {
            DaysGreaterThan = daysGreaterThan;
            DaysLessThan = daysLessThan;
        }
    }
}
