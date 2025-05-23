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
    /// &lt;p&gt;The incremental refresh configuration for a dataset.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetIncrementalRefresh
    {
        /// <summary>
        /// The lookback window setup for an incremental refresh configuration.
        /// </summary>
        public readonly Outputs.DataSetLookbackWindow LookbackWindow;

        [OutputConstructor]
        private DataSetIncrementalRefresh(Outputs.DataSetLookbackWindow lookbackWindow)
        {
            LookbackWindow = lookbackWindow;
        }
    }
}
