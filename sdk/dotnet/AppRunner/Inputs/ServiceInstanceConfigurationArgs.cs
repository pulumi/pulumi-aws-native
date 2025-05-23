// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppRunner.Inputs
{

    /// <summary>
    /// Instance Configuration
    /// </summary>
    public sealed class ServiceInstanceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CPU
        /// </summary>
        [Input("cpu")]
        public Input<string>? Cpu { get; set; }

        /// <summary>
        /// Instance Role Arn
        /// </summary>
        [Input("instanceRoleArn")]
        public Input<string>? InstanceRoleArn { get; set; }

        /// <summary>
        /// Memory
        /// </summary>
        [Input("memory")]
        public Input<string>? Memory { get; set; }

        public ServiceInstanceConfigurationArgs()
        {
        }
        public static new ServiceInstanceConfigurationArgs Empty => new ServiceInstanceConfigurationArgs();
    }
}
