// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Amplify.Inputs
{

    public sealed class AppJobConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the size of the build instance. Amplify supports three instance sizes: `STANDARD_8GB` , `LARGE_16GB` , and `XLARGE_72GB` . If you don't specify a value, Amplify uses the `STANDARD_8GB` default.
        /// 
        /// The following list describes the CPU, memory, and storage capacity for each build instance type:
        /// 
        /// - **STANDARD_8GB** - - vCPUs: 4
        /// - Memory: 8 GiB
        /// - Disk space: 128 GB
        /// - **LARGE_16GB** - - vCPUs: 8
        /// - Memory: 16 GiB
        /// - Disk space: 128 GB
        /// - **XLARGE_72GB** - - vCPUs: 36
        /// - Memory: 72 GiB
        /// - Disk space: 256 GB
        /// </summary>
        [Input("buildComputeType", required: true)]
        public Input<Pulumi.AwsNative.Amplify.AppJobConfigBuildComputeType> BuildComputeType { get; set; } = null!;

        public AppJobConfigArgs()
        {
        }
        public static new AppJobConfigArgs Empty => new AppJobConfigArgs();
    }
}
