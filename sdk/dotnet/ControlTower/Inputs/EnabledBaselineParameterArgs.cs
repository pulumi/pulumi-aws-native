// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ControlTower.Inputs
{

    public sealed class EnabledBaselineParameterArgs : global::Pulumi.ResourceArgs
    {
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("value")]
        public object? Value { get; set; }

        public EnabledBaselineParameterArgs()
        {
        }
        public static new EnabledBaselineParameterArgs Empty => new EnabledBaselineParameterArgs();
    }
}