// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDb.Inputs
{

    public sealed class TableSseSpecificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("kmsMasterKeyId")]
        public Input<string>? KmsMasterKeyId { get; set; }

        [Input("sseEnabled", required: true)]
        public Input<bool> SseEnabled { get; set; } = null!;

        [Input("sseType")]
        public Input<string>? SseType { get; set; }

        public TableSseSpecificationArgs()
        {
        }
        public static new TableSseSpecificationArgs Empty => new TableSseSpecificationArgs();
    }
}