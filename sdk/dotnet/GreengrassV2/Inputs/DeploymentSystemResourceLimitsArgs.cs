// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Inputs
{

    public sealed class DeploymentSystemResourceLimitsArgs : global::Pulumi.ResourceArgs
    {
        [Input("cpus")]
        public Input<double>? Cpus { get; set; }

        [Input("memory")]
        public Input<int>? Memory { get; set; }

        public DeploymentSystemResourceLimitsArgs()
        {
        }
        public static new DeploymentSystemResourceLimitsArgs Empty => new DeploymentSystemResourceLimitsArgs();
    }
}
