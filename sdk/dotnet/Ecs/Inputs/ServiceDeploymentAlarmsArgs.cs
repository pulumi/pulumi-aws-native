// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    public sealed class ServiceDeploymentAlarmsArgs : global::Pulumi.ResourceArgs
    {
        [Input("alarmNames", required: true)]
        private InputList<string>? _alarmNames;
        public InputList<string> AlarmNames
        {
            get => _alarmNames ?? (_alarmNames = new InputList<string>());
            set => _alarmNames = value;
        }

        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        [Input("rollback", required: true)]
        public Input<bool> Rollback { get; set; } = null!;

        public ServiceDeploymentAlarmsArgs()
        {
        }
        public static new ServiceDeploymentAlarmsArgs Empty => new ServiceDeploymentAlarmsArgs();
    }
}