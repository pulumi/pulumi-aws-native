// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline.Inputs
{

    public sealed class QueuePosixUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public QueuePosixUserArgs()
        {
        }
        public static new QueuePosixUserArgs Empty => new QueuePosixUserArgs();
    }
}