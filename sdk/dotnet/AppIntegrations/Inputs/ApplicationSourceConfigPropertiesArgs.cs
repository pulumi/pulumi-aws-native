// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppIntegrations.Inputs
{

    /// <summary>
    /// Application source config
    /// </summary>
    public sealed class ApplicationSourceConfigPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The external URL source for the application.
        /// </summary>
        [Input("externalUrlConfig", required: true)]
        public Input<Inputs.ApplicationExternalUrlConfigArgs> ExternalUrlConfig { get; set; } = null!;

        public ApplicationSourceConfigPropertiesArgs()
        {
        }
        public static new ApplicationSourceConfigPropertiesArgs Empty => new ApplicationSourceConfigPropertiesArgs();
    }
}