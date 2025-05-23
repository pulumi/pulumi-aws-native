// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline.Inputs
{

    public sealed class QueueJobRunAsUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The user and group that the jobs in the queue run as.
        /// </summary>
        [Input("posix")]
        public Input<Inputs.QueuePosixUserArgs>? Posix { get; set; }

        /// <summary>
        /// Specifies whether the job should run using the queue's system user or if the job should run using the worker agent system user.
        /// </summary>
        [Input("runAs", required: true)]
        public Input<Pulumi.AwsNative.Deadline.QueueRunAs> RunAs { get; set; } = null!;

        /// <summary>
        /// Identifies a Microsoft Windows user.
        /// </summary>
        [Input("windows")]
        public Input<Inputs.QueueWindowsUserArgs>? Windows { get; set; }

        public QueueJobRunAsUserArgs()
        {
        }
        public static new QueueJobRunAsUserArgs Empty => new QueueJobRunAsUserArgs();
    }
}
