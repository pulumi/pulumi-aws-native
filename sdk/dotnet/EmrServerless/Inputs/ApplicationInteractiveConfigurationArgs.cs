// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EmrServerless.Inputs
{

    public sealed class ApplicationInteractiveConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables an Apache Livy endpoint that you can connect to and run interactive jobs
        /// </summary>
        [Input("livyEndpointEnabled")]
        public Input<bool>? LivyEndpointEnabled { get; set; }

        /// <summary>
        /// Enabled you to connect an Application to Amazon EMR Studio to run interactive workloads in a notebook
        /// </summary>
        [Input("studioEnabled")]
        public Input<bool>? StudioEnabled { get; set; }

        public ApplicationInteractiveConfigurationArgs()
        {
        }
        public static new ApplicationInteractiveConfigurationArgs Empty => new ApplicationInteractiveConfigurationArgs();
    }
}