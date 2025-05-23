// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EmrServerless.Outputs
{

    [OutputType]
    public sealed class ApplicationInteractiveConfiguration
    {
        /// <summary>
        /// Enables an Apache Livy endpoint that you can connect to and run interactive jobs
        /// </summary>
        public readonly bool? LivyEndpointEnabled;
        /// <summary>
        /// Enabled you to connect an Application to Amazon EMR Studio to run interactive workloads in a notebook
        /// </summary>
        public readonly bool? StudioEnabled;

        [OutputConstructor]
        private ApplicationInteractiveConfiguration(
            bool? livyEndpointEnabled,

            bool? studioEnabled)
        {
            LivyEndpointEnabled = livyEndpointEnabled;
            StudioEnabled = studioEnabled;
        }
    }
}
