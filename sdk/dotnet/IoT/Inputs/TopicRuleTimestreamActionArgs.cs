// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class TopicRuleTimestreamActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of an Amazon Timestream database that has the table to write records into.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        [Input("dimensions", required: true)]
        private InputList<Inputs.TopicRuleTimestreamDimensionArgs>? _dimensions;

        /// <summary>
        /// Metadata attributes of the time series that are written in each measure record.
        /// </summary>
        public InputList<Inputs.TopicRuleTimestreamDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.TopicRuleTimestreamDimensionArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the role that grants AWS IoT permission to write to the Timestream database table.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The table where the message data will be written.
        /// </summary>
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        /// <summary>
        /// The value to use for the entry's timestamp. If blank, the time that the entry was processed is used.
        /// </summary>
        [Input("timestamp")]
        public Input<Inputs.TopicRuleTimestreamTimestampArgs>? Timestamp { get; set; }

        public TopicRuleTimestreamActionArgs()
        {
        }
        public static new TopicRuleTimestreamActionArgs Empty => new TopicRuleTimestreamActionArgs();
    }
}
