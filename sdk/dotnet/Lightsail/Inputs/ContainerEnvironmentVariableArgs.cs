// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    public sealed class ContainerEnvironmentVariableArgs : global::Pulumi.ResourceArgs
    {
        [Input("value")]
        public Input<string>? Value { get; set; }

        [Input("variable")]
        public Input<string>? Variable { get; set; }

        public ContainerEnvironmentVariableArgs()
        {
        }
        public static new ContainerEnvironmentVariableArgs Empty => new ContainerEnvironmentVariableArgs();
    }
}