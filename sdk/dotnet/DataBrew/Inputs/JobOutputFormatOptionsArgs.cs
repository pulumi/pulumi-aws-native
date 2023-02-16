// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Inputs
{

    /// <summary>
    /// Format options for job Output
    /// </summary>
    public sealed class JobOutputFormatOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("csv")]
        public Input<Inputs.JobCsvOutputOptionsArgs>? Csv { get; set; }

        public JobOutputFormatOptionsArgs()
        {
        }
        public static new JobOutputFormatOptionsArgs Empty => new JobOutputFormatOptionsArgs();
    }
}