// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationInsights.Outputs
{

    /// <summary>
    /// The log pattern.
    /// </summary>
    [OutputType]
    public sealed class ApplicationLogPattern
    {
        /// <summary>
        /// The log pattern.
        /// </summary>
        public readonly string Pattern;
        /// <summary>
        /// The name of the log pattern.
        /// </summary>
        public readonly string PatternName;
        /// <summary>
        /// Rank of the log pattern.
        /// </summary>
        public readonly int Rank;

        [OutputConstructor]
        private ApplicationLogPattern(
            string pattern,

            string patternName,

            int rank)
        {
            Pattern = pattern;
            PatternName = patternName;
            Rank = rank;
        }
    }
}