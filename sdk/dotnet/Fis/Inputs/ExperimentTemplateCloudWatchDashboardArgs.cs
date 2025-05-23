// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fis.Inputs
{

    public sealed class ExperimentTemplateCloudWatchDashboardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the CloudWatch dashboard to include in the experiment report.
        /// </summary>
        [Input("dashboardIdentifier", required: true)]
        public Input<string> DashboardIdentifier { get; set; } = null!;

        public ExperimentTemplateCloudWatchDashboardArgs()
        {
        }
        public static new ExperimentTemplateCloudWatchDashboardArgs Empty => new ExperimentTemplateCloudWatchDashboardArgs();
    }
}
