// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fis.Inputs
{

    /// <summary>
    /// The configuration for experiment logging to CloudWatch Logs .
    /// </summary>
    public sealed class ExperimentTemplateLogConfigurationCloudWatchLogsConfigurationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("logGroupArn", required: true)]
        public Input<string> LogGroupArn { get; set; } = null!;

        public ExperimentTemplateLogConfigurationCloudWatchLogsConfigurationPropertiesArgs()
        {
        }
        public static new ExperimentTemplateLogConfigurationCloudWatchLogsConfigurationPropertiesArgs Empty => new ExperimentTemplateLogConfigurationCloudWatchLogsConfigurationPropertiesArgs();
    }
}
