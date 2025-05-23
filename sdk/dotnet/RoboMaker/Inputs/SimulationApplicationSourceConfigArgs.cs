// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RoboMaker.Inputs
{

    /// <summary>
    /// Information about a source configuration.
    /// </summary>
    public sealed class SimulationApplicationSourceConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The target processor architecture for the application.
        /// </summary>
        [Input("architecture", required: true)]
        public Input<Pulumi.AwsNative.RoboMaker.SimulationApplicationSourceConfigArchitecture> Architecture { get; set; } = null!;

        /// <summary>
        /// The Amazon S3 bucket name.
        /// </summary>
        [Input("s3Bucket", required: true)]
        public Input<string> S3Bucket { get; set; } = null!;

        /// <summary>
        /// The s3 object key.
        /// </summary>
        [Input("s3Key", required: true)]
        public Input<string> S3Key { get; set; } = null!;

        public SimulationApplicationSourceConfigArgs()
        {
        }
        public static new SimulationApplicationSourceConfigArgs Empty => new SimulationApplicationSourceConfigArgs();
    }
}
