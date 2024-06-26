// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class FlowSapoDataSourcePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object path specified in the SAPOData flow source.
        /// </summary>
        [Input("objectPath", required: true)]
        public Input<string> ObjectPath { get; set; } = null!;

        [Input("paginationConfig")]
        public Input<Inputs.FlowSapoDataPaginationConfigArgs>? PaginationConfig { get; set; }

        [Input("parallelismConfig")]
        public Input<Inputs.FlowSapoDataParallelismConfigArgs>? ParallelismConfig { get; set; }

        public FlowSapoDataSourcePropertiesArgs()
        {
        }
        public static new FlowSapoDataSourcePropertiesArgs Empty => new FlowSapoDataSourcePropertiesArgs();
    }
}
