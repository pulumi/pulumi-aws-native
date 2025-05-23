// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync.Inputs
{

    /// <summary>
    /// Specifies the level of reporting for the files, objects, and directories that Datasync attempted to transfer.
    /// </summary>
    public sealed class TaskReportConfigOverridesPropertiesTransferredPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't.
        /// </summary>
        [Input("reportLevel")]
        public Input<Pulumi.AwsNative.DataSync.TaskReportConfigOverridesPropertiesTransferredPropertiesReportLevel>? ReportLevel { get; set; }

        public TaskReportConfigOverridesPropertiesTransferredPropertiesArgs()
        {
        }
        public static new TaskReportConfigOverridesPropertiesTransferredPropertiesArgs Empty => new TaskReportConfigOverridesPropertiesTransferredPropertiesArgs();
    }
}
