// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    [OutputType]
    public sealed class JobStatisticOverride
    {
        /// <summary>
        /// A map that includes overrides of an evaluation’s parameters.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Parameters;
        /// <summary>
        /// The name of an evaluation
        /// </summary>
        public readonly string Statistic;

        [OutputConstructor]
        private JobStatisticOverride(
            ImmutableDictionary<string, string> parameters,

            string statistic)
        {
            Parameters = parameters;
            Statistic = statistic;
        }
    }
}
