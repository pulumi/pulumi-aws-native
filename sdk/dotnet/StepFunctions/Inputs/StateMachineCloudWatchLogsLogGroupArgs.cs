// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.StepFunctions.Inputs
{

    public sealed class StateMachineCloudWatchLogsLogGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("logGroupArn")]
        public Input<string>? LogGroupArn { get; set; }

        public StateMachineCloudWatchLogsLogGroupArgs()
        {
        }
        public static new StateMachineCloudWatchLogsLogGroupArgs Empty => new StateMachineCloudWatchLogsLogGroupArgs();
    }
}