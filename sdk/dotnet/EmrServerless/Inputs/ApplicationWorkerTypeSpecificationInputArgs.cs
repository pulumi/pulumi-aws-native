// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EmrServerless.Inputs
{

    /// <summary>
    /// The specifications for a worker type.
    /// </summary>
    public sealed class ApplicationWorkerTypeSpecificationInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The image configuration for a worker type.
        /// </summary>
        [Input("imageConfiguration")]
        public Input<Inputs.ApplicationImageConfigurationInputArgs>? ImageConfiguration { get; set; }

        public ApplicationWorkerTypeSpecificationInputArgs()
        {
        }
        public static new ApplicationWorkerTypeSpecificationInputArgs Empty => new ApplicationWorkerTypeSpecificationInputArgs();
    }
}
