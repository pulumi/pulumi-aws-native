// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Inputs
{

    public sealed class DatasetQueryActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("filters")]
        private InputList<Inputs.DatasetFilterArgs>? _filters;
        public InputList<Inputs.DatasetFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.DatasetFilterArgs>());
            set => _filters = value;
        }

        [Input("sqlQuery", required: true)]
        public Input<string> SqlQuery { get; set; } = null!;

        public DatasetQueryActionArgs()
        {
        }
        public static new DatasetQueryActionArgs Empty => new DatasetQueryActionArgs();
    }
}