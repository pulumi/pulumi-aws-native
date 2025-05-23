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
    /// Source Code configuration
    /// </summary>
    public sealed class ServiceSourceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the resources that are needed to authenticate access to some source repositories.
        /// </summary>
        [Input("authenticationConfiguration")]
        public Input<Inputs.ServiceAuthenticationConfigurationArgs>? AuthenticationConfiguration { get; set; }

        /// <summary>
        /// Auto Deployment enabled
        /// </summary>
        [Input("autoDeploymentsEnabled")]
        public Input<bool>? AutoDeploymentsEnabled { get; set; }

        /// <summary>
        /// The description of a source code repository.
        /// 
        /// You must provide either this member or `ImageRepository` (but not both).
        /// </summary>
        [Input("codeRepository")]
        public Input<Inputs.ServiceCodeRepositoryArgs>? CodeRepository { get; set; }

        /// <summary>
        /// The description of a source image repository.
        /// 
        /// You must provide either this member or `CodeRepository` (but not both).
        /// </summary>
        [Input("imageRepository")]
        public Input<Inputs.ServiceImageRepositoryArgs>? ImageRepository { get; set; }

        public ServiceSourceConfigurationArgs()
        {
        }
        public static new ServiceSourceConfigurationArgs Empty => new ServiceSourceConfigurationArgs();
    }
}
