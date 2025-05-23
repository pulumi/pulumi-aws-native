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
    /// &lt;p&gt;The default values of a string parameter.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetStringDatasetParameterDefaultValues
    {
        /// <summary>
        /// &lt;p&gt;A list of static default values for a given string parameter.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<string> StaticValues;

        [OutputConstructor]
        private DataSetStringDatasetParameterDefaultValues(ImmutableArray<string> staticValues)
        {
            StaticValues = staticValues;
        }
    }
}
