// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    /// <summary>
    /// An object to describe the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories.
    /// </summary>
    public sealed class ContainerPrivateRegistryAccessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An object to describe a request to activate or deactivate the role that you can use to grant an Amazon Lightsail container service access to Amazon Elastic Container Registry (Amazon ECR) private repositories.
        /// </summary>
        [Input("ecrImagePullerRole")]
        public Input<Inputs.ContainerPrivateRegistryAccessEcrImagePullerRolePropertiesArgs>? EcrImagePullerRole { get; set; }

        public ContainerPrivateRegistryAccessArgs()
        {
        }
        public static new ContainerPrivateRegistryAccessArgs Empty => new ContainerPrivateRegistryAccessArgs();
    }
}