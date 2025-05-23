// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardTablePinnedFieldOptions
    {
        /// <summary>
        /// A list of columns to be pinned to the left of a table visual.
        /// </summary>
        public readonly ImmutableArray<string> PinnedLeftFields;

        [OutputConstructor]
        private DashboardTablePinnedFieldOptions(ImmutableArray<string> pinnedLeftFields)
        {
            PinnedLeftFields = pinnedLeftFields;
        }
    }
}
