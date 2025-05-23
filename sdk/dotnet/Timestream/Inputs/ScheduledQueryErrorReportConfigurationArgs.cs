// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream.Inputs
{

    /// <summary>
    /// Configuration for error reporting. Error reports will be generated when a problem is encountered when writing the query results.
    /// </summary>
    public sealed class ScheduledQueryErrorReportConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The S3 configuration for the error reports.
        /// </summary>
        [Input("s3Configuration", required: true)]
        public Input<Inputs.ScheduledQueryS3ConfigurationArgs> S3Configuration { get; set; } = null!;

        public ScheduledQueryErrorReportConfigurationArgs()
        {
        }
        public static new ScheduledQueryErrorReportConfigurationArgs Empty => new ScheduledQueryErrorReportConfigurationArgs();
    }
}
