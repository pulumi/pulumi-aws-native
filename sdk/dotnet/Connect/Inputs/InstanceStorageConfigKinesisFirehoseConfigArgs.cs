// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    public sealed class InstanceStorageConfigKinesisFirehoseConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("firehoseArn", required: true)]
        public Input<string> FirehoseArn { get; set; } = null!;

        public InstanceStorageConfigKinesisFirehoseConfigArgs()
        {
        }
        public static new InstanceStorageConfigKinesisFirehoseConfigArgs Empty => new InstanceStorageConfigKinesisFirehoseConfigArgs();
    }
}