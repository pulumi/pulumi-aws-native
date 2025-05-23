// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;A date-time parameter.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DashboardDateTimeParameter
    {
        /// <summary>
        /// &lt;p&gt;A display name for the date-time parameter.&lt;/p&gt;
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// &lt;p&gt;The values for the date-time parameter.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private DashboardDateTimeParameter(
            string name,

            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
}
