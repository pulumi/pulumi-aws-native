// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream.Inputs
{

    /// <summary>
    /// Configuration of target store where scheduled query results are written to.
    /// </summary>
    public sealed class ScheduledQueryTargetConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("timestreamConfiguration", required: true)]
        public Input<Inputs.ScheduledQueryTimestreamConfigurationArgs> TimestreamConfiguration { get; set; } = null!;

        public ScheduledQueryTargetConfigurationArgs()
        {
        }
        public static new ScheduledQueryTargetConfigurationArgs Empty => new ScheduledQueryTargetConfigurationArgs();
    }
}