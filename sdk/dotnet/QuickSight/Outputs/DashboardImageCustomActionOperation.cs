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
    public sealed class DashboardImageCustomActionOperation
    {
        public readonly Outputs.DashboardCustomActionNavigationOperation? NavigationOperation;
        public readonly Outputs.DashboardCustomActionSetParametersOperation? SetParametersOperation;
        public readonly Outputs.DashboardCustomActionUrlOperation? UrlOperation;

        [OutputConstructor]
        private DashboardImageCustomActionOperation(
            Outputs.DashboardCustomActionNavigationOperation? navigationOperation,

            Outputs.DashboardCustomActionSetParametersOperation? setParametersOperation,

            Outputs.DashboardCustomActionUrlOperation? urlOperation)
        {
            NavigationOperation = navigationOperation;
            SetParametersOperation = setParametersOperation;
            UrlOperation = urlOperation;
        }
    }
}
