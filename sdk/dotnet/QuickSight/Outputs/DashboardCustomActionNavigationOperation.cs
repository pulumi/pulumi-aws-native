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
    public sealed class DashboardCustomActionNavigationOperation
    {
        /// <summary>
        /// The configuration that chooses the navigation target.
        /// </summary>
        public readonly Outputs.DashboardLocalNavigationConfiguration? LocalNavigationConfiguration;

        [OutputConstructor]
        private DashboardCustomActionNavigationOperation(Outputs.DashboardLocalNavigationConfiguration? localNavigationConfiguration)
        {
            LocalNavigationConfiguration = localNavigationConfiguration;
        }
    }
}
